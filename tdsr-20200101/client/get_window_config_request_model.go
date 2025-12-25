// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWindowConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPreviewToken(v string) *GetWindowConfigRequest
	GetPreviewToken() *string
}

type GetWindowConfigRequest struct {
	// example:
	//
	// 5dc5c2dd927e45039dadb312384b****
	PreviewToken *string `json:"PreviewToken,omitempty" xml:"PreviewToken,omitempty"`
}

func (s GetWindowConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWindowConfigRequest) GoString() string {
	return s.String()
}

func (s *GetWindowConfigRequest) GetPreviewToken() *string {
	return s.PreviewToken
}

func (s *GetWindowConfigRequest) SetPreviewToken(v string) *GetWindowConfigRequest {
	s.PreviewToken = &v
	return s
}

func (s *GetWindowConfigRequest) Validate() error {
	return dara.Validate(s)
}
