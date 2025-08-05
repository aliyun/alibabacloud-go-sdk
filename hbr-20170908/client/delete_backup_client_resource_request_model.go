// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackupClientResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientIds(v map[string]interface{}) *DeleteBackupClientResourceRequest
	GetClientIds() map[string]interface{}
}

type DeleteBackupClientResourceRequest struct {
	// The IDs of HBR clients. The value can be a JSON array that consists of up to 100 client IDs. Separate the IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ["c-0007kyu045r0********", "c-000b6818umvo********"]
	ClientIds map[string]interface{} `json:"ClientIds,omitempty" xml:"ClientIds,omitempty"`
}

func (s DeleteBackupClientResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupClientResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteBackupClientResourceRequest) GetClientIds() map[string]interface{} {
	return s.ClientIds
}

func (s *DeleteBackupClientResourceRequest) SetClientIds(v map[string]interface{}) *DeleteBackupClientResourceRequest {
	s.ClientIds = v
	return s
}

func (s *DeleteBackupClientResourceRequest) Validate() error {
	return dara.Validate(s)
}
