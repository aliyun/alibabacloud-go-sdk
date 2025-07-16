// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDriveSpacesShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListDriveSpacesShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *ListDriveSpacesShrinkHeaders
	GetAccountContextShrink() *string
}

type ListDriveSpacesShrinkHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	AccountContextShrink *string `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s ListDriveSpacesShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListDriveSpacesShrinkHeaders) GoString() string {
	return s.String()
}

func (s *ListDriveSpacesShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListDriveSpacesShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *ListDriveSpacesShrinkHeaders) SetCommonHeaders(v map[string]*string) *ListDriveSpacesShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListDriveSpacesShrinkHeaders) SetAccountContextShrink(v string) *ListDriveSpacesShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *ListDriveSpacesShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
