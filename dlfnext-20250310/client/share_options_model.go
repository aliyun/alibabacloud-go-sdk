// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iShareOptions interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogSizeLimit(v int32) *ShareOptions
	GetCatalogSizeLimit() *int32
	SetReceiverSizeLimit(v int32) *ShareOptions
	GetReceiverSizeLimit() *int32
	SetShareResourceSizeLimit(v int32) *ShareOptions
	GetShareResourceSizeLimit() *int32
	SetShareSizeLimit(v int32) *ShareOptions
	GetShareSizeLimit() *int32
}

type ShareOptions struct {
	CatalogSizeLimit       *int32 `json:"catalogSizeLimit,omitempty" xml:"catalogSizeLimit,omitempty"`
	ReceiverSizeLimit      *int32 `json:"receiverSizeLimit,omitempty" xml:"receiverSizeLimit,omitempty"`
	ShareResourceSizeLimit *int32 `json:"shareResourceSizeLimit,omitempty" xml:"shareResourceSizeLimit,omitempty"`
	ShareSizeLimit         *int32 `json:"shareSizeLimit,omitempty" xml:"shareSizeLimit,omitempty"`
}

func (s ShareOptions) String() string {
	return dara.Prettify(s)
}

func (s ShareOptions) GoString() string {
	return s.String()
}

func (s *ShareOptions) GetCatalogSizeLimit() *int32 {
	return s.CatalogSizeLimit
}

func (s *ShareOptions) GetReceiverSizeLimit() *int32 {
	return s.ReceiverSizeLimit
}

func (s *ShareOptions) GetShareResourceSizeLimit() *int32 {
	return s.ShareResourceSizeLimit
}

func (s *ShareOptions) GetShareSizeLimit() *int32 {
	return s.ShareSizeLimit
}

func (s *ShareOptions) SetCatalogSizeLimit(v int32) *ShareOptions {
	s.CatalogSizeLimit = &v
	return s
}

func (s *ShareOptions) SetReceiverSizeLimit(v int32) *ShareOptions {
	s.ReceiverSizeLimit = &v
	return s
}

func (s *ShareOptions) SetShareResourceSizeLimit(v int32) *ShareOptions {
	s.ShareResourceSizeLimit = &v
	return s
}

func (s *ShareOptions) SetShareSizeLimit(v int32) *ShareOptions {
	s.ShareSizeLimit = &v
	return s
}

func (s *ShareOptions) Validate() error {
	return dara.Validate(s)
}
