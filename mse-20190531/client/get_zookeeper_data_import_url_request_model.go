// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetZookeeperDataImportUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GetZookeeperDataImportUrlRequest
	GetAcceptLanguage() *string
	SetContentType(v string) *GetZookeeperDataImportUrlRequest
	GetContentType() *string
	SetInstanceId(v string) *GetZookeeperDataImportUrlRequest
	GetInstanceId() *string
}

type GetZookeeperDataImportUrlRequest struct {
	// RestResult
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zip
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The type of the file.
	//
	// example:
	//
	// mse_prepaid_public_cn-zvp2xzzkk06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetZookeeperDataImportUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetZookeeperDataImportUrlRequest) GoString() string {
	return s.String()
}

func (s *GetZookeeperDataImportUrlRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GetZookeeperDataImportUrlRequest) GetContentType() *string {
	return s.ContentType
}

func (s *GetZookeeperDataImportUrlRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetZookeeperDataImportUrlRequest) SetAcceptLanguage(v string) *GetZookeeperDataImportUrlRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetZookeeperDataImportUrlRequest) SetContentType(v string) *GetZookeeperDataImportUrlRequest {
	s.ContentType = &v
	return s
}

func (s *GetZookeeperDataImportUrlRequest) SetInstanceId(v string) *GetZookeeperDataImportUrlRequest {
	s.InstanceId = &v
	return s
}

func (s *GetZookeeperDataImportUrlRequest) Validate() error {
	return dara.Validate(s)
}
