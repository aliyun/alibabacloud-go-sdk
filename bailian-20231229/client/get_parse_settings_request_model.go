// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetParseSettingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v string) *GetParseSettingsRequest
	GetCategoryId() *string
}

type GetParseSettingsRequest struct {
	// The category ID, which is the `CategoryId` returned by the **AddCategory*	- operation. You can also obtain it from the <props="china">[Application Data](https://bailian.console.aliyun.com/?tab=app#/data-center) - File tab<props="intl">[Application Data](https://bailian.console.alibabacloud.com/?tab=app#/data-center) - Unstructured Data tab by clicking the ID icon next to the category name.
	//
	// example:
	//
	// cate_cdd11b1b79a74e8bbd675c356a91ee35xxxxxxxx
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
}

func (s GetParseSettingsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetParseSettingsRequest) GoString() string {
	return s.String()
}

func (s *GetParseSettingsRequest) GetCategoryId() *string {
	return s.CategoryId
}

func (s *GetParseSettingsRequest) SetCategoryId(v string) *GetParseSettingsRequest {
	s.CategoryId = &v
	return s
}

func (s *GetParseSettingsRequest) Validate() error {
	return dara.Validate(s)
}
