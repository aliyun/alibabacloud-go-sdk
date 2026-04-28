// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelShareLinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetShareId(v string) *CancelShareLinkRequest
	GetShareId() *string
}

type CancelShareLinkRequest struct {
	// The share ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7JQX1FswpQ8
	ShareId *string `json:"share_id,omitempty" xml:"share_id,omitempty"`
}

func (s CancelShareLinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelShareLinkRequest) GoString() string {
	return s.String()
}

func (s *CancelShareLinkRequest) GetShareId() *string {
	return s.ShareId
}

func (s *CancelShareLinkRequest) SetShareId(v string) *CancelShareLinkRequest {
	s.ShareId = &v
	return s
}

func (s *CancelShareLinkRequest) Validate() error {
	return dara.Validate(s)
}
