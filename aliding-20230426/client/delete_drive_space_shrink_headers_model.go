// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDriveSpaceShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteDriveSpaceShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *DeleteDriveSpaceShrinkHeaders
	GetAccountContextShrink() *string
}

type DeleteDriveSpaceShrinkHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	AccountContextShrink *string `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s DeleteDriveSpaceShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteDriveSpaceShrinkHeaders) GoString() string {
	return s.String()
}

func (s *DeleteDriveSpaceShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteDriveSpaceShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *DeleteDriveSpaceShrinkHeaders) SetCommonHeaders(v map[string]*string) *DeleteDriveSpaceShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteDriveSpaceShrinkHeaders) SetAccountContextShrink(v string) *DeleteDriveSpaceShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *DeleteDriveSpaceShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
