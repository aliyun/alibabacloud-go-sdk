// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitializeOperationRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *InitializeOperationRoleRequest
	GetClientToken() *string
	SetBody(v string) *InitializeOperationRoleRequest
	GetBody() *string
}

type InitializeOperationRoleRequest struct {
	// The client token that is used to ensure the idempotence of the request. The client generates the value. The value must be unique among different requests and can be up to 64 ASCII characters long.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The name of the Elasticsearch service-linked role to be created.
	//
	// example:
	//
	// {"rolename":"AliyunServiceRoleForElasticsearchCollector"}
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitializeOperationRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s InitializeOperationRoleRequest) GoString() string {
	return s.String()
}

func (s *InitializeOperationRoleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *InitializeOperationRoleRequest) GetBody() *string {
	return s.Body
}

func (s *InitializeOperationRoleRequest) SetClientToken(v string) *InitializeOperationRoleRequest {
	s.ClientToken = &v
	return s
}

func (s *InitializeOperationRoleRequest) SetBody(v string) *InitializeOperationRoleRequest {
	s.Body = &v
	return s
}

func (s *InitializeOperationRoleRequest) Validate() error {
	return dara.Validate(s)
}
