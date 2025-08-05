// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackupClientRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *DeleteBackupClientRequest
	GetClientId() *string
}

type DeleteBackupClientRequest struct {
	// The ID of the Cloud Backup client.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-*********************
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
}

func (s DeleteBackupClientRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupClientRequest) GoString() string {
	return s.String()
}

func (s *DeleteBackupClientRequest) GetClientId() *string {
	return s.ClientId
}

func (s *DeleteBackupClientRequest) SetClientId(v string) *DeleteBackupClientRequest {
	s.ClientId = &v
	return s
}

func (s *DeleteBackupClientRequest) Validate() error {
	return dara.Validate(s)
}
