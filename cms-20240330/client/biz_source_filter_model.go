// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBizSourceFilter interface {
	dara.Model
	String() string
	GoString() string
	SetEq(v string) *BizSourceFilter
	GetEq() *string
}

type BizSourceFilter struct {
	Eq *string `json:"eq,omitempty" xml:"eq,omitempty"`
}

func (s BizSourceFilter) String() string {
	return dara.Prettify(s)
}

func (s BizSourceFilter) GoString() string {
	return s.String()
}

func (s *BizSourceFilter) GetEq() *string {
	return s.Eq
}

func (s *BizSourceFilter) SetEq(v string) *BizSourceFilter {
	s.Eq = &v
	return s
}

func (s *BizSourceFilter) Validate() error {
	return dara.Validate(s)
}
