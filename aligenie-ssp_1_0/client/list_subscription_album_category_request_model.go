// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubscriptionAlbumCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryName(v string) *ListSubscriptionAlbumCategoryRequest
	GetCategoryName() *string
}

type ListSubscriptionAlbumCategoryRequest struct {
	// example:
	//
	// 儿歌
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
}

func (s ListSubscriptionAlbumCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSubscriptionAlbumCategoryRequest) GoString() string {
	return s.String()
}

func (s *ListSubscriptionAlbumCategoryRequest) GetCategoryName() *string {
	return s.CategoryName
}

func (s *ListSubscriptionAlbumCategoryRequest) SetCategoryName(v string) *ListSubscriptionAlbumCategoryRequest {
	s.CategoryName = &v
	return s
}

func (s *ListSubscriptionAlbumCategoryRequest) Validate() error {
	return dara.Validate(s)
}
