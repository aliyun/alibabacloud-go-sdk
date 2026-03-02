// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpShelfPbcVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMarketId(v int64) *UpShelfPbcVersionRequest
	GetMarketId() *int64
}

type UpShelfPbcVersionRequest struct {
	// This parameter is required.
	MarketId *int64 `json:"marketId,omitempty" xml:"marketId,omitempty"`
}

func (s UpShelfPbcVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpShelfPbcVersionRequest) GoString() string {
	return s.String()
}

func (s *UpShelfPbcVersionRequest) GetMarketId() *int64 {
	return s.MarketId
}

func (s *UpShelfPbcVersionRequest) SetMarketId(v int64) *UpShelfPbcVersionRequest {
	s.MarketId = &v
	return s
}

func (s *UpShelfPbcVersionRequest) Validate() error {
	return dara.Validate(s)
}
