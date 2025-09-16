// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIndexRecoverRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDesc(v string) *ListIndexRecoverRecordsResponseBody
	GetDesc() *string
	SetFinishedTime(v string) *ListIndexRecoverRecordsResponseBody
	GetFinishedTime() *string
	SetGenerationId(v string) *ListIndexRecoverRecordsResponseBody
	GetGenerationId() *string
}

type ListIndexRecoverRecordsResponseBody struct {
	// The description.
	//
	// example:
	//
	// test
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// The time when the index version was published.
	//
	// example:
	//
	// 2024-06-07 16:43:00
	FinishedTime *string `json:"finishedTime,omitempty" xml:"finishedTime,omitempty"`
	// The ID of the full index version.
	//
	// example:
	//
	// 1708674867
	GenerationId *string `json:"generationId,omitempty" xml:"generationId,omitempty"`
}

func (s ListIndexRecoverRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIndexRecoverRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIndexRecoverRecordsResponseBody) GetDesc() *string {
	return s.Desc
}

func (s *ListIndexRecoverRecordsResponseBody) GetFinishedTime() *string {
	return s.FinishedTime
}

func (s *ListIndexRecoverRecordsResponseBody) GetGenerationId() *string {
	return s.GenerationId
}

func (s *ListIndexRecoverRecordsResponseBody) SetDesc(v string) *ListIndexRecoverRecordsResponseBody {
	s.Desc = &v
	return s
}

func (s *ListIndexRecoverRecordsResponseBody) SetFinishedTime(v string) *ListIndexRecoverRecordsResponseBody {
	s.FinishedTime = &v
	return s
}

func (s *ListIndexRecoverRecordsResponseBody) SetGenerationId(v string) *ListIndexRecoverRecordsResponseBody {
	s.GenerationId = &v
	return s
}

func (s *ListIndexRecoverRecordsResponseBody) Validate() error {
	return dara.Validate(s)
}
