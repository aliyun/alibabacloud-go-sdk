// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserGroupsMappingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateUserGroupsMappingResponseBody
	GetRequestId() *string
}

type CreateUserGroupsMappingResponseBody struct {
	// example:
	//
	// 55C5FFD6-BF99-41BD-9C66-FFF39189****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateUserGroupsMappingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUserGroupsMappingResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserGroupsMappingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUserGroupsMappingResponseBody) SetRequestId(v string) *CreateUserGroupsMappingResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserGroupsMappingResponseBody) Validate() error {
	return dara.Validate(s)
}
