// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBaselineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *CreateBaselineResponseBody
	GetData() *int64
	SetErrorCode(v string) *CreateBaselineResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateBaselineResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *CreateBaselineResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CreateBaselineResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateBaselineResponseBody
	GetSuccess() *bool
}

type CreateBaselineResponseBody struct {
	// The baseline ID.
	//
	// example:
	//
	// 100003
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// 401
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified parameters are invalid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F05080B0-CCE6-5D22-B284-34A51C5D4E28
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

func (s CreateBaselineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBaselineResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBaselineResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateBaselineResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateBaselineResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateBaselineResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateBaselineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBaselineResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateBaselineResponseBody) SetData(v int64) *CreateBaselineResponseBody {
	s.Data = &v
	return s
}

func (s *CreateBaselineResponseBody) SetErrorCode(v string) *CreateBaselineResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateBaselineResponseBody) SetErrorMessage(v string) *CreateBaselineResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateBaselineResponseBody) SetHttpStatusCode(v int32) *CreateBaselineResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateBaselineResponseBody) SetRequestId(v string) *CreateBaselineResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBaselineResponseBody) SetSuccess(v bool) *CreateBaselineResponseBody {
	s.Success = &v
	return s
}

func (s *CreateBaselineResponseBody) Validate() error {
	return dara.Validate(s)
}
