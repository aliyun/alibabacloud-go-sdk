// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetShareLinkByAnonymousRequest interface {
	dara.Model
	String() string
	GoString() string
	SetShareId(v string) *GetShareLinkByAnonymousRequest
	GetShareId() *string
}

type GetShareLinkByAnonymousRequest struct {
	// The share ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7JQX1FswpQ8
	ShareId *string `json:"share_id,omitempty" xml:"share_id,omitempty"`
}

func (s GetShareLinkByAnonymousRequest) String() string {
	return dara.Prettify(s)
}

func (s GetShareLinkByAnonymousRequest) GoString() string {
	return s.String()
}

func (s *GetShareLinkByAnonymousRequest) GetShareId() *string {
	return s.ShareId
}

func (s *GetShareLinkByAnonymousRequest) SetShareId(v string) *GetShareLinkByAnonymousRequest {
	s.ShareId = &v
	return s
}

func (s *GetShareLinkByAnonymousRequest) Validate() error {
	return dara.Validate(s)
}
