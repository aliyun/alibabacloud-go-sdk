// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifySentenceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *VerifySentenceRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *VerifySentenceRequest
	GetJsonStr() *string
}

type VerifySentenceRequest struct {
	// Workspace ID.
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// Full JSON string. For details, see the following table.
	//
	// This parameter is required.
	//
	// example:
	//
	// "{"taskId":"EA701F66-8CA2-4A79-8E3C-A6F2F****","fileName":"人工校验测试-订购茶叶.wav","dialogueId":1,"roleCorrect":false,"sourceRole":0,"textCorrect":false,"sourceText":"我要订购大量的信阳毛尖。","oldIncorrectWords":3,"targetText":"我要订购大大的南阳毛巾。","targetRole":1}"
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s VerifySentenceRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifySentenceRequest) GoString() string {
	return s.String()
}

func (s *VerifySentenceRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *VerifySentenceRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *VerifySentenceRequest) SetBaseMeAgentId(v int64) *VerifySentenceRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *VerifySentenceRequest) SetJsonStr(v string) *VerifySentenceRequest {
	s.JsonStr = &v
	return s
}

func (s *VerifySentenceRequest) Validate() error {
	return dara.Validate(s)
}
