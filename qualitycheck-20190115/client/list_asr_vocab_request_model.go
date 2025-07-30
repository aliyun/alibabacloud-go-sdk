// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAsrVocabRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *ListAsrVocabRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *ListAsrVocabRequest
	GetJsonStr() *string
}

type ListAsrVocabRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"pageSize":1}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s ListAsrVocabRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAsrVocabRequest) GoString() string {
	return s.String()
}

func (s *ListAsrVocabRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *ListAsrVocabRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *ListAsrVocabRequest) SetBaseMeAgentId(v int64) *ListAsrVocabRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *ListAsrVocabRequest) SetJsonStr(v string) *ListAsrVocabRequest {
	s.JsonStr = &v
	return s
}

func (s *ListAsrVocabRequest) Validate() error {
	return dara.Validate(s)
}
