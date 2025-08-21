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
	// The information about the IdP.
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
	return dara.Validate(s)
}

type UpdateSAMLProviderResponseBodySAMLProvider struct {
	// The Alibaba Cloud Resource Name (ARN) of the IdP.
	//
	// example:
	//
	// acs:ram::177242285274****:saml-provider/test-provider
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The point in time at which the IdP was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-10-22T02:37:05Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the IdP.
	//
	// example:
	//
	// This is a new provider.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the IdP.
	//
	// example:
	//
	// test-provider
	SAMLProviderName *string `json:"SAMLProviderName,omitempty" xml:"SAMLProviderName,omitempty"`
	// The point in time at which the information about the IdP was modified. The time is displayed in UTC.
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
