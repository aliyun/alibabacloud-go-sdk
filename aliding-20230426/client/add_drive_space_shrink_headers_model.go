// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDriveSpaceShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddDriveSpaceShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *AddDriveSpaceShrinkHeaders
	GetAccountContextShrink() *string
}

type AddDriveSpaceShrinkHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	AccountContextShrink *string `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s AddDriveSpaceShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddDriveSpaceShrinkHeaders) GoString() string {
	return s.String()
}

func (s *AddDriveSpaceShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddDriveSpaceShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *AddDriveSpaceShrinkHeaders) SetCommonHeaders(v map[string]*string) *AddDriveSpaceShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddDriveSpaceShrinkHeaders) SetAccountContextShrink(v string) *AddDriveSpaceShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *AddDriveSpaceShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
