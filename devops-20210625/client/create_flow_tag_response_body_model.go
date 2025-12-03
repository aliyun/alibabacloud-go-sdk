// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlowTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateFlowTagResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateFlowTagResponseBody
	GetErrorMessage() *string
	SetId(v int64) *CreateFlowTagResponseBody
	GetId() *int64
	SetRequestId(v string) *CreateFlowTagResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateFlowTagResponseBody
	GetSuccess() *bool
}

type CreateFlowTagResponseBody struct {
	// example:
	//
	// ”“
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ”“
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 111
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateFlowTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowTagResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFlowTagResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateFlowTagResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateFlowTagResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateFlowTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFlowTagResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateFlowTagResponseBody) SetErrorCode(v string) *CreateFlowTagResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateFlowTagResponseBody) SetErrorMessage(v string) *CreateFlowTagResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateFlowTagResponseBody) SetId(v int64) *CreateFlowTagResponseBody {
	s.Id = &v
	return s
}

func (s *CreateFlowTagResponseBody) SetRequestId(v string) *CreateFlowTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFlowTagResponseBody) SetSuccess(v bool) *CreateFlowTagResponseBody {
	s.Success = &v
	return s
}

func (s *CreateFlowTagResponseBody) Validate() error {
	return dara.Validate(s)
}
