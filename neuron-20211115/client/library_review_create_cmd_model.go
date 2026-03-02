// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLibraryReviewCreateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetLibraryId(v int64) *LibraryReviewCreateCmd
	GetLibraryId() *int64
	SetLibraryUrl(v string) *LibraryReviewCreateCmd
	GetLibraryUrl() *string
	SetMarketId(v int64) *LibraryReviewCreateCmd
	GetMarketId() *int64
	SetReviewerId(v string) *LibraryReviewCreateCmd
	GetReviewerId() *string
}

type LibraryReviewCreateCmd struct {
	LibraryId  *int64  `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	LibraryUrl *string `json:"libraryUrl,omitempty" xml:"libraryUrl,omitempty"`
	MarketId   *int64  `json:"marketId,omitempty" xml:"marketId,omitempty"`
	ReviewerId *string `json:"reviewerId,omitempty" xml:"reviewerId,omitempty"`
}

func (s LibraryReviewCreateCmd) String() string {
	return dara.Prettify(s)
}

func (s LibraryReviewCreateCmd) GoString() string {
	return s.String()
}

func (s *LibraryReviewCreateCmd) GetLibraryId() *int64 {
	return s.LibraryId
}

func (s *LibraryReviewCreateCmd) GetLibraryUrl() *string {
	return s.LibraryUrl
}

func (s *LibraryReviewCreateCmd) GetMarketId() *int64 {
	return s.MarketId
}

func (s *LibraryReviewCreateCmd) GetReviewerId() *string {
	return s.ReviewerId
}

func (s *LibraryReviewCreateCmd) SetLibraryId(v int64) *LibraryReviewCreateCmd {
	s.LibraryId = &v
	return s
}

func (s *LibraryReviewCreateCmd) SetLibraryUrl(v string) *LibraryReviewCreateCmd {
	s.LibraryUrl = &v
	return s
}

func (s *LibraryReviewCreateCmd) SetMarketId(v int64) *LibraryReviewCreateCmd {
	s.MarketId = &v
	return s
}

func (s *LibraryReviewCreateCmd) SetReviewerId(v string) *LibraryReviewCreateCmd {
	s.ReviewerId = &v
	return s
}

func (s *LibraryReviewCreateCmd) Validate() error {
	return dara.Validate(s)
}
