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
	Data      *GetNamespaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	PromptCount *int32  `json:"PromptCount,omitempty" xml:"PromptCount,omitempty"`
	ScanPolicy  *string `json:"ScanPolicy,omitempty" xml:"ScanPolicy,omitempty"`
	SkillCount  *int32  `json:"SkillCount,omitempty" xml:"SkillCount,omitempty"`
	Source      *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceIndex *int32  `json:"SourceIndex,omitempty" xml:"SourceIndex,omitempty"`
	Tags        *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
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

func (s *GetNamespaceResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetNamespaceResponseBodyData) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *GetNamespaceResponseBodyData) GetPromptCount() *int32 {
	return s.PromptCount
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
