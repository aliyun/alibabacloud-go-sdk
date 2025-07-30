// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadDataSyncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *UploadDataSyncRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *UploadDataSyncRequest
	GetJsonStr() *string
}

type UploadDataSyncRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"tickets":xxx}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UploadDataSyncRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadDataSyncRequest) GoString() string {
	return s.String()
}

func (s *UploadDataSyncRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *UploadDataSyncRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *UploadDataSyncRequest) SetBaseMeAgentId(v int64) *UploadDataSyncRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UploadDataSyncRequest) SetJsonStr(v string) *UploadDataSyncRequest {
	s.JsonStr = &v
	return s
}

func (s *UploadDataSyncRequest) Validate() error {
	return dara.Validate(s)
}
