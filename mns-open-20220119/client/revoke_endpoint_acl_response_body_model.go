// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeEndpointAclResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *RevokeEndpointAclResponseBody
	GetCode() *int64
	SetMessage(v string) *RevokeEndpointAclResponseBody
	GetMessage() *string
	SetRequestId(v string) *RevokeEndpointAclResponseBody
	GetRequestId() *string
	SetStatus(v string) *RevokeEndpointAclResponseBody
	GetStatus() *string
	SetSuccess(v bool) *RevokeEndpointAclResponseBody
	GetSuccess() *bool
}

type RevokeEndpointAclResponseBody struct {
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

func (s RevokeEndpointAclResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeEndpointAclResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeEndpointAclResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *RevokeEndpointAclResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RevokeEndpointAclResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeEndpointAclResponseBody) GetStatus() *string {
	return s.Status
}

func (s *RevokeEndpointAclResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RevokeEndpointAclResponseBody) SetCode(v int64) *RevokeEndpointAclResponseBody {
	s.Code = &v
	return s
}

func (s *RevokeEndpointAclResponseBody) SetMessage(v string) *RevokeEndpointAclResponseBody {
	s.Message = &v
	return s
}

func (s *RevokeEndpointAclResponseBody) SetRequestId(v string) *RevokeEndpointAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeEndpointAclResponseBody) SetStatus(v string) *RevokeEndpointAclResponseBody {
	s.Status = &v
	return s
}

func (s *RevokeEndpointAclResponseBody) SetSuccess(v bool) *RevokeEndpointAclResponseBody {
	s.Success = &v
	return s
}

func (s *RevokeEndpointAclResponseBody) Validate() error {
	return dara.Validate(s)
}
