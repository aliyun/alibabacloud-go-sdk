// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadAudioDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *UploadAudioDataRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *UploadAudioDataRequest
	GetJsonStr() *string
}

type UploadAudioDataRequest struct {
	// The business space ID. In multi-business space scenarios, this parameter specifies the business space to use. Default value: the default business space.
	//
	// example:
	//
	// 123456
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// The complete JSON string. For details about the content, see the following detailed information.
	//
	// This parameter is required.
	//
	// example:
	//
	// {“callList”:“xxxxx”}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UploadAudioDataRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadAudioDataRequest) GoString() string {
	return s.String()
}

func (s *UploadAudioDataRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *UploadAudioDataRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *UploadAudioDataRequest) SetBaseMeAgentId(v int64) *UploadAudioDataRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UploadAudioDataRequest) SetJsonStr(v string) *UploadAudioDataRequest {
	s.JsonStr = &v
	return s
}

func (s *UploadAudioDataRequest) Validate() error {
	return dara.Validate(s)
}
