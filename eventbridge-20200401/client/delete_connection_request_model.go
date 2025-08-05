// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionName(v string) *DeleteConnectionRequest
	GetConnectionName() *string
}

type DeleteConnectionRequest struct {
	// The name of the connection that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// connection-name
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
}

func (s DeleteConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteConnectionRequest) GoString() string {
	return s.String()
}

func (s *DeleteConnectionRequest) GetConnectionName() *string {
	return s.ConnectionName
}

func (s *DeleteConnectionRequest) SetConnectionName(v string) *DeleteConnectionRequest {
	s.ConnectionName = &v
	return s
}

func (s *DeleteConnectionRequest) Validate() error {
	return dara.Validate(s)
}
