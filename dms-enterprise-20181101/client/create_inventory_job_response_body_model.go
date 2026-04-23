// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInventoryJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *CreateInventoryJobResponseBody
	GetData() *int64
	SetErrorCode(v string) *CreateInventoryJobResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateInventoryJobResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateInventoryJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateInventoryJobResponseBody
	GetSuccess() *bool
}

type CreateInventoryJobResponseBody struct {
	// example:
	//
	// 1001
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateInventoryJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInventoryJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInventoryJobResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateInventoryJobResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateInventoryJobResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateInventoryJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInventoryJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateInventoryJobResponseBody) SetData(v int64) *CreateInventoryJobResponseBody {
	s.Data = &v
	return s
}

func (s *CreateInventoryJobResponseBody) SetErrorCode(v string) *CreateInventoryJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateInventoryJobResponseBody) SetErrorMessage(v string) *CreateInventoryJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateInventoryJobResponseBody) SetRequestId(v string) *CreateInventoryJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInventoryJobResponseBody) SetSuccess(v bool) *CreateInventoryJobResponseBody {
	s.Success = &v
	return s
}

func (s *CreateInventoryJobResponseBody) Validate() error {
	return dara.Validate(s)
}
