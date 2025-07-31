// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *OperateInstanceResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *OperateInstanceResponseBody
	GetHttpStatusCode() *int32
	SetInstanceStatusList(v []*OperateInstanceResponseBodyInstanceStatusList) *OperateInstanceResponseBody
	GetInstanceStatusList() []*OperateInstanceResponseBodyInstanceStatusList
	SetMessage(v string) *OperateInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *OperateInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OperateInstanceResponseBody
	GetSuccess() *bool
}

type OperateInstanceResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode     *int32                                           `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	InstanceStatusList []*OperateInstanceResponseBodyInstanceStatusList `json:"InstanceStatusList,omitempty" xml:"InstanceStatusList,omitempty" type:"Repeated"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OperateInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *OperateInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *OperateInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *OperateInstanceResponseBody) GetInstanceStatusList() []*OperateInstanceResponseBodyInstanceStatusList {
	return s.InstanceStatusList
}

func (s *OperateInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OperateInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OperateInstanceResponseBody) SetCode(v string) *OperateInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *OperateInstanceResponseBody) SetHttpStatusCode(v int32) *OperateInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *OperateInstanceResponseBody) SetInstanceStatusList(v []*OperateInstanceResponseBodyInstanceStatusList) *OperateInstanceResponseBody {
	s.InstanceStatusList = v
	return s
}

func (s *OperateInstanceResponseBody) SetMessage(v string) *OperateInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *OperateInstanceResponseBody) SetRequestId(v string) *OperateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateInstanceResponseBody) SetSuccess(v bool) *OperateInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *OperateInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type OperateInstanceResponseBodyInstanceStatusList struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// xx
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// t_132435
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 121311
	OwnerId   *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s OperateInstanceResponseBodyInstanceStatusList) String() string {
	return dara.Prettify(s)
}

func (s OperateInstanceResponseBodyInstanceStatusList) GoString() string {
	return s.String()
}

func (s *OperateInstanceResponseBodyInstanceStatusList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *OperateInstanceResponseBodyInstanceStatusList) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *OperateInstanceResponseBodyInstanceStatusList) GetId() *string {
	return s.Id
}

func (s *OperateInstanceResponseBodyInstanceStatusList) GetName() *string {
	return s.Name
}

func (s *OperateInstanceResponseBodyInstanceStatusList) GetOwnerId() *string {
	return s.OwnerId
}

func (s *OperateInstanceResponseBodyInstanceStatusList) GetOwnerName() *string {
	return s.OwnerName
}

func (s *OperateInstanceResponseBodyInstanceStatusList) GetStatus() *string {
	return s.Status
}

func (s *OperateInstanceResponseBodyInstanceStatusList) SetDisplayName(v string) *OperateInstanceResponseBodyInstanceStatusList {
	s.DisplayName = &v
	return s
}

func (s *OperateInstanceResponseBodyInstanceStatusList) SetErrorMessage(v string) *OperateInstanceResponseBodyInstanceStatusList {
	s.ErrorMessage = &v
	return s
}

func (s *OperateInstanceResponseBodyInstanceStatusList) SetId(v string) *OperateInstanceResponseBodyInstanceStatusList {
	s.Id = &v
	return s
}

func (s *OperateInstanceResponseBodyInstanceStatusList) SetName(v string) *OperateInstanceResponseBodyInstanceStatusList {
	s.Name = &v
	return s
}

func (s *OperateInstanceResponseBodyInstanceStatusList) SetOwnerId(v string) *OperateInstanceResponseBodyInstanceStatusList {
	s.OwnerId = &v
	return s
}

func (s *OperateInstanceResponseBodyInstanceStatusList) SetOwnerName(v string) *OperateInstanceResponseBodyInstanceStatusList {
	s.OwnerName = &v
	return s
}

func (s *OperateInstanceResponseBodyInstanceStatusList) SetStatus(v string) *OperateInstanceResponseBodyInstanceStatusList {
	s.Status = &v
	return s
}

func (s *OperateInstanceResponseBodyInstanceStatusList) Validate() error {
	return dara.Validate(s)
}
