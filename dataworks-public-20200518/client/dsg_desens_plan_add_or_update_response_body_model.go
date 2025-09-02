// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgDesensPlanAddOrUpdateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DsgDesensPlanAddOrUpdateResponseBody
	GetData() *bool
	SetErrorCode(v string) *DsgDesensPlanAddOrUpdateResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DsgDesensPlanAddOrUpdateResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DsgDesensPlanAddOrUpdateResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DsgDesensPlanAddOrUpdateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DsgDesensPlanAddOrUpdateResponseBody
	GetSuccess() *bool
}

type DsgDesensPlanAddOrUpdateResponseBody struct {
	// The execution result of adding or modifying a data masking rule.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// 1029030003
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// param error
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID. You can locate logs and troubleshoot issues based on the ID.
	//
	// example:
	//
	// 102400001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DsgDesensPlanAddOrUpdateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DsgDesensPlanAddOrUpdateResponseBody) GoString() string {
	return s.String()
}

func (s *DsgDesensPlanAddOrUpdateResponseBody) GetData() *bool {
	return s.Data
}

func (s *DsgDesensPlanAddOrUpdateResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DsgDesensPlanAddOrUpdateResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DsgDesensPlanAddOrUpdateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DsgDesensPlanAddOrUpdateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DsgDesensPlanAddOrUpdateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DsgDesensPlanAddOrUpdateResponseBody) SetData(v bool) *DsgDesensPlanAddOrUpdateResponseBody {
	s.Data = &v
	return s
}

func (s *DsgDesensPlanAddOrUpdateResponseBody) SetErrorCode(v string) *DsgDesensPlanAddOrUpdateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DsgDesensPlanAddOrUpdateResponseBody) SetErrorMessage(v string) *DsgDesensPlanAddOrUpdateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DsgDesensPlanAddOrUpdateResponseBody) SetHttpStatusCode(v int32) *DsgDesensPlanAddOrUpdateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DsgDesensPlanAddOrUpdateResponseBody) SetRequestId(v string) *DsgDesensPlanAddOrUpdateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DsgDesensPlanAddOrUpdateResponseBody) SetSuccess(v bool) *DsgDesensPlanAddOrUpdateResponseBody {
	s.Success = &v
	return s
}

func (s *DsgDesensPlanAddOrUpdateResponseBody) Validate() error {
	return dara.Validate(s)
}
