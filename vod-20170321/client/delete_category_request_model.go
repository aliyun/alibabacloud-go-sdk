// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCateId(v int64) *DeleteCategoryRequest
	GetCateId() *int64
}

type DeleteCategoryRequest struct {
	// The ID of the category. You can specify only one ID. You can use one of the following methods to obtain the ID:
	//
	// 	- Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com). Choose **Configuration Management*	- > **Media Management*	- > **Categories**. On the Audio and Video / Image Category or Short Video Material Category tab, view the category ID.
	//
	// 	- Obtain the category ID from the response to the [AddCategory](~~AddCategory~~) operation.
	//
	// >  If you specify the ID of a parent category, all subcategories under the parent category are deleted at the same time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3300****
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
}

func (s DeleteCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCategoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteCategoryRequest) GetCateId() *int64 {
	return s.CateId
}

func (s *DeleteCategoryRequest) SetCateId(v int64) *DeleteCategoryRequest {
	s.CateId = &v
	return s
}

func (s *DeleteCategoryRequest) Validate() error {
	return dara.Validate(s)
}
