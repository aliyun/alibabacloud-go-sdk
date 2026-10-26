// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserTagMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTagDescription(v string) *AddUserTagMetaRequest
	GetTagDescription() *string
	SetTagName(v string) *AddUserTagMetaRequest
	GetTagName() *string
}

type AddUserTagMetaRequest struct {
	// Tag description. Format check: maximum length of 255 characters.
	//
	// example:
	//
	// test
	TagDescription *string `json:"TagDescription,omitempty" xml:"TagDescription,omitempty"`
	// Tag name. Format check:
	//
	// - Maximum length of 50 characters.
	//
	// - Only Chinese, English, numbers, and /\\|[]() symbols are allowed.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s AddUserTagMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s AddUserTagMetaRequest) GoString() string {
	return s.String()
}

func (s *AddUserTagMetaRequest) GetTagDescription() *string {
	return s.TagDescription
}

func (s *AddUserTagMetaRequest) GetTagName() *string {
	return s.TagName
}

func (s *AddUserTagMetaRequest) SetTagDescription(v string) *AddUserTagMetaRequest {
	s.TagDescription = &v
	return s
}

func (s *AddUserTagMetaRequest) SetTagName(v string) *AddUserTagMetaRequest {
	s.TagName = &v
	return s
}

func (s *AddUserTagMetaRequest) Validate() error {
	return dara.Validate(s)
}
