// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseProduceAuthorizationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ReleaseProduceAuthorizationResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *ReleaseProduceAuthorizationResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *ReleaseProduceAuthorizationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReleaseProduceAuthorizationResponseBody
	GetSuccess() *bool
}

type ReleaseProduceAuthorizationResponseBody struct {
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 2174AA97-56FB-50FA-B243-0460B9E4CE0C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReleaseProduceAuthorizationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseProduceAuthorizationResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseProduceAuthorizationResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ReleaseProduceAuthorizationResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ReleaseProduceAuthorizationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseProduceAuthorizationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReleaseProduceAuthorizationResponseBody) SetErrorCode(v string) *ReleaseProduceAuthorizationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ReleaseProduceAuthorizationResponseBody) SetErrorMsg(v string) *ReleaseProduceAuthorizationResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ReleaseProduceAuthorizationResponseBody) SetRequestId(v string) *ReleaseProduceAuthorizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseProduceAuthorizationResponseBody) SetSuccess(v bool) *ReleaseProduceAuthorizationResponseBody {
	s.Success = &v
	return s
}

func (s *ReleaseProduceAuthorizationResponseBody) Validate() error {
	return dara.Validate(s)
}
