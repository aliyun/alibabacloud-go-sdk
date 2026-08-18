// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterGetDeptBalanceSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeptBalanceSummaryDTO) *ModelRouterGetDeptBalanceSummaryResponseBody
	GetData() *DeptBalanceSummaryDTO
	SetErrCode(v string) *ModelRouterGetDeptBalanceSummaryResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterGetDeptBalanceSummaryResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterGetDeptBalanceSummaryResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterGetDeptBalanceSummaryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterGetDeptBalanceSummaryResponseBody
	GetSuccess() *bool
}

type ModelRouterGetDeptBalanceSummaryResponseBody struct {
	// The data object.
	//
	// example:
	//
	// {}
	Data *DeptBalanceSummaryDTO `json:"data,omitempty" xml:"data,omitempty"`
	// The fault message code.
	//
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Unknown error
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterGetDeptBalanceSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterGetDeptBalanceSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterGetDeptBalanceSummaryResponseBody) GetData() *DeptBalanceSummaryDTO {
	return s.Data
}

func (s *ModelRouterGetDeptBalanceSummaryResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterGetDeptBalanceSummaryResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterGetDeptBalanceSummaryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterGetDeptBalanceSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterGetDeptBalanceSummaryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterGetDeptBalanceSummaryResponseBody) SetData(v *DeptBalanceSummaryDTO) *ModelRouterGetDeptBalanceSummaryResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterGetDeptBalanceSummaryResponseBody) SetErrCode(v string) *ModelRouterGetDeptBalanceSummaryResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterGetDeptBalanceSummaryResponseBody) SetErrMessage(v string) *ModelRouterGetDeptBalanceSummaryResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterGetDeptBalanceSummaryResponseBody) SetHttpStatusCode(v int32) *ModelRouterGetDeptBalanceSummaryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterGetDeptBalanceSummaryResponseBody) SetRequestId(v string) *ModelRouterGetDeptBalanceSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterGetDeptBalanceSummaryResponseBody) SetSuccess(v bool) *ModelRouterGetDeptBalanceSummaryResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterGetDeptBalanceSummaryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
