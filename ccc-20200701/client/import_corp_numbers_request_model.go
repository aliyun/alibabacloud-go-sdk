// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportCorpNumbersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCity(v string) *ImportCorpNumbersRequest
	GetCity() *string
	SetCorpName(v string) *ImportCorpNumbersRequest
	GetCorpName() *string
	SetNumberList(v string) *ImportCorpNumbersRequest
	GetNumberList() *string
	SetProvider(v string) *ImportCorpNumbersRequest
	GetProvider() *string
	SetProvince(v string) *ImportCorpNumbersRequest
	GetProvince() *string
	SetTagList(v string) *ImportCorpNumbersRequest
	GetTagList() *string
}

type ImportCorpNumbersRequest struct {
	// The city.
	//
	// example:
	//
	// 西安
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// The enterprise name.
	//
	// example:
	//
	// xxx公司
	CorpName *string `json:"CorpName,omitempty" xml:"CorpName,omitempty"`
	// The list of numbers.
	//
	// This parameter is required.
	//
	// example:
	//
	// [\\"029xxxxxx\\"]
	NumberList *string `json:"NumberList,omitempty" xml:"NumberList,omitempty"`
	// The line information.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx通信-陕西线路
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// The province.
	//
	// example:
	//
	// 陕西
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// The business tags.
	//
	// example:
	//
	// 无
	TagList *string `json:"TagList,omitempty" xml:"TagList,omitempty"`
}

func (s ImportCorpNumbersRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportCorpNumbersRequest) GoString() string {
	return s.String()
}

func (s *ImportCorpNumbersRequest) GetCity() *string {
	return s.City
}

func (s *ImportCorpNumbersRequest) GetCorpName() *string {
	return s.CorpName
}

func (s *ImportCorpNumbersRequest) GetNumberList() *string {
	return s.NumberList
}

func (s *ImportCorpNumbersRequest) GetProvider() *string {
	return s.Provider
}

func (s *ImportCorpNumbersRequest) GetProvince() *string {
	return s.Province
}

func (s *ImportCorpNumbersRequest) GetTagList() *string {
	return s.TagList
}

func (s *ImportCorpNumbersRequest) SetCity(v string) *ImportCorpNumbersRequest {
	s.City = &v
	return s
}

func (s *ImportCorpNumbersRequest) SetCorpName(v string) *ImportCorpNumbersRequest {
	s.CorpName = &v
	return s
}

func (s *ImportCorpNumbersRequest) SetNumberList(v string) *ImportCorpNumbersRequest {
	s.NumberList = &v
	return s
}

func (s *ImportCorpNumbersRequest) SetProvider(v string) *ImportCorpNumbersRequest {
	s.Provider = &v
	return s
}

func (s *ImportCorpNumbersRequest) SetProvince(v string) *ImportCorpNumbersRequest {
	s.Province = &v
	return s
}

func (s *ImportCorpNumbersRequest) SetTagList(v string) *ImportCorpNumbersRequest {
	s.TagList = &v
	return s
}

func (s *ImportCorpNumbersRequest) Validate() error {
	return dara.Validate(s)
}
