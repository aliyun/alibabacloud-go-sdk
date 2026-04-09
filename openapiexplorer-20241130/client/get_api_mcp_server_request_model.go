// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApiMcpServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetApiMcpServerRequest
	GetId() *string
}

type GetApiMcpServerRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// v6ZZ7ftCzEILW***
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s GetApiMcpServerRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApiMcpServerRequest) GoString() string {
	return s.String()
}

func (s *GetApiMcpServerRequest) GetId() *string {
	return s.Id
}

func (s *GetApiMcpServerRequest) SetId(v string) *GetApiMcpServerRequest {
	s.Id = &v
	return s
}

func (s *GetApiMcpServerRequest) Validate() error {
	return dara.Validate(s)
}
