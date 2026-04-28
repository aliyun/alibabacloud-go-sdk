// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemorySkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLatency(v int32) *CreateMemorySkillResponseBody
	GetLatency() *int32
	SetRequestId(v string) *CreateMemorySkillResponseBody
	GetRequestId() *string
	SetResult(v *CreateMemorySkillResponseBodyResult) *CreateMemorySkillResponseBody
	GetResult() *CreateMemorySkillResponseBodyResult
	SetStatus(v string) *CreateMemorySkillResponseBody
	GetStatus() *string
}

type CreateMemorySkillResponseBody struct {
	Latency   *int32                               `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                              `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *CreateMemorySkillResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Status    *string                              `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateMemorySkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMemorySkillResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMemorySkillResponseBody) GetLatency() *int32 {
	return s.Latency
}

func (s *CreateMemorySkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMemorySkillResponseBody) GetResult() *CreateMemorySkillResponseBodyResult {
	return s.Result
}

func (s *CreateMemorySkillResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateMemorySkillResponseBody) SetLatency(v int32) *CreateMemorySkillResponseBody {
	s.Latency = &v
	return s
}

func (s *CreateMemorySkillResponseBody) SetRequestId(v string) *CreateMemorySkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMemorySkillResponseBody) SetResult(v *CreateMemorySkillResponseBodyResult) *CreateMemorySkillResponseBody {
	s.Result = v
	return s
}

func (s *CreateMemorySkillResponseBody) SetStatus(v string) *CreateMemorySkillResponseBody {
	s.Status = &v
	return s
}

func (s *CreateMemorySkillResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMemorySkillResponseBodyResult struct {
	Data          []*CreateMemorySkillResponseBodyResultData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	ImportedCount *int32                                     `json:"imported_count,omitempty" xml:"imported_count,omitempty"`
}

func (s CreateMemorySkillResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateMemorySkillResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateMemorySkillResponseBodyResult) GetData() []*CreateMemorySkillResponseBodyResultData {
	return s.Data
}

func (s *CreateMemorySkillResponseBodyResult) GetImportedCount() *int32 {
	return s.ImportedCount
}

func (s *CreateMemorySkillResponseBodyResult) SetData(v []*CreateMemorySkillResponseBodyResultData) *CreateMemorySkillResponseBodyResult {
	s.Data = v
	return s
}

func (s *CreateMemorySkillResponseBodyResult) SetImportedCount(v int32) *CreateMemorySkillResponseBodyResult {
	s.ImportedCount = &v
	return s
}

func (s *CreateMemorySkillResponseBodyResult) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateMemorySkillResponseBodyResultData struct {
	Description   *string   `json:"description,omitempty" xml:"description,omitempty"`
	Id            *string   `json:"id,omitempty" xml:"id,omitempty"`
	Name          *string   `json:"name,omitempty" xml:"name,omitempty"`
	Owner         *string   `json:"owner,omitempty" xml:"owner,omitempty"`
	ResourcePaths []*string `json:"resource_paths,omitempty" xml:"resource_paths,omitempty" type:"Repeated"`
	Tags          []*string `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	Triggers      []*string `json:"triggers,omitempty" xml:"triggers,omitempty" type:"Repeated"`
	UpdatedAt     *string   `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	Version       *string   `json:"version,omitempty" xml:"version,omitempty"`
}

func (s CreateMemorySkillResponseBodyResultData) String() string {
	return dara.Prettify(s)
}

func (s CreateMemorySkillResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *CreateMemorySkillResponseBodyResultData) GetDescription() *string {
	return s.Description
}

func (s *CreateMemorySkillResponseBodyResultData) GetId() *string {
	return s.Id
}

func (s *CreateMemorySkillResponseBodyResultData) GetName() *string {
	return s.Name
}

func (s *CreateMemorySkillResponseBodyResultData) GetOwner() *string {
	return s.Owner
}

func (s *CreateMemorySkillResponseBodyResultData) GetResourcePaths() []*string {
	return s.ResourcePaths
}

func (s *CreateMemorySkillResponseBodyResultData) GetTags() []*string {
	return s.Tags
}

func (s *CreateMemorySkillResponseBodyResultData) GetTriggers() []*string {
	return s.Triggers
}

func (s *CreateMemorySkillResponseBodyResultData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *CreateMemorySkillResponseBodyResultData) GetVersion() *string {
	return s.Version
}

func (s *CreateMemorySkillResponseBodyResultData) SetDescription(v string) *CreateMemorySkillResponseBodyResultData {
	s.Description = &v
	return s
}

func (s *CreateMemorySkillResponseBodyResultData) SetId(v string) *CreateMemorySkillResponseBodyResultData {
	s.Id = &v
	return s
}

func (s *CreateMemorySkillResponseBodyResultData) SetName(v string) *CreateMemorySkillResponseBodyResultData {
	s.Name = &v
	return s
}

func (s *CreateMemorySkillResponseBodyResultData) SetOwner(v string) *CreateMemorySkillResponseBodyResultData {
	s.Owner = &v
	return s
}

func (s *CreateMemorySkillResponseBodyResultData) SetResourcePaths(v []*string) *CreateMemorySkillResponseBodyResultData {
	s.ResourcePaths = v
	return s
}

func (s *CreateMemorySkillResponseBodyResultData) SetTags(v []*string) *CreateMemorySkillResponseBodyResultData {
	s.Tags = v
	return s
}

func (s *CreateMemorySkillResponseBodyResultData) SetTriggers(v []*string) *CreateMemorySkillResponseBodyResultData {
	s.Triggers = v
	return s
}

func (s *CreateMemorySkillResponseBodyResultData) SetUpdatedAt(v string) *CreateMemorySkillResponseBodyResultData {
	s.UpdatedAt = &v
	return s
}

func (s *CreateMemorySkillResponseBodyResultData) SetVersion(v string) *CreateMemorySkillResponseBodyResultData {
	s.Version = &v
	return s
}

func (s *CreateMemorySkillResponseBodyResultData) Validate() error {
	return dara.Validate(s)
}
