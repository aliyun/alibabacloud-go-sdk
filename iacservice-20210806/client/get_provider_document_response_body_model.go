// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProviderDocumentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDocument(v string) *GetProviderDocumentResponseBody
	GetDocument() *string
	SetProviderVersion(v string) *GetProviderDocumentResponseBody
	GetProviderVersion() *string
	SetRequestId(v string) *GetProviderDocumentResponseBody
	GetRequestId() *string
	SetTerraformResourceType(v string) *GetProviderDocumentResponseBody
	GetTerraformResourceType() *string
}

type GetProviderDocumentResponseBody struct {
	Document        *string `json:"document,omitempty" xml:"document,omitempty"`
	ProviderVersion *string `json:"providerVersion,omitempty" xml:"providerVersion,omitempty"`
	// Id of the request
	RequestId             *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TerraformResourceType *string `json:"terraformResourceType,omitempty" xml:"terraformResourceType,omitempty"`
}

func (s GetProviderDocumentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProviderDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *GetProviderDocumentResponseBody) GetDocument() *string {
	return s.Document
}

func (s *GetProviderDocumentResponseBody) GetProviderVersion() *string {
	return s.ProviderVersion
}

func (s *GetProviderDocumentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProviderDocumentResponseBody) GetTerraformResourceType() *string {
	return s.TerraformResourceType
}

func (s *GetProviderDocumentResponseBody) SetDocument(v string) *GetProviderDocumentResponseBody {
	s.Document = &v
	return s
}

func (s *GetProviderDocumentResponseBody) SetProviderVersion(v string) *GetProviderDocumentResponseBody {
	s.ProviderVersion = &v
	return s
}

func (s *GetProviderDocumentResponseBody) SetRequestId(v string) *GetProviderDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProviderDocumentResponseBody) SetTerraformResourceType(v string) *GetProviderDocumentResponseBody {
	s.TerraformResourceType = &v
	return s
}

func (s *GetProviderDocumentResponseBody) Validate() error {
	return dara.Validate(s)
}
