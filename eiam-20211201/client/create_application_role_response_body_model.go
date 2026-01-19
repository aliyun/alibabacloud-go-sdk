// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationRoleId(v string) *CreateApplicationRoleResponseBody
	GetApplicationRoleId() *string
	SetRequestId(v string) *CreateApplicationRoleResponseBody
	GetRequestId() *string
}

type CreateApplicationRoleResponseBody struct {
	// example:
	//
	// approle_01kghbvoptu5262q35aalvq7cxxxx
	ApplicationRoleId *string `json:"ApplicationRoleId,omitempty" xml:"ApplicationRoleId,omitempty"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateApplicationRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApplicationRoleResponseBody) GetApplicationRoleId() *string {
	return s.ApplicationRoleId
}

func (s *CreateApplicationRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApplicationRoleResponseBody) SetApplicationRoleId(v string) *CreateApplicationRoleResponseBody {
	s.ApplicationRoleId = &v
	return s
}

func (s *CreateApplicationRoleResponseBody) SetRequestId(v string) *CreateApplicationRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApplicationRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
