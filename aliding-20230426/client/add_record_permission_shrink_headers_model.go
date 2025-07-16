// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRecordPermissionShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddRecordPermissionShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *AddRecordPermissionShrinkHeaders
	GetAccountContextShrink() *string
}

type AddRecordPermissionShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s AddRecordPermissionShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddRecordPermissionShrinkHeaders) GoString() string {
	return s.String()
}

func (s *AddRecordPermissionShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddRecordPermissionShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *AddRecordPermissionShrinkHeaders) SetCommonHeaders(v map[string]*string) *AddRecordPermissionShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddRecordPermissionShrinkHeaders) SetAccountContextShrink(v string) *AddRecordPermissionShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *AddRecordPermissionShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
