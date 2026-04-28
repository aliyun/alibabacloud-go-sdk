// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemorySkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLatency(v int32) *UpdateMemorySkillResponseBody
	GetLatency() *int32
	SetRequestId(v string) *UpdateMemorySkillResponseBody
	GetRequestId() *string
	SetResult(v *UpdateMemorySkillResponseBodyResult) *UpdateMemorySkillResponseBody
	GetResult() *UpdateMemorySkillResponseBodyResult
	SetStatus(v string) *UpdateMemorySkillResponseBody
	GetStatus() *string
}

type UpdateMemorySkillResponseBody struct {
	Latency   *int32                               `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                              `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *UpdateMemorySkillResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Status    *string                              `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateMemorySkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemorySkillResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMemorySkillResponseBody) GetLatency() *int32 {
	return s.Latency
}

func (s *UpdateMemorySkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMemorySkillResponseBody) GetResult() *UpdateMemorySkillResponseBodyResult {
	return s.Result
}

func (s *UpdateMemorySkillResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UpdateMemorySkillResponseBody) SetLatency(v int32) *UpdateMemorySkillResponseBody {
	s.Latency = &v
	return s
}

func (s *UpdateMemorySkillResponseBody) SetRequestId(v string) *UpdateMemorySkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMemorySkillResponseBody) SetResult(v *UpdateMemorySkillResponseBodyResult) *UpdateMemorySkillResponseBody {
	s.Result = v
	return s
}

func (s *UpdateMemorySkillResponseBody) SetStatus(v string) *UpdateMemorySkillResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateMemorySkillResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMemorySkillResponseBodyResult struct {
	Files   map[string]*string `json:"files,omitempty" xml:"files,omitempty"`
	Name    *string            `json:"name,omitempty" xml:"name,omitempty"`
	Version *string            `json:"version,omitempty" xml:"version,omitempty"`
}

func (s UpdateMemorySkillResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemorySkillResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateMemorySkillResponseBodyResult) GetFiles() map[string]*string {
	return s.Files
}

func (s *UpdateMemorySkillResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *UpdateMemorySkillResponseBodyResult) GetVersion() *string {
	return s.Version
}

func (s *UpdateMemorySkillResponseBodyResult) SetFiles(v map[string]*string) *UpdateMemorySkillResponseBodyResult {
	s.Files = v
	return s
}

func (s *UpdateMemorySkillResponseBodyResult) SetName(v string) *UpdateMemorySkillResponseBodyResult {
	s.Name = &v
	return s
}

func (s *UpdateMemorySkillResponseBodyResult) SetVersion(v string) *UpdateMemorySkillResponseBodyResult {
	s.Version = &v
	return s
}

func (s *UpdateMemorySkillResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
