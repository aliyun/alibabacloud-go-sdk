// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHiMarketProductPublicationConifg interface {
	dara.Model
	String() string
	GoString() string
	SetPublicationId(v string) *HiMarketProductPublicationConifg
	GetPublicationId() *string
}

type HiMarketProductPublicationConifg struct {
	PublicationId *string `json:"publicationId,omitempty" xml:"publicationId,omitempty"`
}

func (s HiMarketProductPublicationConifg) String() string {
	return dara.Prettify(s)
}

func (s HiMarketProductPublicationConifg) GoString() string {
	return s.String()
}

func (s *HiMarketProductPublicationConifg) GetPublicationId() *string {
	return s.PublicationId
}

func (s *HiMarketProductPublicationConifg) SetPublicationId(v string) *HiMarketProductPublicationConifg {
	s.PublicationId = &v
	return s
}

func (s *HiMarketProductPublicationConifg) Validate() error {
	return dara.Validate(s)
}
