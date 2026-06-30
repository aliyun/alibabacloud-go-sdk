// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *UploadDataRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *UploadDataRequest
	GetJsonStr() *string
}

type UploadDataRequest struct {
	// Workspace ID. Use this to select a specific workspace in multi-workspace scenarios. Defaults to the default workspace.
	//
	// example:
	//
	// 123456
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// JsonStr is a JSON string that contains all custom parameters for this API. See the JsonStr property description below.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UploadDataRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadDataRequest) GoString() string {
	return s.String()
}

func (s *UploadDataRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *UploadDataRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *UploadDataRequest) SetBaseMeAgentId(v int64) *UploadDataRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UploadDataRequest) SetJsonStr(v string) *UploadDataRequest {
	s.JsonStr = &v
	return s
}

func (s *UploadDataRequest) Validate() error {
	return dara.Validate(s)
}
