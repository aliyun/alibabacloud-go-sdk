// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionName(v string) *GetConnectionRequest
	GetConnectionName() *string
}

type GetConnectionRequest struct {
	// The connection name.
	//
	// This parameter is required.
	//
	// example:
	//
	// connection-name
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
}

func (s GetConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetConnectionRequest) GoString() string {
	return s.String()
}

func (s *GetConnectionRequest) GetConnectionName() *string {
	return s.ConnectionName
}

func (s *GetConnectionRequest) SetConnectionName(v string) *GetConnectionRequest {
	s.ConnectionName = &v
	return s
}

func (s *GetConnectionRequest) Validate() error {
	return dara.Validate(s)
}
