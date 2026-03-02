// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPbcReviewCreateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetMarketId(v int64) *PbcReviewCreateCmd
	GetMarketId() *int64
	SetPbcUrl(v string) *PbcReviewCreateCmd
	GetPbcUrl() *string
	SetPbcVersionId(v int64) *PbcReviewCreateCmd
	GetPbcVersionId() *int64
	SetReviewerId(v int64) *PbcReviewCreateCmd
	GetReviewerId() *int64
}

type PbcReviewCreateCmd struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	MarketId *int64 `json:"marketId,omitempty" xml:"marketId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://catalog.e2.aliyun.com/pbc/product
	PbcUrl *string `json:"pbcUrl,omitempty" xml:"pbcUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PbcVersionId *int64 `json:"pbcVersionId,omitempty" xml:"pbcVersionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ReviewerId *int64 `json:"reviewerId,omitempty" xml:"reviewerId,omitempty"`
}

func (s PbcReviewCreateCmd) String() string {
	return dara.Prettify(s)
}

func (s PbcReviewCreateCmd) GoString() string {
	return s.String()
}

func (s *PbcReviewCreateCmd) GetMarketId() *int64 {
	return s.MarketId
}

func (s *PbcReviewCreateCmd) GetPbcUrl() *string {
	return s.PbcUrl
}

func (s *PbcReviewCreateCmd) GetPbcVersionId() *int64 {
	return s.PbcVersionId
}

func (s *PbcReviewCreateCmd) GetReviewerId() *int64 {
	return s.ReviewerId
}

func (s *PbcReviewCreateCmd) SetMarketId(v int64) *PbcReviewCreateCmd {
	s.MarketId = &v
	return s
}

func (s *PbcReviewCreateCmd) SetPbcUrl(v string) *PbcReviewCreateCmd {
	s.PbcUrl = &v
	return s
}

func (s *PbcReviewCreateCmd) SetPbcVersionId(v int64) *PbcReviewCreateCmd {
	s.PbcVersionId = &v
	return s
}

func (s *PbcReviewCreateCmd) SetReviewerId(v int64) *PbcReviewCreateCmd {
	s.ReviewerId = &v
	return s
}

func (s *PbcReviewCreateCmd) Validate() error {
	return dara.Validate(s)
}
