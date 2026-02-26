// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGoodsShippingNoticeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *GoodsShippingNoticeCreateCmd) *CreateGoodsShippingNoticeRequest
	GetBody() *GoodsShippingNoticeCreateCmd
}

type CreateGoodsShippingNoticeRequest struct {
	// This parameter is required.
	Body *GoodsShippingNoticeCreateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGoodsShippingNoticeRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGoodsShippingNoticeRequest) GoString() string {
	return s.String()
}

func (s *CreateGoodsShippingNoticeRequest) GetBody() *GoodsShippingNoticeCreateCmd {
	return s.Body
}

func (s *CreateGoodsShippingNoticeRequest) SetBody(v *GoodsShippingNoticeCreateCmd) *CreateGoodsShippingNoticeRequest {
	s.Body = v
	return s
}

func (s *CreateGoodsShippingNoticeRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
