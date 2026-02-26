// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCategoryListResult interface {
	dara.Model
	String() string
	GoString() string
	SetCategories(v []*Category) *CategoryListResult
	GetCategories() []*Category
	SetRequestId(v string) *CategoryListResult
	GetRequestId() *string
}

type CategoryListResult struct {
	Categories []*Category `json:"categories,omitempty" xml:"categories,omitempty" type:"Repeated"`
	RequestId  *string     `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CategoryListResult) String() string {
	return dara.Prettify(s)
}

func (s CategoryListResult) GoString() string {
	return s.String()
}

func (s *CategoryListResult) GetCategories() []*Category {
	return s.Categories
}

func (s *CategoryListResult) GetRequestId() *string {
	return s.RequestId
}

func (s *CategoryListResult) SetCategories(v []*Category) *CategoryListResult {
	s.Categories = v
	return s
}

func (s *CategoryListResult) SetRequestId(v string) *CategoryListResult {
	s.RequestId = &v
	return s
}

func (s *CategoryListResult) Validate() error {
	if s.Categories != nil {
		for _, item := range s.Categories {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
