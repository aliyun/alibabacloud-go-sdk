// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDate(v string) *GetTableSummaryRequest
	GetDate() *string
}

type GetTableSummaryRequest struct {
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
}

func (s GetTableSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTableSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetTableSummaryRequest) GetDate() *string {
	return s.Date
}

func (s *GetTableSummaryRequest) SetDate(v string) *GetTableSummaryRequest {
	s.Date = &v
	return s
}

func (s *GetTableSummaryRequest) Validate() error {
	return dara.Validate(s)
}
