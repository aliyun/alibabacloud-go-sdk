// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBaselineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteBaselineResponseBody
	GetData() *bool
	SetErrorCode(v string) *DeleteBaselineResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteBaselineResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DeleteBaselineResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DeleteBaselineResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteBaselineResponseBody
	GetSuccess() *bool
}

type DeleteBaselineResponseBody struct {
	// Indicates whether the deletion was successful.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// 1031203110000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Baseline deletion failed with nodes dependent on baseline
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
	// 0000-ABCD-EF****
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

func (s DeleteBaselineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBaselineResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBaselineResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteBaselineResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteBaselineResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteBaselineResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteBaselineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBaselineResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteBaselineResponseBody) SetData(v bool) *DeleteBaselineResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteBaselineResponseBody) SetErrorCode(v string) *DeleteBaselineResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteBaselineResponseBody) SetErrorMessage(v string) *DeleteBaselineResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteBaselineResponseBody) SetHttpStatusCode(v int32) *DeleteBaselineResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteBaselineResponseBody) SetRequestId(v string) *DeleteBaselineResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBaselineResponseBody) SetSuccess(v bool) *DeleteBaselineResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteBaselineResponseBody) Validate() error {
	return dara.Validate(s)
}
