// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoReadSummaryOption interface {
	dara.Model
	String() string
	GoString() string
	SetSummarize(v bool) *VideoReadSummaryOption
	GetSummarize() *bool
}

type VideoReadSummaryOption struct {
	Summarize *bool `json:"Summarize,omitempty" xml:"Summarize,omitempty"`
}

func (s VideoReadSummaryOption) String() string {
	return dara.Prettify(s)
}

func (s VideoReadSummaryOption) GoString() string {
	return s.String()
}

func (s *VideoReadSummaryOption) GetSummarize() *bool {
	return s.Summarize
}

func (s *VideoReadSummaryOption) SetSummarize(v bool) *VideoReadSummaryOption {
	s.Summarize = &v
	return s
}

func (s *VideoReadSummaryOption) Validate() error {
	return dara.Validate(s)
}
