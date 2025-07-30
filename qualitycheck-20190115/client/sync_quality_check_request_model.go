// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncQualityCheckRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *SyncQualityCheckRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *SyncQualityCheckRequest
	GetJsonStr() *string
}

type SyncQualityCheckRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"tid":"20200823-234234","dialogue":"{}"}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s SyncQualityCheckRequest) String() string {
	return dara.Prettify(s)
}

func (s SyncQualityCheckRequest) GoString() string {
	return s.String()
}

func (s *SyncQualityCheckRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *SyncQualityCheckRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *SyncQualityCheckRequest) SetBaseMeAgentId(v int64) *SyncQualityCheckRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *SyncQualityCheckRequest) SetJsonStr(v string) *SyncQualityCheckRequest {
	s.JsonStr = &v
	return s
}

func (s *SyncQualityCheckRequest) Validate() error {
	return dara.Validate(s)
}
