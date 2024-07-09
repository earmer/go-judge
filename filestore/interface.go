package filestore

import (
	"errors"
	"os"

	"github.com/criyle/go-judge/envexec"
	"github.com/google/uuid"
)

var errUniqueIDNotGenerated = errors.New("unique id not generated")

// FileStore defines interface to store file
type FileStore interface {
	Add(name, path string) (string, error) // Add creates a file with path to the storage, returns id
	Remove(string) bool                    // Remove deletes a file by id
	Get(string) (string, envexec.File)     // Get file by id, nil if not exists
	List() map[string]string               // List return all file ids to original name
	New() (*os.File, error)                // Create a temporary file to the file store, can be added through Add to save it
}

func generateID() (string, error) {
	uuid1, err := uuid.NewUUID() // Using UUID V1 for unique id
	if err != nil {
		return "", errUniqueIDNotGenerated
	}
	return uuid1.String(), nil
}
