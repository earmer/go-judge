package filestore

import (
	"errors"
	"github.com/google/uuid"
	"os"

	"github.com/criyle/go-judge/envexec"
)

var errUniqueIDNotGenerated = errors.New("failed to generate unique id")

// FileStore defines interface to store file
type FileStore interface {
	Add(name, path string) (string, error) // Add creates a file with path to the storage, returns id
	Remove(string) bool                    // Remove deletes a file by id
	Get(string) (string, envexec.File)     // Get file by id, nil if not exists
	List() map[string]string               // List return all file ids to original name
	New() (*os.File, error)                // Create a temporary file to the file store, can be added through Add to save it
}

func generateID() (string, error) {
	u, err := uuid.NewUUID()
	if err != nil {
		return "", errUniqueIDNotGenerated
	}
	return u.String(), nil
}
