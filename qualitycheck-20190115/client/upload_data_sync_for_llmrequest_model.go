// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadDataSyncForLLMRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *UploadDataSyncForLLMRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *UploadDataSyncForLLMRequest
	GetJsonStr() *string
}

type UploadDataSyncForLLMRequest struct {
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UploadDataSyncForLLMRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadDataSyncForLLMRequest) GoString() string {
	return s.String()
}

func (s *UploadDataSyncForLLMRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *UploadDataSyncForLLMRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *UploadDataSyncForLLMRequest) SetBaseMeAgentId(v int64) *UploadDataSyncForLLMRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UploadDataSyncForLLMRequest) SetJsonStr(v string) *UploadDataSyncForLLMRequest {
	s.JsonStr = &v
	return s
}

func (s *UploadDataSyncForLLMRequest) Validate() error {
	return dara.Validate(s)
}
