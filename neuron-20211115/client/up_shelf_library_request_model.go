// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpShelfLibraryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMarketId(v int64) *UpShelfLibraryRequest
	GetMarketId() *int64
}

type UpShelfLibraryRequest struct {
	MarketId *int64 `json:"marketId,omitempty" xml:"marketId,omitempty"`
}

func (s UpShelfLibraryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpShelfLibraryRequest) GoString() string {
	return s.String()
}

func (s *UpShelfLibraryRequest) GetMarketId() *int64 {
	return s.MarketId
}

func (s *UpShelfLibraryRequest) SetMarketId(v int64) *UpShelfLibraryRequest {
	s.MarketId = &v
	return s
}

func (s *UpShelfLibraryRequest) Validate() error {
	return dara.Validate(s)
}
