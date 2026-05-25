// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStageModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateStageModelResponseBody
	GetRequestId() *string
	SetStageModelId(v string) *CreateStageModelResponseBody
	GetStageModelId() *string
}

type CreateStageModelResponseBody struct {
	// example:
	//
	// EF924FE4-2EDD-4CD3-89EC-34E4708xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 39sdfjsd8fudsfxxx
	StageModelId *string `json:"StageModelId,omitempty" xml:"StageModelId,omitempty"`
}

func (s CreateStageModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateStageModelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStageModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateStageModelResponseBody) GetStageModelId() *string {
	return s.StageModelId
}

func (s *CreateStageModelResponseBody) SetRequestId(v string) *CreateStageModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStageModelResponseBody) SetStageModelId(v string) *CreateStageModelResponseBody {
	s.StageModelId = &v
	return s
}

func (s *CreateStageModelResponseBody) Validate() error {
	return dara.Validate(s)
}
