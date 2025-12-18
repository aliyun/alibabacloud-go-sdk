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
	// The ID of the account group.
	//
	// For more information about how to obtain the ID of an account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-4b1b626622af000c****
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
