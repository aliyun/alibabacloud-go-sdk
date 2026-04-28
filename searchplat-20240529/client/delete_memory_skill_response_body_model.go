// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMemorySkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLatency(v int32) *DeleteMemorySkillResponseBody
	GetLatency() *int32
	SetRequestId(v string) *DeleteMemorySkillResponseBody
	GetRequestId() *string
	SetResult(v *DeleteMemorySkillResponseBodyResult) *DeleteMemorySkillResponseBody
	GetResult() *DeleteMemorySkillResponseBodyResult
	SetStatus(v string) *DeleteMemorySkillResponseBody
	GetStatus() *string
}

type DeleteMemorySkillResponseBody struct {
	Latency   *int32                               `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                              `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *DeleteMemorySkillResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Status    *string                              `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DeleteMemorySkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMemorySkillResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMemorySkillResponseBody) GetLatency() *int32 {
	return s.Latency
}

func (s *DeleteMemorySkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMemorySkillResponseBody) GetResult() *DeleteMemorySkillResponseBodyResult {
	return s.Result
}

func (s *DeleteMemorySkillResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DeleteMemorySkillResponseBody) SetLatency(v int32) *DeleteMemorySkillResponseBody {
	s.Latency = &v
	return s
}

func (s *DeleteMemorySkillResponseBody) SetRequestId(v string) *DeleteMemorySkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMemorySkillResponseBody) SetResult(v *DeleteMemorySkillResponseBodyResult) *DeleteMemorySkillResponseBody {
	s.Result = v
	return s
}

func (s *DeleteMemorySkillResponseBody) SetStatus(v string) *DeleteMemorySkillResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteMemorySkillResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteMemorySkillResponseBodyResult struct {
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	SkillId *string `json:"skill_id,omitempty" xml:"skill_id,omitempty"`
	Status  *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DeleteMemorySkillResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteMemorySkillResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteMemorySkillResponseBodyResult) GetMessage() *string {
	return s.Message
}

func (s *DeleteMemorySkillResponseBodyResult) GetSkillId() *string {
	return s.SkillId
}

func (s *DeleteMemorySkillResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *DeleteMemorySkillResponseBodyResult) SetMessage(v string) *DeleteMemorySkillResponseBodyResult {
	s.Message = &v
	return s
}

func (s *DeleteMemorySkillResponseBodyResult) SetSkillId(v string) *DeleteMemorySkillResponseBodyResult {
	s.SkillId = &v
	return s
}

func (s *DeleteMemorySkillResponseBodyResult) SetStatus(v string) *DeleteMemorySkillResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DeleteMemorySkillResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
