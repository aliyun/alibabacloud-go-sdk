// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeProductListResult interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizedProductList(v []*ProductInfo) *AuthorizeProductListResult
	GetAuthorizedProductList() []*ProductInfo
	SetRequestId(v string) *AuthorizeProductListResult
	GetRequestId() *string
	SetTotal(v int64) *AuthorizeProductListResult
	GetTotal() *int64
}

type AuthorizeProductListResult struct {
	AuthorizedProductList []*ProductInfo `json:"authorizedProductList,omitempty" xml:"authorizedProductList,omitempty" type:"Repeated"`
	RequestId             *string        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Total                 *int64         `json:"total,omitempty" xml:"total,omitempty"`
}

func (s AuthorizeProductListResult) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeProductListResult) GoString() string {
	return s.String()
}

func (s *AuthorizeProductListResult) GetAuthorizedProductList() []*ProductInfo {
	return s.AuthorizedProductList
}

func (s *AuthorizeProductListResult) GetRequestId() *string {
	return s.RequestId
}

func (s *AuthorizeProductListResult) GetTotal() *int64 {
	return s.Total
}

func (s *AuthorizeProductListResult) SetAuthorizedProductList(v []*ProductInfo) *AuthorizeProductListResult {
	s.AuthorizedProductList = v
	return s
}

func (s *AuthorizeProductListResult) SetRequestId(v string) *AuthorizeProductListResult {
	s.RequestId = &v
	return s
}

func (s *AuthorizeProductListResult) SetTotal(v int64) *AuthorizeProductListResult {
	s.Total = &v
	return s
}

func (s *AuthorizeProductListResult) Validate() error {
	if s.AuthorizedProductList != nil {
		for _, item := range s.AuthorizedProductList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
