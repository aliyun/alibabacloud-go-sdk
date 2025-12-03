// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlowTagGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateFlowTagGroupResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateFlowTagGroupResponseBody
	GetErrorMessage() *string
	SetId(v int64) *CreateFlowTagGroupResponseBody
	GetId() *int64
	SetRequestId(v string) *CreateFlowTagGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateFlowTagGroupResponseBody
	GetSuccess() *bool
}

type CreateFlowTagGroupResponseBody struct {
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
	// 1223
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

func (s CreateFlowTagGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowTagGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFlowTagGroupResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateFlowTagGroupResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateFlowTagGroupResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateFlowTagGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFlowTagGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateFlowTagGroupResponseBody) SetErrorCode(v string) *CreateFlowTagGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateFlowTagGroupResponseBody) SetErrorMessage(v string) *CreateFlowTagGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateFlowTagGroupResponseBody) SetId(v int64) *CreateFlowTagGroupResponseBody {
	s.Id = &v
	return s
}

func (s *CreateFlowTagGroupResponseBody) SetRequestId(v string) *CreateFlowTagGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFlowTagGroupResponseBody) SetSuccess(v bool) *CreateFlowTagGroupResponseBody {
	s.Success = &v
	return s
}

func (s *CreateFlowTagGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
