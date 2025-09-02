// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodesByBaselineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaselineId(v int64) *ListNodesByBaselineRequest
	GetBaselineId() *int64
}

type ListNodesByBaselineRequest struct {
	// The baseline ID. You can call the [ListBaselineConfigs](https://help.aliyun.com/document_detail/173964.html) operation to query the baseline ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	BaselineId *int64 `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
}

func (s ListNodesByBaselineRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNodesByBaselineRequest) GoString() string {
	return s.String()
}

func (s *ListNodesByBaselineRequest) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *ListNodesByBaselineRequest) SetBaselineId(v int64) *ListNodesByBaselineRequest {
	s.BaselineId = &v
	return s
}

func (s *ListNodesByBaselineRequest) Validate() error {
	return dara.Validate(s)
}
