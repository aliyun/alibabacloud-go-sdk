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
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
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
