// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFolderShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddFolderShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *AddFolderShrinkHeaders
	GetAccountContextShrink() *string
}

type AddFolderShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s AddFolderShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddFolderShrinkHeaders) GoString() string {
	return s.String()
}

func (s *AddFolderShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddFolderShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *AddFolderShrinkHeaders) SetCommonHeaders(v map[string]*string) *AddFolderShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddFolderShrinkHeaders) SetAccountContextShrink(v string) *AddFolderShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *AddFolderShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
