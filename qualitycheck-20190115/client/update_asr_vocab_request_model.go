// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAsrVocabRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *UpdateAsrVocabRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *UpdateAsrVocabRequest
	GetJsonStr() *string
}

type UpdateAsrVocabRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UpdateAsrVocabRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAsrVocabRequest) GoString() string {
	return s.String()
}

func (s *UpdateAsrVocabRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *UpdateAsrVocabRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *UpdateAsrVocabRequest) SetBaseMeAgentId(v int64) *UpdateAsrVocabRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UpdateAsrVocabRequest) SetJsonStr(v string) *UpdateAsrVocabRequest {
	s.JsonStr = &v
	return s
}

func (s *UpdateAsrVocabRequest) Validate() error {
	return dara.Validate(s)
}
