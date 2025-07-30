// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResultToReviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *GetResultToReviewRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *GetResultToReviewRequest
	GetJsonStr() *string
}

type GetResultToReviewRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetResultToReviewRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResultToReviewRequest) GoString() string {
	return s.String()
}

func (s *GetResultToReviewRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *GetResultToReviewRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *GetResultToReviewRequest) SetBaseMeAgentId(v int64) *GetResultToReviewRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetResultToReviewRequest) SetJsonStr(v string) *GetResultToReviewRequest {
	s.JsonStr = &v
	return s
}

func (s *GetResultToReviewRequest) Validate() error {
	return dara.Validate(s)
}
