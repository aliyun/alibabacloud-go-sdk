// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEpdSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEpdSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEpdSummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetEpdSummaryResponseBody) *GetEpdSummaryResponse
	GetBody() *GetEpdSummaryResponseBody
}

type GetEpdSummaryResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEpdSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEpdSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEpdSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetEpdSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEpdSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEpdSummaryResponse) GetBody() *GetEpdSummaryResponseBody {
	return s.Body
}

func (s *GetEpdSummaryResponse) SetHeaders(v map[string]*string) *GetEpdSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetEpdSummaryResponse) SetStatusCode(v int32) *GetEpdSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEpdSummaryResponse) SetBody(v *GetEpdSummaryResponseBody) *GetEpdSummaryResponse {
	s.Body = v
	return s
}

func (s *GetEpdSummaryResponse) Validate() error {
	return dara.Validate(s)
}
