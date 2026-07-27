// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmReceiptCmd interface {
	dara.Model
	String() string
	GoString() string
	SetDisputeId(v string) *ConfirmReceiptCmd
	GetDisputeId() *string
}

type ConfirmReceiptCmd struct {
	DisputeId *string `json:"disputeId,omitempty" xml:"disputeId,omitempty"`
}

func (s ConfirmReceiptCmd) String() string {
	return dara.Prettify(s)
}

func (s ConfirmReceiptCmd) GoString() string {
	return s.String()
}

func (s *ConfirmReceiptCmd) GetDisputeId() *string {
	return s.DisputeId
}

func (s *ConfirmReceiptCmd) SetDisputeId(v string) *ConfirmReceiptCmd {
	s.DisputeId = &v
	return s
}

func (s *ConfirmReceiptCmd) Validate() error {
	return dara.Validate(s)
}
