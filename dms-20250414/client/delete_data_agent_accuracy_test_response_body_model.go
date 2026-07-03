// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataAgentAccuracyTestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteDataAgentAccuracyTestResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteDataAgentAccuracyTestResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteDataAgentAccuracyTestResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DeleteDataAgentAccuracyTestResponseBody
	GetSuccess() *string
}

type DeleteDataAgentAccuracyTestResponseBody struct {
	// The error code.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the call failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E0D21075-xxx-FD8AD04A63B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDataAgentAccuracyTestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataAgentAccuracyTestResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataAgentAccuracyTestResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteDataAgentAccuracyTestResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteDataAgentAccuracyTestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataAgentAccuracyTestResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DeleteDataAgentAccuracyTestResponseBody) SetErrorCode(v string) *DeleteDataAgentAccuracyTestResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteDataAgentAccuracyTestResponseBody) SetErrorMessage(v string) *DeleteDataAgentAccuracyTestResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteDataAgentAccuracyTestResponseBody) SetRequestId(v string) *DeleteDataAgentAccuracyTestResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataAgentAccuracyTestResponseBody) SetSuccess(v string) *DeleteDataAgentAccuracyTestResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDataAgentAccuracyTestResponseBody) Validate() error {
	return dara.Validate(s)
}
