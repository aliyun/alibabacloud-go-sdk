// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStandardGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteStandardGroupResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteStandardGroupResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteStandardGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteStandardGroupResponseBody
	GetSuccess() *bool
}

type DeleteStandardGroupResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message that is returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request. You can use the request ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// C5B8E84B-42B6-4374-AD5A-6264E1753378
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteStandardGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteStandardGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStandardGroupResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteStandardGroupResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteStandardGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteStandardGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteStandardGroupResponseBody) SetErrorCode(v string) *DeleteStandardGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteStandardGroupResponseBody) SetErrorMessage(v string) *DeleteStandardGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteStandardGroupResponseBody) SetRequestId(v string) *DeleteStandardGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteStandardGroupResponseBody) SetSuccess(v bool) *DeleteStandardGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteStandardGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
