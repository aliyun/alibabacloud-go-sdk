// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSyncResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *GetSyncResultRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *GetSyncResultRequest
	GetJsonStr() *string
}

type GetSyncResultRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"pageNumber":1,"pageSize":10,"requiredFields":"asrResult,agent,status,errorMessage,reviewStatus,reviewResult,score,taskId,reviewer,resolver,recording.name,recording.duration,recording.url,hitResult,business","startTime":"2020-12-25 00:00:00","endTime":"2020-12-31 23:59:59"}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetSyncResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSyncResultRequest) GoString() string {
	return s.String()
}

func (s *GetSyncResultRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *GetSyncResultRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *GetSyncResultRequest) SetBaseMeAgentId(v int64) *GetSyncResultRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetSyncResultRequest) SetJsonStr(v string) *GetSyncResultRequest {
	s.JsonStr = &v
	return s
}

func (s *GetSyncResultRequest) Validate() error {
	return dara.Validate(s)
}
