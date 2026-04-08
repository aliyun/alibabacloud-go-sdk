// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMmsDataSourceConfigItemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListMmsDataSourceConfigItemsRequest
	GetLang() *string
	SetSourceType(v string) *ListMmsDataSourceConfigItemsRequest
	GetSourceType() *string
}

type ListMmsDataSourceConfigItemsRequest struct {
	// example:
	//
	// zh_CN
	Lang *string `json:"lang,omitempty" xml:"lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Hive
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
}

func (s ListMmsDataSourceConfigItemsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMmsDataSourceConfigItemsRequest) GoString() string {
	return s.String()
}

func (s *ListMmsDataSourceConfigItemsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListMmsDataSourceConfigItemsRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *ListMmsDataSourceConfigItemsRequest) SetLang(v string) *ListMmsDataSourceConfigItemsRequest {
	s.Lang = &v
	return s
}

func (s *ListMmsDataSourceConfigItemsRequest) SetSourceType(v string) *ListMmsDataSourceConfigItemsRequest {
	s.SourceType = &v
	return s
}

func (s *ListMmsDataSourceConfigItemsRequest) Validate() error {
	return dara.Validate(s)
}
