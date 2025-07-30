// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitReviewInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *SubmitReviewInfoRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *SubmitReviewInfoRequest
	GetJsonStr() *string
}

type SubmitReviewInfoRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s SubmitReviewInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitReviewInfoRequest) GoString() string {
	return s.String()
}

func (s *SubmitReviewInfoRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *SubmitReviewInfoRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *SubmitReviewInfoRequest) SetBaseMeAgentId(v int64) *SubmitReviewInfoRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *SubmitReviewInfoRequest) SetJsonStr(v string) *SubmitReviewInfoRequest {
	s.JsonStr = &v
	return s
}

func (s *SubmitReviewInfoRequest) Validate() error {
	return dara.Validate(s)
}
