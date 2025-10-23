// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEmissionSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEmissionSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEmissionSummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetEmissionSummaryResponseBody) *GetEmissionSummaryResponse
	GetBody() *GetEmissionSummaryResponseBody
}

type GetEmissionSummaryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEmissionSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEmissionSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEmissionSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetEmissionSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEmissionSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEmissionSummaryResponse) GetBody() *GetEmissionSummaryResponseBody {
	return s.Body
}

func (s *GetEmissionSummaryResponse) SetHeaders(v map[string]*string) *GetEmissionSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetEmissionSummaryResponse) SetStatusCode(v int32) *GetEmissionSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEmissionSummaryResponse) SetBody(v *GetEmissionSummaryResponseBody) *GetEmissionSummaryResponse {
	s.Body = v
	return s
}

func (s *GetEmissionSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
