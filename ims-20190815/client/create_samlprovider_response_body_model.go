// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSAMLProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSAMLProviderResponseBody
	GetRequestId() *string
	SetSAMLProvider(v *CreateSAMLProviderResponseBodySAMLProvider) *CreateSAMLProviderResponseBody
	GetSAMLProvider() *CreateSAMLProviderResponseBodySAMLProvider
}

type CreateSAMLProviderResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A95A763D-F6B7-5242-83EB-AB45CE67F358
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the IdP.
	SAMLProvider *CreateSAMLProviderResponseBodySAMLProvider `json:"SAMLProvider,omitempty" xml:"SAMLProvider,omitempty" type:"Struct"`
}

func (s CreateSAMLProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSAMLProviderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSAMLProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSAMLProviderResponseBody) GetSAMLProvider() *CreateSAMLProviderResponseBodySAMLProvider {
	return s.SAMLProvider
}

func (s *CreateSAMLProviderResponseBody) SetRequestId(v string) *CreateSAMLProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSAMLProviderResponseBody) SetSAMLProvider(v *CreateSAMLProviderResponseBodySAMLProvider) *CreateSAMLProviderResponseBody {
	s.SAMLProvider = v
	return s
}

func (s *CreateSAMLProviderResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateSAMLProviderResponseBodySAMLProvider struct {
	// The Alibaba Cloud Resource Name (ARN) of the IdP.
	//
	// example:
	//
	// acs:ram::177242285274****:saml-provider/test-provider
	Arn           *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	AuthnSignAlgo *string `json:"AuthnSignAlgo,omitempty" xml:"AuthnSignAlgo,omitempty"`
	// The creation time. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-10-22T02:37:05Z
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
	// The update time. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-10-22T02:51:20Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s CreateSAMLProviderResponseBodySAMLProvider) String() string {
	return dara.Prettify(s)
}

func (s CreateSAMLProviderResponseBodySAMLProvider) GoString() string {
	return s.String()
}

func (s *CreateSAMLProviderResponseBodySAMLProvider) GetArn() *string {
	return s.Arn
}

func (s *CreateSAMLProviderResponseBodySAMLProvider) GetAuthnSignAlgo() *string {
	return s.AuthnSignAlgo
}

func (s *CreateSAMLProviderResponseBodySAMLProvider) GetCreateDate() *string {
	return s.CreateDate
}

func (s *CreateSAMLProviderResponseBodySAMLProvider) GetDescription() *string {
	return s.Description
}

func (s *CreateSAMLProviderResponseBodySAMLProvider) GetSAMLProviderName() *string {
	return s.SAMLProviderName
}

func (s *CreateSAMLProviderResponseBodySAMLProvider) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *CreateSAMLProviderResponseBodySAMLProvider) SetArn(v string) *CreateSAMLProviderResponseBodySAMLProvider {
	s.Arn = &v
	return s
}

func (s *CreateSAMLProviderResponseBodySAMLProvider) SetAuthnSignAlgo(v string) *CreateSAMLProviderResponseBodySAMLProvider {
	s.AuthnSignAlgo = &v
	return s
}

func (s *CreateSAMLProviderResponseBodySAMLProvider) SetCreateDate(v string) *CreateSAMLProviderResponseBodySAMLProvider {
	s.CreateDate = &v
	return s
}

func (s *CreateSAMLProviderResponseBodySAMLProvider) SetDescription(v string) *CreateSAMLProviderResponseBodySAMLProvider {
	s.Description = &v
	return s
}

func (s *CreateSAMLProviderResponseBodySAMLProvider) SetSAMLProviderName(v string) *CreateSAMLProviderResponseBodySAMLProvider {
	s.SAMLProviderName = &v
	return s
}

func (s *CreateSAMLProviderResponseBodySAMLProvider) SetUpdateDate(v string) *CreateSAMLProviderResponseBodySAMLProvider {
	s.UpdateDate = &v
	return s
}

func (s *CreateSAMLProviderResponseBodySAMLProvider) Validate() error {
	return dara.Validate(s)
}
