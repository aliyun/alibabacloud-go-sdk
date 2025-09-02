// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBaselineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *UpdateBaselineResponseBody
	GetData() *bool
	SetErrorCode(v string) *UpdateBaselineResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateBaselineResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *UpdateBaselineResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *UpdateBaselineResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateBaselineResponseBody
	GetSuccess() *bool
}

type UpdateBaselineResponseBody struct {
	// Indicates whether the baseline was updated.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// 6E07E90B-D9BC-5D6B-896A-82BA41A34AE1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateBaselineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateBaselineResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBaselineResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateBaselineResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateBaselineResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateBaselineResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateBaselineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateBaselineResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateBaselineResponseBody) SetData(v bool) *UpdateBaselineResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateBaselineResponseBody) SetErrorCode(v string) *UpdateBaselineResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateBaselineResponseBody) SetErrorMessage(v string) *UpdateBaselineResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateBaselineResponseBody) SetHttpStatusCode(v int32) *UpdateBaselineResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateBaselineResponseBody) SetRequestId(v string) *UpdateBaselineResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateBaselineResponseBody) SetSuccess(v bool) *UpdateBaselineResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateBaselineResponseBody) Validate() error {
	return dara.Validate(s)
}
