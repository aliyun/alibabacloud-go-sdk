// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFailedReason interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *FailedReason
	GetErrorCode() *string
	SetErrorMessage(v string) *FailedReason
	GetErrorMessage() *string
	SetRequestId(v string) *FailedReason
	GetRequestId() *string
}

type FailedReason struct {
	// 错误码。
	//
	// example:
	//
	// MissingParameter.InstanceType
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息。
	//
	// example:
	//
	// The instance type is required.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求ID。
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944abcd
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FailedReason) String() string {
	return dara.Prettify(s)
}

func (s FailedReason) GoString() string {
	return s.String()
}

func (s *FailedReason) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *FailedReason) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *FailedReason) GetRequestId() *string {
	return s.RequestId
}

func (s *FailedReason) SetErrorCode(v string) *FailedReason {
	s.ErrorCode = &v
	return s
}

func (s *FailedReason) SetErrorMessage(v string) *FailedReason {
	s.ErrorMessage = &v
	return s
}

func (s *FailedReason) SetRequestId(v string) *FailedReason {
	s.RequestId = &v
	return s
}

func (s *FailedReason) Validate() error {
	return dara.Validate(s)
}
