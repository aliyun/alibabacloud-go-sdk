// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitQualityCheckTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *SubmitQualityCheckTaskRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *SubmitQualityCheckTaskRequest
	GetJsonStr() *string
}

type SubmitQualityCheckTaskRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s SubmitQualityCheckTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitQualityCheckTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitQualityCheckTaskRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *SubmitQualityCheckTaskRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *SubmitQualityCheckTaskRequest) SetBaseMeAgentId(v int64) *SubmitQualityCheckTaskRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *SubmitQualityCheckTaskRequest) SetJsonStr(v string) *SubmitQualityCheckTaskRequest {
	s.JsonStr = &v
	return s
}

func (s *SubmitQualityCheckTaskRequest) Validate() error {
	return dara.Validate(s)
}
