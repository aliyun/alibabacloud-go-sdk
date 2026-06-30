// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAsrVocabRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *CreateAsrVocabRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *CreateAsrVocabRequest
	GetJsonStr() *string
}

type CreateAsrVocabRequest struct {
	// Workspace ID
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// Full JSON string. See the details below.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"name":"热词名称"}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s CreateAsrVocabRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAsrVocabRequest) GoString() string {
	return s.String()
}

func (s *CreateAsrVocabRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *CreateAsrVocabRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *CreateAsrVocabRequest) SetBaseMeAgentId(v int64) *CreateAsrVocabRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *CreateAsrVocabRequest) SetJsonStr(v string) *CreateAsrVocabRequest {
	s.JsonStr = &v
	return s
}

func (s *CreateAsrVocabRequest) Validate() error {
	return dara.Validate(s)
}
