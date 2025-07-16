// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddScenegroupMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddScenegroupMemberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddScenegroupMemberResponseBody
	GetSuccess() *bool
}

type AddScenegroupMemberResponseBody struct {
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddScenegroupMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddScenegroupMemberResponseBody) GoString() string {
	return s.String()
}

func (s *AddScenegroupMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddScenegroupMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddScenegroupMemberResponseBody) SetRequestId(v string) *AddScenegroupMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddScenegroupMemberResponseBody) SetSuccess(v bool) *AddScenegroupMemberResponseBody {
	s.Success = &v
	return s
}

func (s *AddScenegroupMemberResponseBody) Validate() error {
	return dara.Validate(s)
}
