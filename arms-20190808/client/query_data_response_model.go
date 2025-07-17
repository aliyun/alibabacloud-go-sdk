// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetResults(v string) *QueryDataResponse
	GetResults() *string
}

type QueryDataResponse struct {
	Results *string `json:"results,omitempty" xml:"results,omitempty"`
}

func (s QueryDataResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDataResponse) GoString() string {
	return s.String()
}

func (s *QueryDataResponse) GetResults() *string {
	return s.Results
}

func (s *QueryDataResponse) SetResults(v string) *QueryDataResponse {
	s.Results = &v
	return s
}

func (s *QueryDataResponse) Validate() error {
	return dara.Validate(s)
}
