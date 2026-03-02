// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPbcReviewListResult interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*PbcReview) *PbcReviewListResult
	GetData() []*PbcReview
	SetTotal(v int32) *PbcReviewListResult
	GetTotal() *int32
}

type PbcReviewListResult struct {
	Data []*PbcReview `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// 总数量
	//
	// example:
	//
	// 24
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s PbcReviewListResult) String() string {
	return dara.Prettify(s)
}

func (s PbcReviewListResult) GoString() string {
	return s.String()
}

func (s *PbcReviewListResult) GetData() []*PbcReview {
	return s.Data
}

func (s *PbcReviewListResult) GetTotal() *int32 {
	return s.Total
}

func (s *PbcReviewListResult) SetData(v []*PbcReview) *PbcReviewListResult {
	s.Data = v
	return s
}

func (s *PbcReviewListResult) SetTotal(v int32) *PbcReviewListResult {
	s.Total = &v
	return s
}

func (s *PbcReviewListResult) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
