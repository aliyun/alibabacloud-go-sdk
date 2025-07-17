// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v string) *AddFileRequest
	GetCategoryId() *string
	SetCategoryType(v string) *AddFileRequest
	GetCategoryType() *string
	SetLeaseId(v string) *AddFileRequest
	GetLeaseId() *string
	SetOriginalFileUrl(v string) *AddFileRequest
	GetOriginalFileUrl() *string
	SetParser(v string) *AddFileRequest
	GetParser() *string
	SetTags(v []*string) *AddFileRequest
	GetTags() []*string
}

type AddFileRequest struct {
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
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s AddFileRequest) String() string {
	return dara.Prettify(s)
}

func (s AddFileRequest) GoString() string {
	return s.String()
}

func (s *AddFileRequest) GetCategoryId() *string {
	return s.CategoryId
}

func (s *AddFileRequest) GetCategoryType() *string {
	return s.CategoryType
}

func (s *AddFileRequest) GetLeaseId() *string {
	return s.LeaseId
}

func (s *AddFileRequest) GetOriginalFileUrl() *string {
	return s.OriginalFileUrl
}

func (s *AddFileRequest) GetParser() *string {
	return s.Parser
}

func (s *AddFileRequest) GetTags() []*string {
	return s.Tags
}

func (s *AddFileRequest) SetCategoryId(v string) *AddFileRequest {
	s.CategoryId = &v
	return s
}

func (s *AddFileRequest) SetCategoryType(v string) *AddFileRequest {
	s.CategoryType = &v
	return s
}

func (s *AddFileRequest) SetLeaseId(v string) *AddFileRequest {
	s.LeaseId = &v
	return s
}

func (s *AddFileRequest) SetOriginalFileUrl(v string) *AddFileRequest {
	s.OriginalFileUrl = &v
	return s
}

func (s *AddFileRequest) SetParser(v string) *AddFileRequest {
	s.Parser = &v
	return s
}

func (s *AddFileRequest) SetTags(v []*string) *AddFileRequest {
	s.Tags = v
	return s
}

func (s *AddFileRequest) Validate() error {
	return dara.Validate(s)
}
