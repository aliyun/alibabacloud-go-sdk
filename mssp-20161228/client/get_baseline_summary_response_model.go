// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBaselineSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBaselineSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBaselineSummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetBaselineSummaryResponseBody) *GetBaselineSummaryResponse
	GetBody() *GetBaselineSummaryResponseBody
}

type GetBaselineSummaryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBaselineSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBaselineSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBaselineSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetBaselineSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBaselineSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBaselineSummaryResponse) GetBody() *GetBaselineSummaryResponseBody {
	return s.Body
}

func (s *GetBaselineSummaryResponse) SetHeaders(v map[string]*string) *GetBaselineSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetBaselineSummaryResponse) SetStatusCode(v int32) *GetBaselineSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBaselineSummaryResponse) SetBody(v *GetBaselineSummaryResponseBody) *GetBaselineSummaryResponse {
	s.Body = v
	return s
}

func (s *GetBaselineSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
