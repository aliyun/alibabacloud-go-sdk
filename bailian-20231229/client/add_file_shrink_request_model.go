// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFileShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v string) *AddFileShrinkRequest
	GetCategoryId() *string
	SetCategoryType(v string) *AddFileShrinkRequest
	GetCategoryType() *string
	SetLeaseId(v string) *AddFileShrinkRequest
	GetLeaseId() *string
	SetOriginalFileUrl(v string) *AddFileShrinkRequest
	GetOriginalFileUrl() *string
	SetParser(v string) *AddFileShrinkRequest
	GetParser() *string
	SetTagsShrink(v string) *AddFileShrinkRequest
	GetTagsShrink() *string
}

type AddFileShrinkRequest struct {
	// The primary key ID of the category to which the document is uploaded. This parameter corresponds to the `CategoryId` returned by the [AddCategory](https://www.alibabacloud.com/help/eh/model-studio/developer-reference/api-bailian-2023-12-29-addcategory) operation. You can also click the ID icon next to the category name on the Unstructured Data tab of the [Application Data](https://modelstudio.console.alibabacloud.com/#/data-center) page to view the ID. You can set the parameter to default, which specifies the Default Category created by the system.
	//
	// This parameter is required.
	//
	// example:
	//
	// cate_cdd11b1b79a74e8bbd675c356a91ee3510024405
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The type of the category. Valid values:
	//
	// - UNSTRUCTURED
	//
	// - SESSION_FILE
	//
	// example:
	//
	// UNSTRUCTURED
	CategoryType *string `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	// The lease ID, which corresponds to the `FileUploadLeaseId` parameter returned by the [ApplyFileUploadLease](https://www.alibabacloud.com/help/en/model-studio/developer-reference/api-bailian-2023-12-29-applyfileuploadlease) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 68abd1dea7b6404d8f7d7b9f7fbd332d.1716698936847
	LeaseId *string `json:"LeaseId,omitempty" xml:"LeaseId,omitempty"`
	// example:
	//
	// https://thisistest.com/abc.pdf
	OriginalFileUrl *string `json:"OriginalFileUrl,omitempty" xml:"OriginalFileUrl,omitempty"`
	// The parser. Valid value:
	//
	// 	- DASHSCOPE_DOCMIND: Intelligent document parsing by Alibaba Cloud.
	//
	// This parameter is required.
	//
	// example:
	//
	// DASHSCOPE_DOCMIND
	Parser *string `json:"Parser,omitempty" xml:"Parser,omitempty"`
	// A list of tags associated with the document. The default value is null, which means no tags. You can specify up to 10 tags.
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s AddFileShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddFileShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddFileShrinkRequest) GetCategoryId() *string {
	return s.CategoryId
}

func (s *AddFileShrinkRequest) GetCategoryType() *string {
	return s.CategoryType
}

func (s *AddFileShrinkRequest) GetLeaseId() *string {
	return s.LeaseId
}

func (s *AddFileShrinkRequest) GetOriginalFileUrl() *string {
	return s.OriginalFileUrl
}

func (s *AddFileShrinkRequest) GetParser() *string {
	return s.Parser
}

func (s *AddFileShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *AddFileShrinkRequest) SetCategoryId(v string) *AddFileShrinkRequest {
	s.CategoryId = &v
	return s
}

func (s *AddFileShrinkRequest) SetCategoryType(v string) *AddFileShrinkRequest {
	s.CategoryType = &v
	return s
}

func (s *AddFileShrinkRequest) SetLeaseId(v string) *AddFileShrinkRequest {
	s.LeaseId = &v
	return s
}

func (s *AddFileShrinkRequest) SetOriginalFileUrl(v string) *AddFileShrinkRequest {
	s.OriginalFileUrl = &v
	return s
}

func (s *AddFileShrinkRequest) SetParser(v string) *AddFileShrinkRequest {
	s.Parser = &v
	return s
}

func (s *AddFileShrinkRequest) SetTagsShrink(v string) *AddFileShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *AddFileShrinkRequest) Validate() error {
	return dara.Validate(s)
}
