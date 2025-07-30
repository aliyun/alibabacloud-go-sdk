// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *GetResultRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *GetResultRequest
	GetJsonStr() *string
}

type GetResultRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"pageNumber":1,"pageSize":10,"excludeFields":"hitResult.hits, recording.url","requiredFields":"agent,status,errorMessage,reviewStatus,reviewResult,score,taskId,reviewer,resolver,recording.name,recording.duration,hitResult,business","dataType":1,"sourceType":0,"startTime":"2020-06-25 00:00:00","endTime":"2020-07-01 23:59:59"}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResultRequest) GoString() string {
	return s.String()
}

func (s *GetResultRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *GetResultRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *GetResultRequest) SetBaseMeAgentId(v int64) *GetResultRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetResultRequest) SetJsonStr(v string) *GetResultRequest {
	s.JsonStr = &v
	return s
}

func (s *GetResultRequest) Validate() error {
	return dara.Validate(s)
}
