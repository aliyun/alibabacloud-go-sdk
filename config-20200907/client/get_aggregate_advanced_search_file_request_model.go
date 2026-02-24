// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateAdvancedSearchFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *GetAggregateAdvancedSearchFileRequest
	GetAggregatorId() *string
}

type GetAggregateAdvancedSearchFileRequest struct {
	// This parameter is required.
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
}

func (s GetAggregateAdvancedSearchFileRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateAdvancedSearchFileRequest) GoString() string {
	return s.String()
}

func (s *GetAggregateAdvancedSearchFileRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *GetAggregateAdvancedSearchFileRequest) SetAggregatorId(v string) *GetAggregateAdvancedSearchFileRequest {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateAdvancedSearchFileRequest) Validate() error {
	return dara.Validate(s)
}
