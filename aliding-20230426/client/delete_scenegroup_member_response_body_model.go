// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScenegroupMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteScenegroupMemberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteScenegroupMemberResponseBody
	GetSuccess() *bool
}

type DeleteScenegroupMemberResponseBody struct {
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

func (s DeleteScenegroupMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteScenegroupMemberResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScenegroupMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteScenegroupMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteScenegroupMemberResponseBody) SetRequestId(v string) *DeleteScenegroupMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteScenegroupMemberResponseBody) SetSuccess(v bool) *DeleteScenegroupMemberResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteScenegroupMemberResponseBody) Validate() error {
	return dara.Validate(s)
}
