// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceSummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceSummaryResponseBody) *GetInstanceSummaryResponse
	GetBody() *GetInstanceSummaryResponseBody
}

type GetInstanceSummaryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceSummaryResponse) GetBody() *GetInstanceSummaryResponseBody {
	return s.Body
}

func (s *GetInstanceSummaryResponse) SetHeaders(v map[string]*string) *GetInstanceSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceSummaryResponse) SetStatusCode(v int32) *GetInstanceSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceSummaryResponse) SetBody(v *GetInstanceSummaryResponseBody) *GetInstanceSummaryResponse {
	s.Body = v
	return s
}

func (s *GetInstanceSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
