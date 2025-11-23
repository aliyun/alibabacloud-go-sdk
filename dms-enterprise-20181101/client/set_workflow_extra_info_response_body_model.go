// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetWorkflowExtraInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *SetWorkflowExtraInfoResponseBody
	GetData() *bool
	SetErrorCode(v string) *SetWorkflowExtraInfoResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *SetWorkflowExtraInfoResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *SetWorkflowExtraInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SetWorkflowExtraInfoResponseBody
	GetSuccess() *bool
}

type SetWorkflowExtraInfoResponseBody struct {
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// The request ID. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// 8401893F-4235-55D5-B563-7CF7A7D037DE
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

func (s SetWorkflowExtraInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetWorkflowExtraInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SetWorkflowExtraInfoResponseBody) GetData() *bool {
	return s.Data
}

func (s *SetWorkflowExtraInfoResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SetWorkflowExtraInfoResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SetWorkflowExtraInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetWorkflowExtraInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetWorkflowExtraInfoResponseBody) SetData(v bool) *SetWorkflowExtraInfoResponseBody {
	s.Data = &v
	return s
}

func (s *SetWorkflowExtraInfoResponseBody) SetErrorCode(v string) *SetWorkflowExtraInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SetWorkflowExtraInfoResponseBody) SetErrorMessage(v string) *SetWorkflowExtraInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SetWorkflowExtraInfoResponseBody) SetRequestId(v string) *SetWorkflowExtraInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetWorkflowExtraInfoResponseBody) SetSuccess(v bool) *SetWorkflowExtraInfoResponseBody {
	s.Success = &v
	return s
}

func (s *SetWorkflowExtraInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
