// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCdsFileShareLinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCdsId(v string) *CancelCdsFileShareLinkRequest
	GetCdsId() *string
	SetShareId(v string) *CancelCdsFileShareLinkRequest
	GetShareId() *string
}

type CancelCdsFileShareLinkRequest struct {
	// The ID of the cloud disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+cds-352282****
	CdsId *string `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
	// The ID of the file sharing task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7JQX1Fsw****
	ShareId *string `json:"ShareId,omitempty" xml:"ShareId,omitempty"`
}

func (s CancelCdsFileShareLinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelCdsFileShareLinkRequest) GoString() string {
	return s.String()
}

func (s *CancelCdsFileShareLinkRequest) GetCdsId() *string {
	return s.CdsId
}

func (s *CancelCdsFileShareLinkRequest) GetShareId() *string {
	return s.ShareId
}

func (s *CancelCdsFileShareLinkRequest) SetCdsId(v string) *CancelCdsFileShareLinkRequest {
	s.CdsId = &v
	return s
}

func (s *CancelCdsFileShareLinkRequest) SetShareId(v string) *CancelCdsFileShareLinkRequest {
	s.ShareId = &v
	return s
}

func (s *CancelCdsFileShareLinkRequest) Validate() error {
	return dara.Validate(s)
}
