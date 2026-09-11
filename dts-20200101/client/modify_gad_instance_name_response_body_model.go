// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGadInstanceNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v int64) *ModifyGadInstanceNameResponseBody
	GetCreateTime() *int64
	SetDynamicCode(v string) *ModifyGadInstanceNameResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ModifyGadInstanceNameResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *ModifyGadInstanceNameResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifyGadInstanceNameResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v string) *ModifyGadInstanceNameResponseBody
	GetHttpStatusCode() *string
	SetInstanceId(v string) *ModifyGadInstanceNameResponseBody
	GetInstanceId() *string
	SetInstanceName(v string) *ModifyGadInstanceNameResponseBody
	GetInstanceName() *string
	SetRegionId(v string) *ModifyGadInstanceNameResponseBody
	GetRegionId() *string
	SetRequestId(v string) *ModifyGadInstanceNameResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *ModifyGadInstanceNameResponseBody
	GetResourceGroupId() *string
	SetSuccess(v string) *ModifyGadInstanceNameResponseBody
	GetSuccess() *string
}

type ModifyGadInstanceNameResponseBody struct {
	// The time when the task was created, in the format of <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC).
	//
	// example:
	//
	// 2021-06-28 17:34:53.0
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The dynamic error code. This parameter will be deprecated.
	//
	// example:
	//
	// 403
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message, which replaces **%s*	- in the **ErrMessage*	- response parameter.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// present environment is not support,so skip.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code returned if the call failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the call failed.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The active geo-redundancy instance group ID.
	//
	// example:
	//
	// rm-bp162d4tp0500****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// rm-2zehh163694qs5c3v
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The region ID of the instance. For more information, see Supported regions.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4D0ADAD5-DD97-41B6-B78F-D1961AB1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyGadInstanceNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyGadInstanceNameResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyGadInstanceNameResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ModifyGadInstanceNameResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ModifyGadInstanceNameResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ModifyGadInstanceNameResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifyGadInstanceNameResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifyGadInstanceNameResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *ModifyGadInstanceNameResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyGadInstanceNameResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyGadInstanceNameResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyGadInstanceNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyGadInstanceNameResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyGadInstanceNameResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ModifyGadInstanceNameResponseBody) SetCreateTime(v int64) *ModifyGadInstanceNameResponseBody {
	s.CreateTime = &v
	return s
}

func (s *ModifyGadInstanceNameResponseBody) SetDynamicCode(v string) *ModifyGadInstanceNameResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ModifyGadInstanceNameResponseBody) SetDynamicMessage(v string) *ModifyGadInstanceNameResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ModifyGadInstanceNameResponseBody) SetErrCode(v string) *ModifyGadInstanceNameResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyGadInstanceNameResponseBody) SetErrMessage(v string) *ModifyGadInstanceNameResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyGadInstanceNameResponseBody) SetHttpStatusCode(v string) *ModifyGadInstanceNameResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyGadInstanceNameResponseBody) SetInstanceId(v string) *ModifyGadInstanceNameResponseBody {
	s.InstanceId = &v
	return s
}

func (s *ModifyGadInstanceNameResponseBody) SetInstanceName(v string) *ModifyGadInstanceNameResponseBody {
	s.InstanceName = &v
	return s
}

func (s *ModifyGadInstanceNameResponseBody) SetRegionId(v string) *ModifyGadInstanceNameResponseBody {
	s.RegionId = &v
	return s
}

func (s *ModifyGadInstanceNameResponseBody) SetRequestId(v string) *ModifyGadInstanceNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyGadInstanceNameResponseBody) SetResourceGroupId(v string) *ModifyGadInstanceNameResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyGadInstanceNameResponseBody) SetSuccess(v string) *ModifyGadInstanceNameResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyGadInstanceNameResponseBody) Validate() error {
	return dara.Validate(s)
}
