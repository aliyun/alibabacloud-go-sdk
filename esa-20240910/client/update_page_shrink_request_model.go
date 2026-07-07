// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *UpdatePageShrinkRequest
	GetContent() *string
	SetContentType(v string) *UpdatePageShrinkRequest
	GetContentType() *string
	SetDescription(v string) *UpdatePageShrinkRequest
	GetDescription() *string
	SetId(v int64) *UpdatePageShrinkRequest
	GetId() *int64
	SetName(v string) *UpdatePageShrinkRequest
	GetName() *string
	SetSiteIdsShrink(v string) *UpdatePageShrinkRequest
	GetSiteIdsShrink() *string
}

type UpdatePageShrinkRequest struct {
	// The BASE64-encoded page content, which must be consistent with `ContentType`.
	//
	// **Encoding method**:
	//
	// 1. Convert the original page content to a UTF-8 byte string.
	//
	// 2. Encode the byte string using standard BASE64 encoding.
	//
	// **Example**: `<html>hello page</html>` → `PGh0bWw+aGVsbG8gcGFnZTwvaHRtbD4=`
	//
	// > The maximum size limit is subject to the server-side custom page specification. If this parameter is not specified, the original page content is retained.
	//
	// This parameter is required.
	//
	// example:
	//
	// PGh0bWw+aGVsbG8gcGFnZTwvaHRtbD4=
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The MIME type of the page content, which is returned to the client as the HTTP `Content-Type` response header when a rule is matched.
	//
	// **Common values**:
	//
	// - `text/html`: HTML page
	//
	// - `application/json`: JSON response
	//
	// > The complete set of supported values is subject to the server-side specification. The actual format of `Content` must match this field. A mismatch may cause browser rendering issues.
	//
	// This parameter is required.
	//
	// example:
	//
	// text/html
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The description of the page after the update. This is used to identify the purpose of the page in the console list. This is an optional field. If this parameter is not specified, the original description is retained. The maximum field length is subject to the server-side limit.
	//
	// This parameter is required.
	//
	// example:
	//
	// a custom deny page
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the custom response page. You can obtain this value by calling the [ListPages](https://help.aliyun.com/document_detail/2850223.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 50000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the custom response page after the update.
	//
	// **Naming suggestion**: Use a combination of letters, digits, and underscores (such as `blocked_page_v2`) for easy reference in rules. The character set, maximum length, and uniqueness constraints are subject to the server-side naming conventions for custom pages. If this parameter is not specified, the original name is retained.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of site IDs associated with this page after the update. This parameter uses full overwrite semantics.
	//
	// - You can obtain site IDs by calling the `ListSites` operation.
	//
	// - Passing an empty list dissociates all sites from the page.
	//
	// - Including a site ID that does not belong to your account returns an `InvalidParameter` error.
	SiteIdsShrink *string `json:"SiteIds,omitempty" xml:"SiteIds,omitempty"`
}

func (s UpdatePageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePageShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdatePageShrinkRequest) GetContent() *string {
	return s.Content
}

func (s *UpdatePageShrinkRequest) GetContentType() *string {
	return s.ContentType
}

func (s *UpdatePageShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdatePageShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdatePageShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdatePageShrinkRequest) GetSiteIdsShrink() *string {
	return s.SiteIdsShrink
}

func (s *UpdatePageShrinkRequest) SetContent(v string) *UpdatePageShrinkRequest {
	s.Content = &v
	return s
}

func (s *UpdatePageShrinkRequest) SetContentType(v string) *UpdatePageShrinkRequest {
	s.ContentType = &v
	return s
}

func (s *UpdatePageShrinkRequest) SetDescription(v string) *UpdatePageShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdatePageShrinkRequest) SetId(v int64) *UpdatePageShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdatePageShrinkRequest) SetName(v string) *UpdatePageShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdatePageShrinkRequest) SetSiteIdsShrink(v string) *UpdatePageShrinkRequest {
	s.SiteIdsShrink = &v
	return s
}

func (s *UpdatePageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
