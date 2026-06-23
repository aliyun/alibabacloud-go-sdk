// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommercializeFetchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelId(v string) *CommercializeFetchRequest
	GetChannelId() *string
	SetData(v string) *CommercializeFetchRequest
	GetData() *string
	SetProductId(v string) *CommercializeFetchRequest
	GetProductId() *string
	SetRequestId(v string) *CommercializeFetchRequest
	GetRequestId() *string
	SetSecretKey(v string) *CommercializeFetchRequest
	GetSecretKey() *string
	SetSign(v string) *CommercializeFetchRequest
	GetSign() *string
}

type CommercializeFetchRequest struct {
	// This parameter is required.
	ChannelId *string `json:"channelId,omitempty" xml:"channelId,omitempty"`
	// This parameter is required.
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// This parameter is required.
	ProductId *string `json:"productId,omitempty" xml:"productId,omitempty"`
	// This parameter is required.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	SecretKey *string `json:"secretKey,omitempty" xml:"secretKey,omitempty"`
	// This parameter is required.
	Sign *string `json:"sign,omitempty" xml:"sign,omitempty"`
}

func (s CommercializeFetchRequest) String() string {
	return dara.Prettify(s)
}

func (s CommercializeFetchRequest) GoString() string {
	return s.String()
}

func (s *CommercializeFetchRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *CommercializeFetchRequest) GetData() *string {
	return s.Data
}

func (s *CommercializeFetchRequest) GetProductId() *string {
	return s.ProductId
}

func (s *CommercializeFetchRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *CommercializeFetchRequest) GetSecretKey() *string {
	return s.SecretKey
}

func (s *CommercializeFetchRequest) GetSign() *string {
	return s.Sign
}

func (s *CommercializeFetchRequest) SetChannelId(v string) *CommercializeFetchRequest {
	s.ChannelId = &v
	return s
}

func (s *CommercializeFetchRequest) SetData(v string) *CommercializeFetchRequest {
	s.Data = &v
	return s
}

func (s *CommercializeFetchRequest) SetProductId(v string) *CommercializeFetchRequest {
	s.ProductId = &v
	return s
}

func (s *CommercializeFetchRequest) SetRequestId(v string) *CommercializeFetchRequest {
	s.RequestId = &v
	return s
}

func (s *CommercializeFetchRequest) SetSecretKey(v string) *CommercializeFetchRequest {
	s.SecretKey = &v
	return s
}

func (s *CommercializeFetchRequest) SetSign(v string) *CommercializeFetchRequest {
	s.Sign = &v
	return s
}

func (s *CommercializeFetchRequest) Validate() error {
	return dara.Validate(s)
}
