// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSkillsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSkillsResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeSkillsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSkillsResponseBody
	GetRequestId() *string
	SetSkillInfo(v []*DescribeSkillsResponseBodySkillInfo) *DescribeSkillsResponseBody
	GetSkillInfo() []*DescribeSkillsResponseBodySkillInfo
	SetTotalCount(v string) *DescribeSkillsResponseBody
	GetTotalCount() *string
}

type DescribeSkillsResponseBody struct {
	// The status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// Success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1A923337-44D9-5CAD-9A53-95084BD4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The skill information.
	SkillInfo []*DescribeSkillsResponseBodySkillInfo `json:"SkillInfo,omitempty" xml:"SkillInfo,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 10
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSkillsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSkillsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSkillsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSkillsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSkillsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSkillsResponseBody) GetSkillInfo() []*DescribeSkillsResponseBodySkillInfo {
	return s.SkillInfo
}

func (s *DescribeSkillsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeSkillsResponseBody) SetCode(v string) *DescribeSkillsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSkillsResponseBody) SetMessage(v string) *DescribeSkillsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSkillsResponseBody) SetRequestId(v string) *DescribeSkillsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSkillsResponseBody) SetSkillInfo(v []*DescribeSkillsResponseBodySkillInfo) *DescribeSkillsResponseBody {
	s.SkillInfo = v
	return s
}

func (s *DescribeSkillsResponseBody) SetTotalCount(v string) *DescribeSkillsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSkillsResponseBody) Validate() error {
	if s.SkillInfo != nil {
		for _, item := range s.SkillInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSkillsResponseBodySkillInfo struct {
	// The skill category.
	//
	// example:
	//
	// System
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-03-13 15:40:51
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The skill description.
	//
	// example:
	//
	// Current weather and forecasts with wttr.in via curl for locations, rain, temperature, travel planning.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The OSS download URL of the skill icon.
	//
	// example:
	//
	// aHR0cDovL2Nsb3VkLXBob25lLWFpLXRlc3QwLm9zcy1jbi1oYW5nemhv****
	IconOssUrl *string `json:"IconOssUrl,omitempty" xml:"IconOssUrl,omitempty"`
	// The number of instances that have the skill installed.
	//
	// example:
	//
	// 10
	InstalledCount *int32 `json:"InstalledCount,omitempty" xml:"InstalledCount,omitempty"`
	// The information about the installed instances.
	InstalledInstances []*DescribeSkillsResponseBodySkillInfoInstalledInstances `json:"InstalledInstances,omitempty" xml:"InstalledInstances,omitempty" type:"Repeated"`
	// The skill summary.
	//
	// example:
	//
	// Current weather and forecasts.
	Instruction *string `json:"Instruction,omitempty" xml:"Instruction,omitempty"`
	// The skill ID.
	//
	// example:
	//
	// s-04zzrgosj6xd1****
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	// The skill name.
	//
	// example:
	//
	// weather
	SkillName *string `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
	// The skill source.
	//
	// example:
	//
	// USER_UPLOAD
	SkillSource *string `json:"SkillSource,omitempty" xml:"SkillSource,omitempty"`
	// The skill lifecycle status.
	//
	// example:
	//
	// UPLOADED
	SkillStatus *string `json:"SkillStatus,omitempty" xml:"SkillStatus,omitempty"`
	// The source node ID of the skill created from a conversation. This value is empty for user-uploaded skills.
	//
	// example:
	//
	// acp-bp4du4v74mc7qw8****
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
	// The skill status.
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The skill type.
	//
	// example:
	//
	// System
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The skill version.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeSkillsResponseBodySkillInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeSkillsResponseBodySkillInfo) GoString() string {
	return s.String()
}

func (s *DescribeSkillsResponseBodySkillInfo) GetCategory() *string {
	return s.Category
}

func (s *DescribeSkillsResponseBodySkillInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeSkillsResponseBodySkillInfo) GetDescription() *string {
	return s.Description
}

func (s *DescribeSkillsResponseBodySkillInfo) GetIconOssUrl() *string {
	return s.IconOssUrl
}

func (s *DescribeSkillsResponseBodySkillInfo) GetInstalledCount() *int32 {
	return s.InstalledCount
}

func (s *DescribeSkillsResponseBodySkillInfo) GetInstalledInstances() []*DescribeSkillsResponseBodySkillInfoInstalledInstances {
	return s.InstalledInstances
}

func (s *DescribeSkillsResponseBodySkillInfo) GetInstruction() *string {
	return s.Instruction
}

func (s *DescribeSkillsResponseBodySkillInfo) GetSkillId() *string {
	return s.SkillId
}

func (s *DescribeSkillsResponseBodySkillInfo) GetSkillName() *string {
	return s.SkillName
}

func (s *DescribeSkillsResponseBodySkillInfo) GetSkillSource() *string {
	return s.SkillSource
}

func (s *DescribeSkillsResponseBodySkillInfo) GetSkillStatus() *string {
	return s.SkillStatus
}

func (s *DescribeSkillsResponseBodySkillInfo) GetSourceInstanceId() *string {
	return s.SourceInstanceId
}

func (s *DescribeSkillsResponseBodySkillInfo) GetStatus() *string {
	return s.Status
}

func (s *DescribeSkillsResponseBodySkillInfo) GetType() *string {
	return s.Type
}

func (s *DescribeSkillsResponseBodySkillInfo) GetVersion() *string {
	return s.Version
}

func (s *DescribeSkillsResponseBodySkillInfo) SetCategory(v string) *DescribeSkillsResponseBodySkillInfo {
	s.Category = &v
	return s
}

func (s *DescribeSkillsResponseBodySkillInfo) SetCreateTime(v string) *DescribeSkillsResponseBodySkillInfo {
	s.CreateTime = &v
	return s
}

func (s *DescribeSkillsResponseBodySkillInfo) SetDescription(v string) *DescribeSkillsResponseBodySkillInfo {
	s.Description = &v
	return s
}

func (s *DescribeSkillsResponseBodySkillInfo) SetIconOssUrl(v string) *DescribeSkillsResponseBodySkillInfo {
	s.IconOssUrl = &v
	return s
}

func (s *DescribeSkillsResponseBodySkillInfo) SetInstalledCount(v int32) *DescribeSkillsResponseBodySkillInfo {
	s.InstalledCount = &v
	return s
}

func (s *DescribeSkillsResponseBodySkillInfo) SetInstalledInstances(v []*DescribeSkillsResponseBodySkillInfoInstalledInstances) *DescribeSkillsResponseBodySkillInfo {
	s.InstalledInstances = v
	return s
}

func (s *DescribeSkillsResponseBodySkillInfo) SetInstruction(v string) *DescribeSkillsResponseBodySkillInfo {
	s.Instruction = &v
	return s
}

func (s *DescribeSkillsResponseBodySkillInfo) SetSkillId(v string) *DescribeSkillsResponseBodySkillInfo {
	s.SkillId = &v
	return s
}

func (s *DescribeSkillsResponseBodySkillInfo) SetSkillName(v string) *DescribeSkillsResponseBodySkillInfo {
	s.SkillName = &v
	return s
}

func (s *DescribeSkillsResponseBodySkillInfo) SetSkillSource(v string) *DescribeSkillsResponseBodySkillInfo {
	s.SkillSource = &v
	return s
}

func (s *DescribeSkillsResponseBodySkillInfo) SetSkillStatus(v string) *DescribeSkillsResponseBodySkillInfo {
	s.SkillStatus = &v
	return s
}

func (s *DescribeSkillsResponseBodySkillInfo) SetSourceInstanceId(v string) *DescribeSkillsResponseBodySkillInfo {
	s.SourceInstanceId = &v
	return s
}

func (s *DescribeSkillsResponseBodySkillInfo) SetStatus(v string) *DescribeSkillsResponseBodySkillInfo {
	s.Status = &v
	return s
}

func (s *DescribeSkillsResponseBodySkillInfo) SetType(v string) *DescribeSkillsResponseBodySkillInfo {
	s.Type = &v
	return s
}

func (s *DescribeSkillsResponseBodySkillInfo) SetVersion(v string) *DescribeSkillsResponseBodySkillInfo {
	s.Version = &v
	return s
}

func (s *DescribeSkillsResponseBodySkillInfo) Validate() error {
	if s.InstalledInstances != nil {
		for _, item := range s.InstalledInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSkillsResponseBodySkillInfoInstalledInstances struct {
	// The installation status.
	//
	// example:
	//
	// INSTALLED
	InstallStatus *string `json:"InstallStatus,omitempty" xml:"InstallStatus,omitempty"`
	// The cloud phone instance ID.
	//
	// example:
	//
	// acp-6g3nocu5y9vaf****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeSkillsResponseBodySkillInfoInstalledInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeSkillsResponseBodySkillInfoInstalledInstances) GoString() string {
	return s.String()
}

func (s *DescribeSkillsResponseBodySkillInfoInstalledInstances) GetInstallStatus() *string {
	return s.InstallStatus
}

func (s *DescribeSkillsResponseBodySkillInfoInstalledInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSkillsResponseBodySkillInfoInstalledInstances) SetInstallStatus(v string) *DescribeSkillsResponseBodySkillInfoInstalledInstances {
	s.InstallStatus = &v
	return s
}

func (s *DescribeSkillsResponseBodySkillInfoInstalledInstances) SetInstanceId(v string) *DescribeSkillsResponseBodySkillInfoInstalledInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeSkillsResponseBodySkillInfoInstalledInstances) Validate() error {
	return dara.Validate(s)
}
