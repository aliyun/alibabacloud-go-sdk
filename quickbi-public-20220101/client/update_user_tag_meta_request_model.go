// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserTagMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTagDescription(v string) *UpdateUserTagMetaRequest
	GetTagDescription() *string
	SetTagId(v string) *UpdateUserTagMetaRequest
	GetTagId() *string
	SetTagName(v string) *UpdateUserTagMetaRequest
	GetTagName() *string
}

type UpdateUserTagMetaRequest struct {
	// The tag description.
	//
	// - Format check: Maximum length is 255 characters.
	//
	// example:
	//
	// Job Positions within the Department
	TagDescription *string `json:"TagDescription,omitempty" xml:"TagDescription,omitempty"`
	// The specified TagID.
	//
	// - Format check: Maximum length is 64 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// e82f6c6c0333431bad0225b2f85e****
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// The tag name.
	//
	// - Format check: Maximum length is 50 characters.
	//
	// - Only Chinese, English, numbers, and /\\|[]() symbols are allowed.
	//
	// This parameter is required.
	//
	// example:
	//
	// Department
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s UpdateUserTagMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserTagMetaRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserTagMetaRequest) GetTagDescription() *string {
	return s.TagDescription
}

func (s *UpdateUserTagMetaRequest) GetTagId() *string {
	return s.TagId
}

func (s *UpdateUserTagMetaRequest) GetTagName() *string {
	return s.TagName
}

func (s *UpdateUserTagMetaRequest) SetTagDescription(v string) *UpdateUserTagMetaRequest {
	s.TagDescription = &v
	return s
}

func (s *UpdateUserTagMetaRequest) SetTagId(v string) *UpdateUserTagMetaRequest {
	s.TagId = &v
	return s
}

func (s *UpdateUserTagMetaRequest) SetTagName(v string) *UpdateUserTagMetaRequest {
	s.TagName = &v
	return s
}

func (s *UpdateUserTagMetaRequest) Validate() error {
	return dara.Validate(s)
}
