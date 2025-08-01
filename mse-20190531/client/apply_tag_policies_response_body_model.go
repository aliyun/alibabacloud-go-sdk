// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyTagPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ApplyTagPoliciesResponseBodyData) *ApplyTagPoliciesResponseBody
	GetData() []*ApplyTagPoliciesResponseBodyData
	SetMessage(v string) *ApplyTagPoliciesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ApplyTagPoliciesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ApplyTagPoliciesResponseBody
	GetSuccess() *bool
}

type ApplyTagPoliciesResponseBody struct {
	// The details of the data.
	Data []*ApplyTagPoliciesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The message returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 316F5F64-F73D-42DC-8632-01E308B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ApplyTagPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyTagPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyTagPoliciesResponseBody) GetData() []*ApplyTagPoliciesResponseBodyData {
	return s.Data
}

func (s *ApplyTagPoliciesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ApplyTagPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyTagPoliciesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ApplyTagPoliciesResponseBody) SetData(v []*ApplyTagPoliciesResponseBodyData) *ApplyTagPoliciesResponseBody {
	s.Data = v
	return s
}

func (s *ApplyTagPoliciesResponseBody) SetMessage(v string) *ApplyTagPoliciesResponseBody {
	s.Message = &v
	return s
}

func (s *ApplyTagPoliciesResponseBody) SetRequestId(v string) *ApplyTagPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyTagPoliciesResponseBody) SetSuccess(v bool) *ApplyTagPoliciesResponseBody {
	s.Success = &v
	return s
}

func (s *ApplyTagPoliciesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ApplyTagPoliciesResponseBodyData struct {
	// Indicates whether the field is the primary key.
	//
	// example:
	//
	// true
	CarryData *bool `json:"CarryData,omitempty" xml:"CarryData,omitempty"`
	// Indicates whether the rule is enabled. Valid values:
	//
	// 	- `true`: The rule is enabled.
	//
	// 	- `false`: The rule is disabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The ID of the primary key.
	//
	// example:
	//
	// 12
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The number of instances.
	//
	// example:
	//
	// 1
	InstanceNum *int32 `json:"InstanceNum,omitempty" xml:"InstanceNum,omitempty"`
	// The policy name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The rate.
	//
	// example:
	//
	// 10
	Rate *int32 `json:"Rate,omitempty" xml:"Rate,omitempty"`
	// Indicates whether the routing rule was deleted.
	Remove *bool `json:"Remove,omitempty" xml:"Remove,omitempty"`
	// The details of the routing rule.
	//
	// example:
	//
	// {\\"_base\\": {\\"rate\\": 100, \\"remove\\": true}, \\"blue\\": {\\"rate\\": 0}}
	Rules *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
	// The status.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag.
	//
	// example:
	//
	// gray
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ApplyTagPoliciesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ApplyTagPoliciesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ApplyTagPoliciesResponseBodyData) GetCarryData() *bool {
	return s.CarryData
}

func (s *ApplyTagPoliciesResponseBodyData) GetEnable() *bool {
	return s.Enable
}

func (s *ApplyTagPoliciesResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ApplyTagPoliciesResponseBodyData) GetInstanceNum() *int32 {
	return s.InstanceNum
}

func (s *ApplyTagPoliciesResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ApplyTagPoliciesResponseBodyData) GetRate() *int32 {
	return s.Rate
}

func (s *ApplyTagPoliciesResponseBodyData) GetRemove() *bool {
	return s.Remove
}

func (s *ApplyTagPoliciesResponseBodyData) GetRules() *string {
	return s.Rules
}

func (s *ApplyTagPoliciesResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *ApplyTagPoliciesResponseBodyData) GetTag() *string {
	return s.Tag
}

func (s *ApplyTagPoliciesResponseBodyData) SetCarryData(v bool) *ApplyTagPoliciesResponseBodyData {
	s.CarryData = &v
	return s
}

func (s *ApplyTagPoliciesResponseBodyData) SetEnable(v bool) *ApplyTagPoliciesResponseBodyData {
	s.Enable = &v
	return s
}

func (s *ApplyTagPoliciesResponseBodyData) SetId(v int64) *ApplyTagPoliciesResponseBodyData {
	s.Id = &v
	return s
}

func (s *ApplyTagPoliciesResponseBodyData) SetInstanceNum(v int32) *ApplyTagPoliciesResponseBodyData {
	s.InstanceNum = &v
	return s
}

func (s *ApplyTagPoliciesResponseBodyData) SetName(v string) *ApplyTagPoliciesResponseBodyData {
	s.Name = &v
	return s
}

func (s *ApplyTagPoliciesResponseBodyData) SetRate(v int32) *ApplyTagPoliciesResponseBodyData {
	s.Rate = &v
	return s
}

func (s *ApplyTagPoliciesResponseBodyData) SetRemove(v bool) *ApplyTagPoliciesResponseBodyData {
	s.Remove = &v
	return s
}

func (s *ApplyTagPoliciesResponseBodyData) SetRules(v string) *ApplyTagPoliciesResponseBodyData {
	s.Rules = &v
	return s
}

func (s *ApplyTagPoliciesResponseBodyData) SetStatus(v int32) *ApplyTagPoliciesResponseBodyData {
	s.Status = &v
	return s
}

func (s *ApplyTagPoliciesResponseBodyData) SetTag(v string) *ApplyTagPoliciesResponseBodyData {
	s.Tag = &v
	return s
}

func (s *ApplyTagPoliciesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
