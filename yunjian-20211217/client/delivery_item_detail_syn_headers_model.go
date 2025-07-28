// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeliveryItemDetailSynHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeliveryItemDetailSynHeaders
	GetCommonHeaders() map[string]*string
	SetYunUserId(v string) *DeliveryItemDetailSynHeaders
	GetYunUserId() *string
}

type DeliveryItemDetailSynHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	YunUserId *string `json:"Yun-User-Id,omitempty" xml:"Yun-User-Id,omitempty"`
}

func (s DeliveryItemDetailSynHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeliveryItemDetailSynHeaders) GoString() string {
	return s.String()
}

func (s *DeliveryItemDetailSynHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeliveryItemDetailSynHeaders) GetYunUserId() *string {
	return s.YunUserId
}

func (s *DeliveryItemDetailSynHeaders) SetCommonHeaders(v map[string]*string) *DeliveryItemDetailSynHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeliveryItemDetailSynHeaders) SetYunUserId(v string) *DeliveryItemDetailSynHeaders {
	s.YunUserId = &v
	return s
}

func (s *DeliveryItemDetailSynHeaders) Validate() error {
	return dara.Validate(s)
}
