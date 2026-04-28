// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetShareLinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetShareId(v string) *GetShareLinkRequest
	GetShareId() *string
}

type GetShareLinkRequest struct {
	// The share ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7JQX1FswpQ8
	ShareId *string `json:"share_id,omitempty" xml:"share_id,omitempty"`
}

func (s GetShareLinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetShareLinkRequest) GoString() string {
	return s.String()
}

func (s *GetShareLinkRequest) GetShareId() *string {
	return s.ShareId
}

func (s *GetShareLinkRequest) SetShareId(v string) *GetShareLinkRequest {
	s.ShareId = &v
	return s
}

func (s *GetShareLinkRequest) Validate() error {
	return dara.Validate(s)
}
