// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UploadDataResponseBody
	GetCode() *string
	SetData(v string) *UploadDataResponseBody
	GetData() *string
	SetMessage(v string) *UploadDataResponseBody
	GetMessage() *string
	SetRequestId(v string) *UploadDataResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UploadDataResponseBody
	GetSuccess() *bool
}

type UploadDataResponseBody struct {
	// Result code. **200*	- means success. Any other value means failure. Use this field to diagnose failures.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Task ID for this request.
	//
	// example:
	//
	// 6F5934C7-C223-4F0F-BBF3-5B3594***
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// Error details if the request failed. Returns successful if the request succeeded.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID. A unique identifier for this request. Use it to trace the request.
	//
	// example:
	//
	// 6F5934C7-C223-4F0F-BBF3-5B3594****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request succeeded. true means success. false or null means failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UploadDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadDataResponseBody) GoString() string {
	return s.String()
}

func (s *UploadDataResponseBody) GetCode() *string {
	return s.Code
}

func (s *UploadDataResponseBody) GetData() *string {
	return s.Data
}

func (s *UploadDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UploadDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadDataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UploadDataResponseBody) SetCode(v string) *UploadDataResponseBody {
	s.Code = &v
	return s
}

func (s *UploadDataResponseBody) SetData(v string) *UploadDataResponseBody {
	s.Data = &v
	return s
}

func (s *UploadDataResponseBody) SetMessage(v string) *UploadDataResponseBody {
	s.Message = &v
	return s
}

func (s *UploadDataResponseBody) SetRequestId(v string) *UploadDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadDataResponseBody) SetSuccess(v bool) *UploadDataResponseBody {
	s.Success = &v
	return s
}

func (s *UploadDataResponseBody) Validate() error {
	return dara.Validate(s)
}
