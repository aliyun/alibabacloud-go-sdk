// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateCallCenterForPartnerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *OperateCallCenterForPartnerResponseBody
	GetData() *string
	SetErrorCode(v string) *OperateCallCenterForPartnerResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *OperateCallCenterForPartnerResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *OperateCallCenterForPartnerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OperateCallCenterForPartnerResponseBody
	GetSuccess() *bool
}

type OperateCallCenterForPartnerResponseBody struct {
	// example:
	//
	// 02131584182
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// NoPermission
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 1A13ABB5-7649-5031-B55C-D2E38F9F189D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OperateCallCenterForPartnerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateCallCenterForPartnerResponseBody) GoString() string {
	return s.String()
}

func (s *OperateCallCenterForPartnerResponseBody) GetData() *string {
	return s.Data
}

func (s *OperateCallCenterForPartnerResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *OperateCallCenterForPartnerResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *OperateCallCenterForPartnerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateCallCenterForPartnerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OperateCallCenterForPartnerResponseBody) SetData(v string) *OperateCallCenterForPartnerResponseBody {
	s.Data = &v
	return s
}

func (s *OperateCallCenterForPartnerResponseBody) SetErrorCode(v string) *OperateCallCenterForPartnerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *OperateCallCenterForPartnerResponseBody) SetErrorMsg(v string) *OperateCallCenterForPartnerResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *OperateCallCenterForPartnerResponseBody) SetRequestId(v string) *OperateCallCenterForPartnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateCallCenterForPartnerResponseBody) SetSuccess(v bool) *OperateCallCenterForPartnerResponseBody {
	s.Success = &v
	return s
}

func (s *OperateCallCenterForPartnerResponseBody) Validate() error {
	return dara.Validate(s)
}
