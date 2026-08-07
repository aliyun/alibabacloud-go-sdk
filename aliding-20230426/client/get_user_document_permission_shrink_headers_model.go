// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserDocumentPermissionShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetUserDocumentPermissionShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetUserDocumentPermissionShrinkHeaders
	GetAccountContextShrink() *string
}

type GetUserDocumentPermissionShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetUserDocumentPermissionShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetUserDocumentPermissionShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetUserDocumentPermissionShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetUserDocumentPermissionShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetUserDocumentPermissionShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetUserDocumentPermissionShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserDocumentPermissionShrinkHeaders) SetAccountContextShrink(v string) *GetUserDocumentPermissionShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetUserDocumentPermissionShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
