// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSAMLServiceProviderInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSAMLServiceProviderInfoResponseBody
	GetRequestId() *string
	SetSAMLServiceProviderInfo(v *GetSAMLServiceProviderInfoResponseBodySAMLServiceProviderInfo) *GetSAMLServiceProviderInfoResponseBody
	GetSAMLServiceProviderInfo() *GetSAMLServiceProviderInfoResponseBodySAMLServiceProviderInfo
}

type GetSAMLServiceProviderInfoResponseBody struct {
	// example:
	//
	// AABD6E03-4B3A-5687-88FF-72232670ED0C
	RequestId               *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SAMLServiceProviderInfo *GetSAMLServiceProviderInfoResponseBodySAMLServiceProviderInfo `json:"SAMLServiceProviderInfo,omitempty" xml:"SAMLServiceProviderInfo,omitempty" type:"Struct"`
}

func (s GetSAMLServiceProviderInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSAMLServiceProviderInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetSAMLServiceProviderInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSAMLServiceProviderInfoResponseBody) GetSAMLServiceProviderInfo() *GetSAMLServiceProviderInfoResponseBodySAMLServiceProviderInfo {
	return s.SAMLServiceProviderInfo
}

func (s *GetSAMLServiceProviderInfoResponseBody) SetRequestId(v string) *GetSAMLServiceProviderInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSAMLServiceProviderInfoResponseBody) SetSAMLServiceProviderInfo(v *GetSAMLServiceProviderInfoResponseBodySAMLServiceProviderInfo) *GetSAMLServiceProviderInfoResponseBody {
	s.SAMLServiceProviderInfo = v
	return s
}

func (s *GetSAMLServiceProviderInfoResponseBody) Validate() error {
	if s.SAMLServiceProviderInfo != nil {
		if err := s.SAMLServiceProviderInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSAMLServiceProviderInfoResponseBodySAMLServiceProviderInfo struct {
	// example:
	//
	// https://signin-cn-beijing.aliyunagentid.com/up_xxxxxxxxxxxxxxxxxxxx/saml/acs
	ACSURL *string `json:"ACSURL,omitempty" xml:"ACSURL,omitempty"`
	// example:
	//
	// https://signin-cn-beijing.aliyunagentid.com/up_xxxxxxxxxxxxxxxxxxxx
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// example:
	//
	// <?xml version=\\"1.0\\" encoding=\\"UTF-8\\" standalone=\\"no\\"?>\\n<md:EntityDescriptor entityID=\\"https://signin......
	SPMetadataDocument *string `json:"SPMetadataDocument,omitempty" xml:"SPMetadataDocument,omitempty"`
	// example:
	//
	// up_xxxxxxxxxxxxxxxxxxxx
	UserPoolId *string `json:"UserPoolId,omitempty" xml:"UserPoolId,omitempty"`
}

func (s GetSAMLServiceProviderInfoResponseBodySAMLServiceProviderInfo) String() string {
	return dara.Prettify(s)
}

func (s GetSAMLServiceProviderInfoResponseBodySAMLServiceProviderInfo) GoString() string {
	return s.String()
}

func (s *GetSAMLServiceProviderInfoResponseBodySAMLServiceProviderInfo) GetACSURL() *string {
	return s.ACSURL
}

func (s *GetSAMLServiceProviderInfoResponseBodySAMLServiceProviderInfo) GetEntityId() *string {
	return s.EntityId
}

func (s *GetSAMLServiceProviderInfoResponseBodySAMLServiceProviderInfo) GetSPMetadataDocument() *string {
	return s.SPMetadataDocument
}

func (s *GetSAMLServiceProviderInfoResponseBodySAMLServiceProviderInfo) GetUserPoolId() *string {
	return s.UserPoolId
}

func (s *GetSAMLServiceProviderInfoResponseBodySAMLServiceProviderInfo) SetACSURL(v string) *GetSAMLServiceProviderInfoResponseBodySAMLServiceProviderInfo {
	s.ACSURL = &v
	return s
}

func (s *GetSAMLServiceProviderInfoResponseBodySAMLServiceProviderInfo) SetEntityId(v string) *GetSAMLServiceProviderInfoResponseBodySAMLServiceProviderInfo {
	s.EntityId = &v
	return s
}

func (s *GetSAMLServiceProviderInfoResponseBodySAMLServiceProviderInfo) SetSPMetadataDocument(v string) *GetSAMLServiceProviderInfoResponseBodySAMLServiceProviderInfo {
	s.SPMetadataDocument = &v
	return s
}

func (s *GetSAMLServiceProviderInfoResponseBodySAMLServiceProviderInfo) SetUserPoolId(v string) *GetSAMLServiceProviderInfoResponseBodySAMLServiceProviderInfo {
	s.UserPoolId = &v
	return s
}

func (s *GetSAMLServiceProviderInfoResponseBodySAMLServiceProviderInfo) Validate() error {
	return dara.Validate(s)
}
