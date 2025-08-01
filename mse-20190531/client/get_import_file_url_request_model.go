// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImportFileUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GetImportFileUrlRequest
	GetAcceptLanguage() *string
	SetContentType(v string) *GetImportFileUrlRequest
	GetContentType() *string
	SetInstanceId(v string) *GetImportFileUrlRequest
	GetInstanceId() *string
	SetNamespaceId(v string) *GetImportFileUrlRequest
	GetNamespaceId() *string
}

type GetImportFileUrlRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The file type.
	//
	// example:
	//
	// application/zip
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// mse_prepaid_public_cn-7pp2b****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// 12788f73-9848-4388-98f1-507778f2****
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s GetImportFileUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetImportFileUrlRequest) GoString() string {
	return s.String()
}

func (s *GetImportFileUrlRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GetImportFileUrlRequest) GetContentType() *string {
	return s.ContentType
}

func (s *GetImportFileUrlRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetImportFileUrlRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *GetImportFileUrlRequest) SetAcceptLanguage(v string) *GetImportFileUrlRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetImportFileUrlRequest) SetContentType(v string) *GetImportFileUrlRequest {
	s.ContentType = &v
	return s
}

func (s *GetImportFileUrlRequest) SetInstanceId(v string) *GetImportFileUrlRequest {
	s.InstanceId = &v
	return s
}

func (s *GetImportFileUrlRequest) SetNamespaceId(v string) *GetImportFileUrlRequest {
	s.NamespaceId = &v
	return s
}

func (s *GetImportFileUrlRequest) Validate() error {
	return dara.Validate(s)
}
