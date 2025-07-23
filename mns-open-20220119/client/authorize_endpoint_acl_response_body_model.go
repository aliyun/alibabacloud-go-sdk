// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeEndpointAclResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *AuthorizeEndpointAclResponseBody
	GetCode() *int64
	SetMessage(v string) *AuthorizeEndpointAclResponseBody
	GetMessage() *string
	SetRequestId(v string) *AuthorizeEndpointAclResponseBody
	GetRequestId() *string
	SetStatus(v string) *AuthorizeEndpointAclResponseBody
	GetStatus() *string
	SetSuccess(v bool) *AuthorizeEndpointAclResponseBody
	GetSuccess() *bool
}

type AuthorizeEndpointAclResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 06273500-249F-5863-121D-74D51123****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response status.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AuthorizeEndpointAclResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeEndpointAclResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeEndpointAclResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *AuthorizeEndpointAclResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AuthorizeEndpointAclResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AuthorizeEndpointAclResponseBody) GetStatus() *string {
	return s.Status
}

func (s *AuthorizeEndpointAclResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AuthorizeEndpointAclResponseBody) SetCode(v int64) *AuthorizeEndpointAclResponseBody {
	s.Code = &v
	return s
}

func (s *AuthorizeEndpointAclResponseBody) SetMessage(v string) *AuthorizeEndpointAclResponseBody {
	s.Message = &v
	return s
}

func (s *AuthorizeEndpointAclResponseBody) SetRequestId(v string) *AuthorizeEndpointAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthorizeEndpointAclResponseBody) SetStatus(v string) *AuthorizeEndpointAclResponseBody {
	s.Status = &v
	return s
}

func (s *AuthorizeEndpointAclResponseBody) SetSuccess(v bool) *AuthorizeEndpointAclResponseBody {
	s.Success = &v
	return s
}

func (s *AuthorizeEndpointAclResponseBody) Validate() error {
	return dara.Validate(s)
}
