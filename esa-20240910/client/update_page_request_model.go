// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *UpdatePageRequest
	GetContent() *string
	SetContentType(v string) *UpdatePageRequest
	GetContentType() *string
	SetDescription(v string) *UpdatePageRequest
	GetDescription() *string
	SetId(v int64) *UpdatePageRequest
	GetId() *int64
	SetName(v string) *UpdatePageRequest
	GetName() *string
	SetSiteIds(v []*int64) *UpdatePageRequest
	GetSiteIds() []*int64
}

type UpdatePageRequest struct {
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
	SiteIds []*int64 `json:"SiteIds,omitempty" xml:"SiteIds,omitempty" type:"Repeated"`
}

func (s UpdatePageRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePageRequest) GoString() string {
	return s.String()
}

func (s *UpdatePageRequest) GetContent() *string {
	return s.Content
}

func (s *UpdatePageRequest) GetContentType() *string {
	return s.ContentType
}

func (s *UpdatePageRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdatePageRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdatePageRequest) GetName() *string {
	return s.Name
}

func (s *UpdatePageRequest) GetSiteIds() []*int64 {
	return s.SiteIds
}

func (s *UpdatePageRequest) SetContent(v string) *UpdatePageRequest {
	s.Content = &v
	return s
}

func (s *UpdatePageRequest) SetContentType(v string) *UpdatePageRequest {
	s.ContentType = &v
	return s
}

func (s *UpdatePageRequest) SetDescription(v string) *UpdatePageRequest {
	s.Description = &v
	return s
}

func (s *UpdatePageRequest) SetId(v int64) *UpdatePageRequest {
	s.Id = &v
	return s
}

func (s *UpdatePageRequest) SetName(v string) *UpdatePageRequest {
	s.Name = &v
	return s
}

func (s *UpdatePageRequest) SetSiteIds(v []*int64) *UpdatePageRequest {
	s.SiteIds = v
	return s
}

func (s *UpdatePageRequest) Validate() error {
	return dara.Validate(s)
}
