// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectAuthorizationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationId(v string) *CreateProjectAuthorizationResponseBody
	GetAuthorizationId() *string
	SetRequestId(v string) *CreateProjectAuthorizationResponseBody
	GetRequestId() *string
}

type CreateProjectAuthorizationResponseBody struct {
	// The project authorization ID.
	//
	// example:
	//
	// 2
	AuthorizationId *string `json:"AuthorizationId,omitempty" xml:"AuthorizationId,omitempty"`
	// The unique identifier that Alibaba Cloud generated for the request.
	//
	// example:
	//
	// 4F6C075F-FC86-476E-943B-097BD4E12948
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateProjectAuthorizationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectAuthorizationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectAuthorizationResponseBody) GetAuthorizationId() *string {
	return s.AuthorizationId
}

func (s *CreateProjectAuthorizationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateProjectAuthorizationResponseBody) SetAuthorizationId(v string) *CreateProjectAuthorizationResponseBody {
	s.AuthorizationId = &v
	return s
}

func (s *CreateProjectAuthorizationResponseBody) SetRequestId(v string) *CreateProjectAuthorizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProjectAuthorizationResponseBody) Validate() error {
	return dara.Validate(s)
}
