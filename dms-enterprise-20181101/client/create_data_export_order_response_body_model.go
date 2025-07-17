// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataExportOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateOrderResult(v *CreateDataExportOrderResponseBodyCreateOrderResult) *CreateDataExportOrderResponseBody
	GetCreateOrderResult() *CreateDataExportOrderResponseBodyCreateOrderResult
	SetErrorCode(v string) *CreateDataExportOrderResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateDataExportOrderResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateDataExportOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDataExportOrderResponseBody
	GetSuccess() *bool
}

type CreateDataExportOrderResponseBody struct {
	// The content of the ticket.
	CreateOrderResult *CreateDataExportOrderResponseBodyCreateOrderResult `json:"CreateOrderResult,omitempty" xml:"CreateOrderResult,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// 7FAD400F-7A5C-4193-8F9A-39D86C4F0231
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDataExportOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataExportOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataExportOrderResponseBody) GetCreateOrderResult() *CreateDataExportOrderResponseBodyCreateOrderResult {
	return s.CreateOrderResult
}

func (s *CreateDataExportOrderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateDataExportOrderResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateDataExportOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataExportOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDataExportOrderResponseBody) SetCreateOrderResult(v *CreateDataExportOrderResponseBodyCreateOrderResult) *CreateDataExportOrderResponseBody {
	s.CreateOrderResult = v
	return s
}

func (s *CreateDataExportOrderResponseBody) SetErrorCode(v string) *CreateDataExportOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDataExportOrderResponseBody) SetErrorMessage(v string) *CreateDataExportOrderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDataExportOrderResponseBody) SetRequestId(v string) *CreateDataExportOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataExportOrderResponseBody) SetSuccess(v bool) *CreateDataExportOrderResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDataExportOrderResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateDataExportOrderResponseBodyCreateOrderResult struct {
	CreateOrderResult []*int64 `json:"CreateOrderResult,omitempty" xml:"CreateOrderResult,omitempty" type:"Repeated"`
}

func (s CreateDataExportOrderResponseBodyCreateOrderResult) String() string {
	return dara.Prettify(s)
}

func (s CreateDataExportOrderResponseBodyCreateOrderResult) GoString() string {
	return s.String()
}

func (s *CreateDataExportOrderResponseBodyCreateOrderResult) GetCreateOrderResult() []*int64 {
	return s.CreateOrderResult
}

func (s *CreateDataExportOrderResponseBodyCreateOrderResult) SetCreateOrderResult(v []*int64) *CreateDataExportOrderResponseBodyCreateOrderResult {
	s.CreateOrderResult = v
	return s
}

func (s *CreateDataExportOrderResponseBodyCreateOrderResult) Validate() error {
	return dara.Validate(s)
}
