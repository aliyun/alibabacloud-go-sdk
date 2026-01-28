// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAuthorizationResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationResourceId(v string) *CreateAuthorizationResourceResponseBody
	GetAuthorizationResourceId() *string
	SetRequestId(v string) *CreateAuthorizationResourceResponseBody
	GetRequestId() *string
}

type CreateAuthorizationResourceResponseBody struct {
	// example:
	//
	// arres_01kgh3jvt7pk093rv6giu0c0qxxxx
	AuthorizationResourceId *string `json:"AuthorizationResourceId,omitempty" xml:"AuthorizationResourceId,omitempty"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAuthorizationResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAuthorizationResourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAuthorizationResourceResponseBody) GetAuthorizationResourceId() *string {
	return s.AuthorizationResourceId
}

func (s *CreateAuthorizationResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAuthorizationResourceResponseBody) SetAuthorizationResourceId(v string) *CreateAuthorizationResourceResponseBody {
	s.AuthorizationResourceId = &v
	return s
}

func (s *CreateAuthorizationResourceResponseBody) SetRequestId(v string) *CreateAuthorizationResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAuthorizationResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
