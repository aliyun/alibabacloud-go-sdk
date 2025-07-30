// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSyncQualityCheckDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *UpdateSyncQualityCheckDataRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *UpdateSyncQualityCheckDataRequest
	GetJsonStr() *string
}

type UpdateSyncQualityCheckDataRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"tid":"xxx"}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UpdateSyncQualityCheckDataRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSyncQualityCheckDataRequest) GoString() string {
	return s.String()
}

func (s *UpdateSyncQualityCheckDataRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *UpdateSyncQualityCheckDataRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *UpdateSyncQualityCheckDataRequest) SetBaseMeAgentId(v int64) *UpdateSyncQualityCheckDataRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UpdateSyncQualityCheckDataRequest) SetJsonStr(v string) *UpdateSyncQualityCheckDataRequest {
	s.JsonStr = &v
	return s
}

func (s *UpdateSyncQualityCheckDataRequest) Validate() error {
	return dara.Validate(s)
}
