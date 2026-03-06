// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApiMcpServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteApiMcpServerRequest
	GetClientToken() *string
	SetId(v string) *DeleteApiMcpServerRequest
	GetId() *string
}

type DeleteApiMcpServerRequest struct {
	// The client token used to ensure the idempotence of the request. Generate this value on your client and make sure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. Use a universally unique identifier (UUID). The token is valid for three days.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The ID of the API MCP service.
	//
	// This parameter is required.
	//
	// example:
	//
	// v6ZZ7ftCzEILW***
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s DeleteApiMcpServerRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteApiMcpServerRequest) GoString() string {
	return s.String()
}

func (s *DeleteApiMcpServerRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteApiMcpServerRequest) GetId() *string {
	return s.Id
}

func (s *DeleteApiMcpServerRequest) SetClientToken(v string) *DeleteApiMcpServerRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteApiMcpServerRequest) SetId(v string) *DeleteApiMcpServerRequest {
	s.Id = &v
	return s
}

func (s *DeleteApiMcpServerRequest) Validate() error {
	return dara.Validate(s)
}
