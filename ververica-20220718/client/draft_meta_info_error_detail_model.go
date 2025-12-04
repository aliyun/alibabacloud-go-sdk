// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDraftMetaInfoErrorDetail interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DraftMetaInfoErrorDetail
	GetMessage() *string
	SetReason(v string) *DraftMetaInfoErrorDetail
	GetReason() *string
}

type DraftMetaInfoErrorDetail struct {
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	Reason  *string `json:"reason,omitempty" xml:"reason,omitempty"`
}

func (s DraftMetaInfoErrorDetail) String() string {
	return dara.Prettify(s)
}

func (s DraftMetaInfoErrorDetail) GoString() string {
	return s.String()
}

func (s *DraftMetaInfoErrorDetail) GetMessage() *string {
	return s.Message
}

func (s *DraftMetaInfoErrorDetail) GetReason() *string {
	return s.Reason
}

func (s *DraftMetaInfoErrorDetail) SetMessage(v string) *DraftMetaInfoErrorDetail {
	s.Message = &v
	return s
}

func (s *DraftMetaInfoErrorDetail) SetReason(v string) *DraftMetaInfoErrorDetail {
	s.Reason = &v
	return s
}

func (s *DraftMetaInfoErrorDetail) Validate() error {
	return dara.Validate(s)
}
