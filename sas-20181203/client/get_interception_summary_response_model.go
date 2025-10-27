// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInterceptionSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInterceptionSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInterceptionSummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetInterceptionSummaryResponseBody) *GetInterceptionSummaryResponse
	GetBody() *GetInterceptionSummaryResponseBody
}

type GetInterceptionSummaryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInterceptionSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInterceptionSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInterceptionSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetInterceptionSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInterceptionSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInterceptionSummaryResponse) GetBody() *GetInterceptionSummaryResponseBody {
	return s.Body
}

func (s *GetInterceptionSummaryResponse) SetHeaders(v map[string]*string) *GetInterceptionSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetInterceptionSummaryResponse) SetStatusCode(v int32) *GetInterceptionSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInterceptionSummaryResponse) SetBody(v *GetInterceptionSummaryResponseBody) *GetInterceptionSummaryResponse {
	s.Body = v
	return s
}

func (s *GetInterceptionSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
