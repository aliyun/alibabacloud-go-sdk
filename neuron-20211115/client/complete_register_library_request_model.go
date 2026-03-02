// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompleteRegisterLibraryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDependIntegral(v int32) *CompleteRegisterLibraryRequest
	GetDependIntegral() *int32
	SetMarketId(v int64) *CompleteRegisterLibraryRequest
	GetMarketId() *int64
}

type CompleteRegisterLibraryRequest struct {
	DependIntegral *int32 `json:"dependIntegral,omitempty" xml:"dependIntegral,omitempty"`
	MarketId       *int64 `json:"marketId,omitempty" xml:"marketId,omitempty"`
}

func (s CompleteRegisterLibraryRequest) String() string {
	return dara.Prettify(s)
}

func (s CompleteRegisterLibraryRequest) GoString() string {
	return s.String()
}

func (s *CompleteRegisterLibraryRequest) GetDependIntegral() *int32 {
	return s.DependIntegral
}

func (s *CompleteRegisterLibraryRequest) GetMarketId() *int64 {
	return s.MarketId
}

func (s *CompleteRegisterLibraryRequest) SetDependIntegral(v int32) *CompleteRegisterLibraryRequest {
	s.DependIntegral = &v
	return s
}

func (s *CompleteRegisterLibraryRequest) SetMarketId(v int64) *CompleteRegisterLibraryRequest {
	s.MarketId = &v
	return s
}

func (s *CompleteRegisterLibraryRequest) Validate() error {
	return dara.Validate(s)
}
