// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAsrVocabRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *GetAsrVocabRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *GetAsrVocabRequest
	GetJsonStr() *string
}

type GetAsrVocabRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64  `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	JsonStr       *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetAsrVocabRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAsrVocabRequest) GoString() string {
	return s.String()
}

func (s *GetAsrVocabRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *GetAsrVocabRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *GetAsrVocabRequest) SetBaseMeAgentId(v int64) *GetAsrVocabRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetAsrVocabRequest) SetJsonStr(v string) *GetAsrVocabRequest {
	s.JsonStr = &v
	return s
}

func (s *GetAsrVocabRequest) Validate() error {
	return dara.Validate(s)
}
