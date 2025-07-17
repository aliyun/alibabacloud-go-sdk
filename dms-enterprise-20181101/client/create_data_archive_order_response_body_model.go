// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataArchiveOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateOrderResult(v []*int64) *CreateDataArchiveOrderResponseBody
	GetCreateOrderResult() []*int64
	SetErrorCode(v string) *CreateDataArchiveOrderResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateDataArchiveOrderResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateDataArchiveOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDataArchiveOrderResponseBody
	GetSuccess() *bool
}

type CreateDataArchiveOrderResponseBody struct {
	// The ID of the data archiving ticket.
	CreateOrderResult []*int64 `json:"CreateOrderResult,omitempty" xml:"CreateOrderResult,omitempty" type:"Repeated"`
	// The error code returned if the request failed.
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
	// The ID of the request, which is used to query logs and troubleshoot issues.
	//
	// example:
	//
	// 283C461F-11D8-48AA-B695-DF092DA32AF3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDataArchiveOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataArchiveOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataArchiveOrderResponseBody) GetCreateOrderResult() []*int64 {
	return s.CreateOrderResult
}

func (s *CreateDataArchiveOrderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateDataArchiveOrderResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateDataArchiveOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataArchiveOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDataArchiveOrderResponseBody) SetCreateOrderResult(v []*int64) *CreateDataArchiveOrderResponseBody {
	s.CreateOrderResult = v
	return s
}

func (s *CreateDataArchiveOrderResponseBody) SetErrorCode(v string) *CreateDataArchiveOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDataArchiveOrderResponseBody) SetErrorMessage(v string) *CreateDataArchiveOrderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDataArchiveOrderResponseBody) SetRequestId(v string) *CreateDataArchiveOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataArchiveOrderResponseBody) SetSuccess(v bool) *CreateDataArchiveOrderResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDataArchiveOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
