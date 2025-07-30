// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachGadInstanceDbMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v int64) *DetachGadInstanceDbMemberResponseBody
	GetCreateTime() *int64
	SetDynamicCode(v string) *DetachGadInstanceDbMemberResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DetachGadInstanceDbMemberResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *DetachGadInstanceDbMemberResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DetachGadInstanceDbMemberResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v string) *DetachGadInstanceDbMemberResponseBody
	GetHttpStatusCode() *string
	SetInstanceId(v string) *DetachGadInstanceDbMemberResponseBody
	GetInstanceId() *string
	SetInstanceName(v string) *DetachGadInstanceDbMemberResponseBody
	GetInstanceName() *string
	SetRegionId(v string) *DetachGadInstanceDbMemberResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DetachGadInstanceDbMemberResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *DetachGadInstanceDbMemberResponseBody
	GetResourceGroupId() *string
	SetSlaveDbInstanceId(v string) *DetachGadInstanceDbMemberResponseBody
	GetSlaveDbInstanceId() *string
	SetSuccess(v string) *DetachGadInstanceDbMemberResponseBody
	GetSuccess() *string
}

type DetachGadInstanceDbMemberResponseBody struct {
	// example:
	//
	// 2021-06-28 17:34:53.0
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 403
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// ****
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// gad-bp1i99e8l7913****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 6063641E-BAD1-4BA7-B70B-26FFFD18****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// rm-bp1i99e8l7913****
	SlaveDbInstanceId *string `json:"SlaveDbInstanceId,omitempty" xml:"SlaveDbInstanceId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DetachGadInstanceDbMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachGadInstanceDbMemberResponseBody) GoString() string {
	return s.String()
}

func (s *DetachGadInstanceDbMemberResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DetachGadInstanceDbMemberResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DetachGadInstanceDbMemberResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DetachGadInstanceDbMemberResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DetachGadInstanceDbMemberResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DetachGadInstanceDbMemberResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DetachGadInstanceDbMemberResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DetachGadInstanceDbMemberResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DetachGadInstanceDbMemberResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachGadInstanceDbMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachGadInstanceDbMemberResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DetachGadInstanceDbMemberResponseBody) GetSlaveDbInstanceId() *string {
	return s.SlaveDbInstanceId
}

func (s *DetachGadInstanceDbMemberResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DetachGadInstanceDbMemberResponseBody) SetCreateTime(v int64) *DetachGadInstanceDbMemberResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DetachGadInstanceDbMemberResponseBody) SetDynamicCode(v string) *DetachGadInstanceDbMemberResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DetachGadInstanceDbMemberResponseBody) SetDynamicMessage(v string) *DetachGadInstanceDbMemberResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DetachGadInstanceDbMemberResponseBody) SetErrCode(v string) *DetachGadInstanceDbMemberResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DetachGadInstanceDbMemberResponseBody) SetErrMessage(v string) *DetachGadInstanceDbMemberResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DetachGadInstanceDbMemberResponseBody) SetHttpStatusCode(v string) *DetachGadInstanceDbMemberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DetachGadInstanceDbMemberResponseBody) SetInstanceId(v string) *DetachGadInstanceDbMemberResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DetachGadInstanceDbMemberResponseBody) SetInstanceName(v string) *DetachGadInstanceDbMemberResponseBody {
	s.InstanceName = &v
	return s
}

func (s *DetachGadInstanceDbMemberResponseBody) SetRegionId(v string) *DetachGadInstanceDbMemberResponseBody {
	s.RegionId = &v
	return s
}

func (s *DetachGadInstanceDbMemberResponseBody) SetRequestId(v string) *DetachGadInstanceDbMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachGadInstanceDbMemberResponseBody) SetResourceGroupId(v string) *DetachGadInstanceDbMemberResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DetachGadInstanceDbMemberResponseBody) SetSlaveDbInstanceId(v string) *DetachGadInstanceDbMemberResponseBody {
	s.SlaveDbInstanceId = &v
	return s
}

func (s *DetachGadInstanceDbMemberResponseBody) SetSuccess(v string) *DetachGadInstanceDbMemberResponseBody {
	s.Success = &v
	return s
}

func (s *DetachGadInstanceDbMemberResponseBody) Validate() error {
	return dara.Validate(s)
}
