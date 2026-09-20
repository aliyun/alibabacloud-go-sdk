// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetNamespaceResponseBodyData) *GetNamespaceResponseBody
	GetData() *GetNamespaceResponseBodyData
	SetRequestId(v string) *GetNamespaceResponseBody
	GetRequestId() *string
}

type GetNamespaceResponseBody struct {
	// The namespace information.
	Data *GetNamespaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D9E87E66-9EF0-5C10-A5E6-924020A0C9B7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNamespaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetNamespaceResponseBody) GetData() *GetNamespaceResponseBodyData {
	return s.Data
}

func (s *GetNamespaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNamespaceResponseBody) SetData(v *GetNamespaceResponseBodyData) *GetNamespaceResponseBody {
	s.Data = v
	return s
}

func (s *GetNamespaceResponseBody) SetRequestId(v string) *GetNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNamespaceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetNamespaceResponseBodyData struct {
	// The time when the namespace was created.
	//
	// example:
	//
	// 2025-11-17T09:57:38+08:00
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The description of the namespace.
	//
	// example:
	//
	// secret for bbtadmin
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	IpWhitelist *string `json:"IpWhitelist,omitempty" xml:"IpWhitelist,omitempty"`
	// The namespace name.
	//
	// example:
	//
	// magic:magic-cn-1us4sed5d01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace ID.
	//
	// example:
	//
	// a2a9310a-9d91-4283-b4e2-844f6d45fe64
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The number of prompts in the namespace.
	//
	// example:
	//
	// 1
	PromptCount         *int32  `json:"PromptCount,omitempty" xml:"PromptCount,omitempty"`
	PublicAccessEnabled *bool   `json:"PublicAccessEnabled,omitempty" xml:"PublicAccessEnabled,omitempty"`
	PublicDomain        *string `json:"PublicDomain,omitempty" xml:"PublicDomain,omitempty"`
	// The scan policy.
	//
	// The policy contains two configuration items:
	//
	// - minBlockRiskLevel: the risk level for blocking.
	//
	//   - high: blocks high-risk items.
	//
	//   - medium: blocks medium- and high-risk items.
	//
	//   - low: blocks all risk levels including high, medium, and low.
	//
	// - maxSkipRatio: the maximum skip ratio. If the scan skip ratio exceeds this value, the scan is considered as failed.
	//
	// example:
	//
	// {"minBlockRiskLevel":"medium","maxSkipRatio":0.2}
	ScanPolicy *string `json:"ScanPolicy,omitempty" xml:"ScanPolicy,omitempty"`
	// The number of skills in the namespace.
	//
	// example:
	//
	// 1
	SkillCount *int32 `json:"SkillCount,omitempty" xml:"SkillCount,omitempty"`
	// The source of the namespace.
	//
	// example:
	//
	// magic:magic-cn-1us4sed5d01
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The source ordinal number of the namespace.
	//
	// example:
	//
	// 0
	SourceIndex *int32 `json:"SourceIndex,omitempty" xml:"SourceIndex,omitempty"`
	// The tags of the namespace.
	//
	// example:
	//
	// {}
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s GetNamespaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetNamespaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetNamespaceResponseBodyData) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *GetNamespaceResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetNamespaceResponseBodyData) GetIpWhitelist() *string {
	return s.IpWhitelist
}

func (s *GetNamespaceResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetNamespaceResponseBodyData) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *GetNamespaceResponseBodyData) GetPromptCount() *int32 {
	return s.PromptCount
}

func (s *GetNamespaceResponseBodyData) GetPublicAccessEnabled() *bool {
	return s.PublicAccessEnabled
}

func (s *GetNamespaceResponseBodyData) GetPublicDomain() *string {
	return s.PublicDomain
}

func (s *GetNamespaceResponseBodyData) GetScanPolicy() *string {
	return s.ScanPolicy
}

func (s *GetNamespaceResponseBodyData) GetSkillCount() *int32 {
	return s.SkillCount
}

func (s *GetNamespaceResponseBodyData) GetSource() *string {
	return s.Source
}

func (s *GetNamespaceResponseBodyData) GetSourceIndex() *int32 {
	return s.SourceIndex
}

func (s *GetNamespaceResponseBodyData) GetTags() *string {
	return s.Tags
}

func (s *GetNamespaceResponseBodyData) SetCreatedTime(v string) *GetNamespaceResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *GetNamespaceResponseBodyData) SetDescription(v string) *GetNamespaceResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetNamespaceResponseBodyData) SetIpWhitelist(v string) *GetNamespaceResponseBodyData {
	s.IpWhitelist = &v
	return s
}

func (s *GetNamespaceResponseBodyData) SetName(v string) *GetNamespaceResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetNamespaceResponseBodyData) SetNamespaceId(v string) *GetNamespaceResponseBodyData {
	s.NamespaceId = &v
	return s
}

func (s *GetNamespaceResponseBodyData) SetPromptCount(v int32) *GetNamespaceResponseBodyData {
	s.PromptCount = &v
	return s
}

func (s *GetNamespaceResponseBodyData) SetPublicAccessEnabled(v bool) *GetNamespaceResponseBodyData {
	s.PublicAccessEnabled = &v
	return s
}

func (s *GetNamespaceResponseBodyData) SetPublicDomain(v string) *GetNamespaceResponseBodyData {
	s.PublicDomain = &v
	return s
}

func (s *GetNamespaceResponseBodyData) SetScanPolicy(v string) *GetNamespaceResponseBodyData {
	s.ScanPolicy = &v
	return s
}

func (s *GetNamespaceResponseBodyData) SetSkillCount(v int32) *GetNamespaceResponseBodyData {
	s.SkillCount = &v
	return s
}

func (s *GetNamespaceResponseBodyData) SetSource(v string) *GetNamespaceResponseBodyData {
	s.Source = &v
	return s
}

func (s *GetNamespaceResponseBodyData) SetSourceIndex(v int32) *GetNamespaceResponseBodyData {
	s.SourceIndex = &v
	return s
}

func (s *GetNamespaceResponseBodyData) SetTags(v string) *GetNamespaceResponseBodyData {
	s.Tags = &v
	return s
}

func (s *GetNamespaceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
