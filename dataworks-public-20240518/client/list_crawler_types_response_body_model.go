// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCrawlerTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCrawlerTypes(v []*CrawlerType) *ListCrawlerTypesResponseBody
	GetCrawlerTypes() []*CrawlerType
	SetRequestId(v string) *ListCrawlerTypesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCrawlerTypesResponseBody
	GetSuccess() *bool
}

type ListCrawlerTypesResponseBody struct {
	CrawlerTypes []*CrawlerType `json:"CrawlerTypes,omitempty" xml:"CrawlerTypes,omitempty" type:"Repeated"`
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCrawlerTypesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCrawlerTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCrawlerTypesResponseBody) GetCrawlerTypes() []*CrawlerType {
	return s.CrawlerTypes
}

func (s *ListCrawlerTypesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCrawlerTypesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCrawlerTypesResponseBody) SetCrawlerTypes(v []*CrawlerType) *ListCrawlerTypesResponseBody {
	s.CrawlerTypes = v
	return s
}

func (s *ListCrawlerTypesResponseBody) SetRequestId(v string) *ListCrawlerTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCrawlerTypesResponseBody) SetSuccess(v bool) *ListCrawlerTypesResponseBody {
	s.Success = &v
	return s
}

func (s *ListCrawlerTypesResponseBody) Validate() error {
	return dara.Validate(s)
}
