// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQualityCheckDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *UpdateQualityCheckDataRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *UpdateQualityCheckDataRequest
	GetJsonStr() *string
}

type UpdateQualityCheckDataRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"taskId":"xxx"}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UpdateQualityCheckDataRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateQualityCheckDataRequest) GoString() string {
	return s.String()
}

func (s *UpdateQualityCheckDataRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *UpdateQualityCheckDataRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *UpdateQualityCheckDataRequest) SetBaseMeAgentId(v int64) *UpdateQualityCheckDataRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UpdateQualityCheckDataRequest) SetJsonStr(v string) *UpdateQualityCheckDataRequest {
	s.JsonStr = &v
	return s
}

func (s *UpdateQualityCheckDataRequest) Validate() error {
	return dara.Validate(s)
}
