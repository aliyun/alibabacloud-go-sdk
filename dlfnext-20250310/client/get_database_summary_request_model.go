// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatabaseSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDate(v string) *GetDatabaseSummaryRequest
	GetDate() *string
}

type GetDatabaseSummaryRequest struct {
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
}

func (s GetDatabaseSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDatabaseSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetDatabaseSummaryRequest) GetDate() *string {
	return s.Date
}

func (s *GetDatabaseSummaryRequest) SetDate(v string) *GetDatabaseSummaryRequest {
	s.Date = &v
	return s
}

func (s *GetDatabaseSummaryRequest) Validate() error {
	return dara.Validate(s)
}
