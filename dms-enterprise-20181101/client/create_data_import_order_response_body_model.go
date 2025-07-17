// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataImportOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateOrderResult(v []*int64) *CreateDataImportOrderResponseBody
	GetCreateOrderResult() []*int64
	SetErrorCode(v string) *CreateDataImportOrderResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateDataImportOrderResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateDataImportOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDataImportOrderResponseBody
	GetSuccess() *bool
}

type CreateDataImportOrderResponseBody struct {
	// The ID of the ticket.
	CreateOrderResult []*int64 `json:"CreateOrderResult,omitempty" xml:"CreateOrderResult,omitempty" type:"Repeated"`
	// The error code returned if the request fails.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request fails.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**: The request is successful.
	//
	// 	- **false**: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDataImportOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataImportOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataImportOrderResponseBody) GetCreateOrderResult() []*int64 {
	return s.CreateOrderResult
}

func (s *CreateDataImportOrderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateDataImportOrderResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateDataImportOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataImportOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDataImportOrderResponseBody) SetCreateOrderResult(v []*int64) *CreateDataImportOrderResponseBody {
	s.CreateOrderResult = v
	return s
}

func (s *CreateDataImportOrderResponseBody) SetErrorCode(v string) *CreateDataImportOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDataImportOrderResponseBody) SetErrorMessage(v string) *CreateDataImportOrderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDataImportOrderResponseBody) SetRequestId(v string) *CreateDataImportOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataImportOrderResponseBody) SetSuccess(v bool) *CreateDataImportOrderResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDataImportOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
