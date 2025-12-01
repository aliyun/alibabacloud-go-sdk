// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSuspPageSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSuspPageSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSuspPageSummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetSuspPageSummaryResponseBody) *GetSuspPageSummaryResponse
	GetBody() *GetSuspPageSummaryResponseBody
}

type GetSuspPageSummaryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSuspPageSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSuspPageSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSuspPageSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetSuspPageSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSuspPageSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSuspPageSummaryResponse) GetBody() *GetSuspPageSummaryResponseBody {
	return s.Body
}

func (s *GetSuspPageSummaryResponse) SetHeaders(v map[string]*string) *GetSuspPageSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetSuspPageSummaryResponse) SetStatusCode(v int32) *GetSuspPageSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSuspPageSummaryResponse) SetBody(v *GetSuspPageSummaryResponseBody) *GetSuspPageSummaryResponse {
	s.Body = v
	return s
}

func (s *GetSuspPageSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
