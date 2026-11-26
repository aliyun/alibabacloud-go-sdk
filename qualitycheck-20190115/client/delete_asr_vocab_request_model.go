// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAsrVocabRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *DeleteAsrVocabRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *DeleteAsrVocabRequest
	GetJsonStr() *string
}

type DeleteAsrVocabRequest struct {
	// The workspace ID.
	//
	// example:
	//
	// 123456
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// The complete JSON string. For more information, see the following details.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"vocabId":"Specified hot word ID"}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s DeleteAsrVocabRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAsrVocabRequest) GoString() string {
	return s.String()
}

func (s *DeleteAsrVocabRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *DeleteAsrVocabRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *DeleteAsrVocabRequest) SetBaseMeAgentId(v int64) *DeleteAsrVocabRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *DeleteAsrVocabRequest) SetJsonStr(v string) *DeleteAsrVocabRequest {
	s.JsonStr = &v
	return s
}

func (s *DeleteAsrVocabRequest) Validate() error {
	return dara.Validate(s)
}
