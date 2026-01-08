// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRequestWhatsappConversionApiShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *RequestWhatsappConversionApiShrinkRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *RequestWhatsappConversionApiShrinkRequest
	GetOwnerId() *int64
	SetPageId(v string) *RequestWhatsappConversionApiShrinkRequest
	GetPageId() *string
	SetRequestDataShrink(v string) *RequestWhatsappConversionApiShrinkRequest
	GetRequestDataShrink() *string
	SetResourceOwnerAccount(v string) *RequestWhatsappConversionApiShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RequestWhatsappConversionApiShrinkRequest
	GetResourceOwnerId() *int64
}

type RequestWhatsappConversionApiShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 929399382
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1939848838
	PageId               *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	RequestDataShrink    *string `json:"RequestData,omitempty" xml:"RequestData,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RequestWhatsappConversionApiShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RequestWhatsappConversionApiShrinkRequest) GoString() string {
	return s.String()
}

func (s *RequestWhatsappConversionApiShrinkRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *RequestWhatsappConversionApiShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RequestWhatsappConversionApiShrinkRequest) GetPageId() *string {
	return s.PageId
}

func (s *RequestWhatsappConversionApiShrinkRequest) GetRequestDataShrink() *string {
	return s.RequestDataShrink
}

func (s *RequestWhatsappConversionApiShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RequestWhatsappConversionApiShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RequestWhatsappConversionApiShrinkRequest) SetCustSpaceId(v string) *RequestWhatsappConversionApiShrinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *RequestWhatsappConversionApiShrinkRequest) SetOwnerId(v int64) *RequestWhatsappConversionApiShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *RequestWhatsappConversionApiShrinkRequest) SetPageId(v string) *RequestWhatsappConversionApiShrinkRequest {
	s.PageId = &v
	return s
}

func (s *RequestWhatsappConversionApiShrinkRequest) SetRequestDataShrink(v string) *RequestWhatsappConversionApiShrinkRequest {
	s.RequestDataShrink = &v
	return s
}

func (s *RequestWhatsappConversionApiShrinkRequest) SetResourceOwnerAccount(v string) *RequestWhatsappConversionApiShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RequestWhatsappConversionApiShrinkRequest) SetResourceOwnerId(v int64) *RequestWhatsappConversionApiShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RequestWhatsappConversionApiShrinkRequest) Validate() error {
	return dara.Validate(s)
}
