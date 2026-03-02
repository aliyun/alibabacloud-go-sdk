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
	IsTruncated   *bool                                       `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	Marker        *string                                     `json:"Marker,omitempty" xml:"Marker,omitempty"`
	RequestId     *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Arn              *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	CreateDate       *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	SAMLProviderName *string `json:"SAMLProviderName,omitempty" xml:"SAMLProviderName,omitempty"`
	UpdateDate       *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
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
