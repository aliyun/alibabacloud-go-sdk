// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSAMLProvidersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsTruncated(v bool) *ListSAMLProvidersResponseBody
	GetIsTruncated() *bool
	SetMarker(v string) *ListSAMLProvidersResponseBody
	GetMarker() *string
	SetRequestId(v string) *ListSAMLProvidersResponseBody
	GetRequestId() *string
	SetSAMLProviders(v *ListSAMLProvidersResponseBodySAMLProviders) *ListSAMLProvidersResponseBody
	GetSAMLProviders() *ListSAMLProvidersResponseBodySAMLProviders
}

type ListSAMLProvidersResponseBody struct {
	// Indicates whether the response is truncated. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The `marker`. This parameter is returned only if the value of `IsTruncated` is `true`. If the parameter is returned, you can call this operation again and set this parameter to obtain the truncated part.``
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2D8B70D3-E194-41C9-93C5-F6A10D716D24
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about IdPs.
	SAMLProviders *ListSAMLProvidersResponseBodySAMLProviders `json:"SAMLProviders,omitempty" xml:"SAMLProviders,omitempty" type:"Struct"`
}

func (s ListSAMLProvidersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSAMLProvidersResponseBody) GoString() string {
	return s.String()
}

func (s *ListSAMLProvidersResponseBody) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListSAMLProvidersResponseBody) GetMarker() *string {
	return s.Marker
}

func (s *ListSAMLProvidersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSAMLProvidersResponseBody) GetSAMLProviders() *ListSAMLProvidersResponseBodySAMLProviders {
	return s.SAMLProviders
}

func (s *ListSAMLProvidersResponseBody) SetIsTruncated(v bool) *ListSAMLProvidersResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListSAMLProvidersResponseBody) SetMarker(v string) *ListSAMLProvidersResponseBody {
	s.Marker = &v
	return s
}

func (s *ListSAMLProvidersResponseBody) SetRequestId(v string) *ListSAMLProvidersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSAMLProvidersResponseBody) SetSAMLProviders(v *ListSAMLProvidersResponseBodySAMLProviders) *ListSAMLProvidersResponseBody {
	s.SAMLProviders = v
	return s
}

func (s *ListSAMLProvidersResponseBody) Validate() error {
	if s.SAMLProviders != nil {
		if err := s.SAMLProviders.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSAMLProvidersResponseBodySAMLProviders struct {
	SAMLProvider []*ListSAMLProvidersResponseBodySAMLProvidersSAMLProvider `json:"SAMLProvider,omitempty" xml:"SAMLProvider,omitempty" type:"Repeated"`
}

func (s ListSAMLProvidersResponseBodySAMLProviders) String() string {
	return dara.Prettify(s)
}

func (s ListSAMLProvidersResponseBodySAMLProviders) GoString() string {
	return s.String()
}

func (s *ListSAMLProvidersResponseBodySAMLProviders) GetSAMLProvider() []*ListSAMLProvidersResponseBodySAMLProvidersSAMLProvider {
	return s.SAMLProvider
}

func (s *ListSAMLProvidersResponseBodySAMLProviders) SetSAMLProvider(v []*ListSAMLProvidersResponseBodySAMLProvidersSAMLProvider) *ListSAMLProvidersResponseBodySAMLProviders {
	s.SAMLProvider = v
	return s
}

func (s *ListSAMLProvidersResponseBodySAMLProviders) Validate() error {
	if s.SAMLProvider != nil {
		for _, item := range s.SAMLProvider {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSAMLProvidersResponseBodySAMLProvidersSAMLProvider struct {
	// The Alibaba Cloud Resource Name (ARN) of the IdP.
	//
	// example:
	//
	// acs:ram::177242285274****:saml-provider/test-provider
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2020-10-22T06:26:15Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description.
	//
	// example:
	//
	// This is a provider.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the IdP.
	//
	// example:
	//
	// test-provider
	SAMLProviderName *string `json:"SAMLProviderName,omitempty" xml:"SAMLProviderName,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2020-10-22T06:26:15Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s ListSAMLProvidersResponseBodySAMLProvidersSAMLProvider) String() string {
	return dara.Prettify(s)
}

func (s ListSAMLProvidersResponseBodySAMLProvidersSAMLProvider) GoString() string {
	return s.String()
}

func (s *ListSAMLProvidersResponseBodySAMLProvidersSAMLProvider) GetArn() *string {
	return s.Arn
}

func (s *ListSAMLProvidersResponseBodySAMLProvidersSAMLProvider) GetCreateDate() *string {
	return s.CreateDate
}

func (s *ListSAMLProvidersResponseBodySAMLProvidersSAMLProvider) GetDescription() *string {
	return s.Description
}

func (s *ListSAMLProvidersResponseBodySAMLProvidersSAMLProvider) GetSAMLProviderName() *string {
	return s.SAMLProviderName
}

func (s *ListSAMLProvidersResponseBodySAMLProvidersSAMLProvider) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *ListSAMLProvidersResponseBodySAMLProvidersSAMLProvider) SetArn(v string) *ListSAMLProvidersResponseBodySAMLProvidersSAMLProvider {
	s.Arn = &v
	return s
}

func (s *ListSAMLProvidersResponseBodySAMLProvidersSAMLProvider) SetCreateDate(v string) *ListSAMLProvidersResponseBodySAMLProvidersSAMLProvider {
	s.CreateDate = &v
	return s
}

func (s *ListSAMLProvidersResponseBodySAMLProvidersSAMLProvider) SetDescription(v string) *ListSAMLProvidersResponseBodySAMLProvidersSAMLProvider {
	s.Description = &v
	return s
}

func (s *ListSAMLProvidersResponseBodySAMLProvidersSAMLProvider) SetSAMLProviderName(v string) *ListSAMLProvidersResponseBodySAMLProvidersSAMLProvider {
	s.SAMLProviderName = &v
	return s
}

func (s *ListSAMLProvidersResponseBodySAMLProvidersSAMLProvider) SetUpdateDate(v string) *ListSAMLProvidersResponseBodySAMLProvidersSAMLProvider {
	s.UpdateDate = &v
	return s
}

func (s *ListSAMLProvidersResponseBodySAMLProvidersSAMLProvider) Validate() error {
	return dara.Validate(s)
}
