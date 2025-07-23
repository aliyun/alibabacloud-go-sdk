// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateAppServiceForPartnerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *OperateAppServiceForPartnerResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *OperateAppServiceForPartnerResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *OperateAppServiceForPartnerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OperateAppServiceForPartnerResponseBody
	GetSuccess() *bool
}

type OperateAppServiceForPartnerResponseBody struct {
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// aliuid:1998006665794443 assumeRole not exist,serviceName:realtimelogpush.dcdnservices.aliyuncs.com
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OperateAppServiceForPartnerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateAppServiceForPartnerResponseBody) GoString() string {
	return s.String()
}

func (s *OperateAppServiceForPartnerResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *OperateAppServiceForPartnerResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *OperateAppServiceForPartnerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateAppServiceForPartnerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OperateAppServiceForPartnerResponseBody) SetErrorCode(v string) *OperateAppServiceForPartnerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *OperateAppServiceForPartnerResponseBody) SetErrorMsg(v string) *OperateAppServiceForPartnerResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *OperateAppServiceForPartnerResponseBody) SetRequestId(v string) *OperateAppServiceForPartnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateAppServiceForPartnerResponseBody) SetSuccess(v bool) *OperateAppServiceForPartnerResponseBody {
	s.Success = &v
	return s
}

func (s *OperateAppServiceForPartnerResponseBody) Validate() error {
	return dara.Validate(s)
}
