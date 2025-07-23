// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateAppInstanceForPartnerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *OperateAppInstanceForPartnerResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *OperateAppInstanceForPartnerResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *OperateAppInstanceForPartnerResponseBody
	GetRequestId() *string
	SetSuccess(v string) *OperateAppInstanceForPartnerResponseBody
	GetSuccess() *string
}

type OperateAppInstanceForPartnerResponseBody struct {
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
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OperateAppInstanceForPartnerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateAppInstanceForPartnerResponseBody) GoString() string {
	return s.String()
}

func (s *OperateAppInstanceForPartnerResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *OperateAppInstanceForPartnerResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *OperateAppInstanceForPartnerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateAppInstanceForPartnerResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *OperateAppInstanceForPartnerResponseBody) SetErrorCode(v string) *OperateAppInstanceForPartnerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *OperateAppInstanceForPartnerResponseBody) SetErrorMsg(v string) *OperateAppInstanceForPartnerResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *OperateAppInstanceForPartnerResponseBody) SetRequestId(v string) *OperateAppInstanceForPartnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateAppInstanceForPartnerResponseBody) SetSuccess(v string) *OperateAppInstanceForPartnerResponseBody {
	s.Success = &v
	return s
}

func (s *OperateAppInstanceForPartnerResponseBody) Validate() error {
	return dara.Validate(s)
}
