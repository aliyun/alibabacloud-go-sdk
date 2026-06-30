// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *VerifyFileRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *VerifyFileRequest
	GetJsonStr() *string
}

type VerifyFileRequest struct {
	// Workspace ID.
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// Complete JSON string information. For details, see the following information.
	//
	// This parameter is required.
	//
	// example:
	//
	// "{"taskId":"EA701F66-8CA2-4A79-8E3C-A6F2****","fileName":"人工校验测试-订购茶叶.wav"}"
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s VerifyFileRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyFileRequest) GoString() string {
	return s.String()
}

func (s *VerifyFileRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *VerifyFileRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *VerifyFileRequest) SetBaseMeAgentId(v int64) *VerifyFileRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *VerifyFileRequest) SetJsonStr(v string) *VerifyFileRequest {
	s.JsonStr = &v
	return s
}

func (s *VerifyFileRequest) Validate() error {
	return dara.Validate(s)
}
