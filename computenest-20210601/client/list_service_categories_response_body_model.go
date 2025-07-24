// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceCategoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCategories(v []*string) *ListServiceCategoriesResponseBody
	GetCategories() []*string
	SetRequestId(v string) *ListServiceCategoriesResponseBody
	GetRequestId() *string
}

type ListServiceCategoriesResponseBody struct {
	// The category list of the service.
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 13FE89A5-C036-56BF-A0FF-A31C59819FD7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListServiceCategoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServiceCategoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceCategoriesResponseBody) GetCategories() []*string {
	return s.Categories
}

func (s *ListServiceCategoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServiceCategoriesResponseBody) SetCategories(v []*string) *ListServiceCategoriesResponseBody {
	s.Categories = v
	return s
}

func (s *ListServiceCategoriesResponseBody) SetRequestId(v string) *ListServiceCategoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceCategoriesResponseBody) Validate() error {
	return dara.Validate(s)
}
