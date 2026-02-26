// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyReason interface {
	dara.Model
	String() string
	GoString() string
	SetReasonTextId(v int64) *ApplyReason
	GetReasonTextId() *int64
	SetReasonTips(v string) *ApplyReason
	GetReasonTips() *string
}

type ApplyReason struct {
	// example:
	//
	// 403769
	ReasonTextId *int64  `json:"reasonTextId,omitempty" xml:"reasonTextId,omitempty"`
	ReasonTips   *string `json:"reasonTips,omitempty" xml:"reasonTips,omitempty"`
}

func (s ApplyReason) String() string {
	return dara.Prettify(s)
}

func (s ApplyReason) GoString() string {
	return s.String()
}

func (s *ApplyReason) GetReasonTextId() *int64 {
	return s.ReasonTextId
}

func (s *ApplyReason) GetReasonTips() *string {
	return s.ReasonTips
}

func (s *ApplyReason) SetReasonTextId(v int64) *ApplyReason {
	s.ReasonTextId = &v
	return s
}

func (s *ApplyReason) SetReasonTips(v string) *ApplyReason {
	s.ReasonTips = &v
	return s
}

func (s *ApplyReason) Validate() error {
	return dara.Validate(s)
}
