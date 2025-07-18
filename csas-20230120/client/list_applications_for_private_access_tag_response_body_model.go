// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsForPrivateAccessTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListApplicationsForPrivateAccessTagResponseBody
	GetRequestId() *string
	SetTags(v []*ListApplicationsForPrivateAccessTagResponseBodyTags) *ListApplicationsForPrivateAccessTagResponseBody
	GetTags() []*ListApplicationsForPrivateAccessTagResponseBodyTags
}

type ListApplicationsForPrivateAccessTagResponseBody struct {
	// example:
	//
	// B608C6AE-623D-55C4-9454-601B88AE937E
	RequestId *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tags      []*ListApplicationsForPrivateAccessTagResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListApplicationsForPrivateAccessTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForPrivateAccessTagResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationsForPrivateAccessTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationsForPrivateAccessTagResponseBody) GetTags() []*ListApplicationsForPrivateAccessTagResponseBodyTags {
	return s.Tags
}

func (s *ListApplicationsForPrivateAccessTagResponseBody) SetRequestId(v string) *ListApplicationsForPrivateAccessTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationsForPrivateAccessTagResponseBody) SetTags(v []*ListApplicationsForPrivateAccessTagResponseBodyTags) *ListApplicationsForPrivateAccessTagResponseBody {
	s.Tags = v
	return s
}

func (s *ListApplicationsForPrivateAccessTagResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListApplicationsForPrivateAccessTagResponseBodyTags struct {
	Applications []*ListApplicationsForPrivateAccessTagResponseBodyTagsApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// example:
	//
	// tag-7ffc82853476****
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s ListApplicationsForPrivateAccessTagResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForPrivateAccessTagResponseBodyTags) GoString() string {
	return s.String()
}

func (s *ListApplicationsForPrivateAccessTagResponseBodyTags) GetApplications() []*ListApplicationsForPrivateAccessTagResponseBodyTagsApplications {
	return s.Applications
}

func (s *ListApplicationsForPrivateAccessTagResponseBodyTags) GetTagId() *string {
	return s.TagId
}

func (s *ListApplicationsForPrivateAccessTagResponseBodyTags) SetApplications(v []*ListApplicationsForPrivateAccessTagResponseBodyTagsApplications) *ListApplicationsForPrivateAccessTagResponseBodyTags {
	s.Applications = v
	return s
}

func (s *ListApplicationsForPrivateAccessTagResponseBodyTags) SetTagId(v string) *ListApplicationsForPrivateAccessTagResponseBodyTags {
	s.TagId = &v
	return s
}

func (s *ListApplicationsForPrivateAccessTagResponseBodyTags) Validate() error {
	return dara.Validate(s)
}

type ListApplicationsForPrivateAccessTagResponseBodyTagsApplications struct {
	Addresses []*string `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	// example:
	//
	// pa-application-7a9243dd02f4****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// 内网访问应用创建时间。
	//
	// example:
	//
	// 2022-09-27 18:10:25
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// private_access_application_name
	Name       *string                                                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	PortRanges []*ListApplicationsForPrivateAccessTagResponseBodyTagsApplicationsPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
	// example:
	//
	// All
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListApplicationsForPrivateAccessTagResponseBodyTagsApplications) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForPrivateAccessTagResponseBodyTagsApplications) GoString() string {
	return s.String()
}

func (s *ListApplicationsForPrivateAccessTagResponseBodyTagsApplications) GetAddresses() []*string {
	return s.Addresses
}

func (s *ListApplicationsForPrivateAccessTagResponseBodyTagsApplications) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListApplicationsForPrivateAccessTagResponseBodyTagsApplications) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListApplicationsForPrivateAccessTagResponseBodyTagsApplications) GetDescription() *string {
	return s.Description
}

func (s *ListApplicationsForPrivateAccessTagResponseBodyTagsApplications) GetName() *string {
	return s.Name
}

func (s *ListApplicationsForPrivateAccessTagResponseBodyTagsApplications) GetPortRanges() []*ListApplicationsForPrivateAccessTagResponseBodyTagsApplicationsPortRanges {
	return s.PortRanges
}

func (s *ListApplicationsForPrivateAccessTagResponseBodyTagsApplications) GetProtocol() *string {
	return s.Protocol
}

func (s *ListApplicationsForPrivateAccessTagResponseBodyTagsApplications) GetStatus() *string {
	return s.Status
}

func (s *ListApplicationsForPrivateAccessTagResponseBodyTagsApplications) SetAddresses(v []*string) *ListApplicationsForPrivateAccessTagResponseBodyTagsApplications {
	s.Addresses = v
	return s
}

func (s *ListApplicationsForPrivateAccessTagResponseBodyTagsApplications) SetApplicationId(v string) *ListApplicationsForPrivateAccessTagResponseBodyTagsApplications {
	s.ApplicationId = &v
	return s
}

func (s *ListApplicationsForPrivateAccessTagResponseBodyTagsApplications) SetCreateTime(v string) *ListApplicationsForPrivateAccessTagResponseBodyTagsApplications {
	s.CreateTime = &v
	return s
}

func (s *ListApplicationsForPrivateAccessTagResponseBodyTagsApplications) SetDescription(v string) *ListApplicationsForPrivateAccessTagResponseBodyTagsApplications {
	s.Description = &v
	return s
}

func (s *ListApplicationsForPrivateAccessTagResponseBodyTagsApplications) SetName(v string) *ListApplicationsForPrivateAccessTagResponseBodyTagsApplications {
	s.Name = &v
	return s
}

func (s *ListApplicationsForPrivateAccessTagResponseBodyTagsApplications) SetPortRanges(v []*ListApplicationsForPrivateAccessTagResponseBodyTagsApplicationsPortRanges) *ListApplicationsForPrivateAccessTagResponseBodyTagsApplications {
	s.PortRanges = v
	return s
}

func (s *ListApplicationsForPrivateAccessTagResponseBodyTagsApplications) SetProtocol(v string) *ListApplicationsForPrivateAccessTagResponseBodyTagsApplications {
	s.Protocol = &v
	return s
}

func (s *ListApplicationsForPrivateAccessTagResponseBodyTagsApplications) SetStatus(v string) *ListApplicationsForPrivateAccessTagResponseBodyTagsApplications {
	s.Status = &v
	return s
}

func (s *ListApplicationsForPrivateAccessTagResponseBodyTagsApplications) Validate() error {
	return dara.Validate(s)
}

type ListApplicationsForPrivateAccessTagResponseBodyTagsApplicationsPortRanges struct {
	// example:
	//
	// 80
	Begin *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	// example:
	//
	// 81
	End *int32 `json:"End,omitempty" xml:"End,omitempty"`
}

func (s ListApplicationsForPrivateAccessTagResponseBodyTagsApplicationsPortRanges) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForPrivateAccessTagResponseBodyTagsApplicationsPortRanges) GoString() string {
	return s.String()
}

func (s *ListApplicationsForPrivateAccessTagResponseBodyTagsApplicationsPortRanges) GetBegin() *int32 {
	return s.Begin
}

func (s *ListApplicationsForPrivateAccessTagResponseBodyTagsApplicationsPortRanges) GetEnd() *int32 {
	return s.End
}

func (s *ListApplicationsForPrivateAccessTagResponseBodyTagsApplicationsPortRanges) SetBegin(v int32) *ListApplicationsForPrivateAccessTagResponseBodyTagsApplicationsPortRanges {
	s.Begin = &v
	return s
}

func (s *ListApplicationsForPrivateAccessTagResponseBodyTagsApplicationsPortRanges) SetEnd(v int32) *ListApplicationsForPrivateAccessTagResponseBodyTagsApplicationsPortRanges {
	s.End = &v
	return s
}

func (s *ListApplicationsForPrivateAccessTagResponseBodyTagsApplicationsPortRanges) Validate() error {
	return dara.Validate(s)
}
