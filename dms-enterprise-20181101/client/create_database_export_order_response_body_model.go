// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatabaseExportOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateOrderResult(v *CreateDatabaseExportOrderResponseBodyCreateOrderResult) *CreateDatabaseExportOrderResponseBody
	GetCreateOrderResult() *CreateDatabaseExportOrderResponseBodyCreateOrderResult
	SetErrorCode(v string) *CreateDatabaseExportOrderResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateDatabaseExportOrderResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateDatabaseExportOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDatabaseExportOrderResponseBody
	GetSuccess() *bool
}

type CreateDatabaseExportOrderResponseBody struct {
	// The information about the ticket.
	CreateOrderResult *CreateDatabaseExportOrderResponseBodyCreateOrderResult `json:"CreateOrderResult,omitempty" xml:"CreateOrderResult,omitempty" type:"Struct"`
	// The error code returned.
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
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
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

func (s CreateDatabaseExportOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDatabaseExportOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDatabaseExportOrderResponseBody) GetCreateOrderResult() *CreateDatabaseExportOrderResponseBodyCreateOrderResult {
	return s.CreateOrderResult
}

func (s *CreateDatabaseExportOrderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateDatabaseExportOrderResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateDatabaseExportOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDatabaseExportOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDatabaseExportOrderResponseBody) SetCreateOrderResult(v *CreateDatabaseExportOrderResponseBodyCreateOrderResult) *CreateDatabaseExportOrderResponseBody {
	s.CreateOrderResult = v
	return s
}

func (s *CreateDatabaseExportOrderResponseBody) SetErrorCode(v string) *CreateDatabaseExportOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDatabaseExportOrderResponseBody) SetErrorMessage(v string) *CreateDatabaseExportOrderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDatabaseExportOrderResponseBody) SetRequestId(v string) *CreateDatabaseExportOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDatabaseExportOrderResponseBody) SetSuccess(v bool) *CreateDatabaseExportOrderResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDatabaseExportOrderResponseBody) Validate() error {
	if s.CreateOrderResult != nil {
		if err := s.CreateOrderResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDatabaseExportOrderResponseBodyCreateOrderResult struct {
	CreateOrderResult []*int64 `json:"CreateOrderResult,omitempty" xml:"CreateOrderResult,omitempty" type:"Repeated"`
}

func (s CreateDatabaseExportOrderResponseBodyCreateOrderResult) String() string {
	return dara.Prettify(s)
}

func (s CreateDatabaseExportOrderResponseBodyCreateOrderResult) GoString() string {
	return s.String()
}

func (s *CreateDatabaseExportOrderResponseBodyCreateOrderResult) GetCreateOrderResult() []*int64 {
	return s.CreateOrderResult
}

func (s *CreateDatabaseExportOrderResponseBodyCreateOrderResult) SetCreateOrderResult(v []*int64) *CreateDatabaseExportOrderResponseBodyCreateOrderResult {
	s.CreateOrderResult = v
	return s
}

func (s *CreateDatabaseExportOrderResponseBodyCreateOrderResult) Validate() error {
	return dara.Validate(s)
}
