// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSAMLProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSAMLProviderResponseBody
	GetRequestId() *string
	SetSAMLProvider(v *UpdateSAMLProviderResponseBodySAMLProvider) *UpdateSAMLProviderResponseBody
	GetSAMLProvider() *UpdateSAMLProviderResponseBodySAMLProvider
}

type UpdateSAMLProviderResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E5EDDFD2-3654-4F9F-9780-4AE7D81823EF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the identity provider.
	SAMLProvider *UpdateSAMLProviderResponseBodySAMLProvider `json:"SAMLProvider,omitempty" xml:"SAMLProvider,omitempty" type:"Struct"`
}

func (s UpdateSAMLProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSAMLProviderResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSAMLProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSAMLProviderResponseBody) GetSAMLProvider() *UpdateSAMLProviderResponseBodySAMLProvider {
	return s.SAMLProvider
}

func (s *UpdateSAMLProviderResponseBody) SetRequestId(v string) *UpdateSAMLProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSAMLProviderResponseBody) SetSAMLProvider(v *UpdateSAMLProviderResponseBodySAMLProvider) *UpdateSAMLProviderResponseBody {
	s.SAMLProvider = v
	return s
}

func (s *UpdateSAMLProviderResponseBody) Validate() error {
	if s.SAMLProvider != nil {
		if err := s.SAMLProvider.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateSAMLProviderResponseBodySAMLProvider struct {
	// The Alibaba Cloud Resource Name (ARN) of the identity provider.
	//
	// example:
	//
	// acs:ram::177242285274****:saml-provider/test-provider
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The signature algorithm supported by the Alibaba Cloud SP. Valid values:
	//
	// - rsa-sha256
	//
	// - rsa-sha1
	//
	// example:
	//
	// rsa-sha1
	AuthnSignAlgo *string `json:"AuthnSignAlgo,omitempty" xml:"AuthnSignAlgo,omitempty"`
	// The time when the identity provider was created. The time is in UTC.
	//
	// example:
	//
	// 2020-10-22T02:37:05Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description.
	//
	// example:
	//
	// This is a new provider.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the identity provider.
	//
	// example:
	//
	// test-provider
	SAMLProviderName *string `json:"SAMLProviderName,omitempty" xml:"SAMLProviderName,omitempty"`
	// The time when the identity provider was last updated. The time is in Coordinated Universal Time (UTC).
	//
	// example:
	//
	// 2020-10-22T02:51:20Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s UpdateSAMLProviderResponseBodySAMLProvider) String() string {
	return dara.Prettify(s)
}

func (s UpdateSAMLProviderResponseBodySAMLProvider) GoString() string {
	return s.String()
}

func (s *UpdateSAMLProviderResponseBodySAMLProvider) GetArn() *string {
	return s.Arn
}

func (s *UpdateSAMLProviderResponseBodySAMLProvider) GetAuthnSignAlgo() *string {
	return s.AuthnSignAlgo
}

func (s *UpdateSAMLProviderResponseBodySAMLProvider) GetCreateDate() *string {
	return s.CreateDate
}

func (s *UpdateSAMLProviderResponseBodySAMLProvider) GetDescription() *string {
	return s.Description
}

func (s *UpdateSAMLProviderResponseBodySAMLProvider) GetSAMLProviderName() *string {
	return s.SAMLProviderName
}

func (s *UpdateSAMLProviderResponseBodySAMLProvider) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *UpdateSAMLProviderResponseBodySAMLProvider) SetArn(v string) *UpdateSAMLProviderResponseBodySAMLProvider {
	s.Arn = &v
	return s
}

func (s *UpdateSAMLProviderResponseBodySAMLProvider) SetAuthnSignAlgo(v string) *UpdateSAMLProviderResponseBodySAMLProvider {
	s.AuthnSignAlgo = &v
	return s
}

func (s *UpdateSAMLProviderResponseBodySAMLProvider) SetCreateDate(v string) *UpdateSAMLProviderResponseBodySAMLProvider {
	s.CreateDate = &v
	return s
}

func (s *UpdateSAMLProviderResponseBodySAMLProvider) SetDescription(v string) *UpdateSAMLProviderResponseBodySAMLProvider {
	s.Description = &v
	return s
}

func (s *UpdateSAMLProviderResponseBodySAMLProvider) SetSAMLProviderName(v string) *UpdateSAMLProviderResponseBodySAMLProvider {
	s.SAMLProviderName = &v
	return s
}

func (s *UpdateSAMLProviderResponseBodySAMLProvider) SetUpdateDate(v string) *UpdateSAMLProviderResponseBodySAMLProvider {
	s.UpdateDate = &v
	return s
}

func (s *UpdateSAMLProviderResponseBodySAMLProvider) Validate() error {
	return dara.Validate(s)
}
