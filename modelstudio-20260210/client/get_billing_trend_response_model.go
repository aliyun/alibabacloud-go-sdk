// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBillingTrendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBillingTrendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBillingTrendResponse
	GetStatusCode() *int32
	SetBody(v *GetBillingTrendResponseBody) *GetBillingTrendResponse
	GetBody() *GetBillingTrendResponseBody
}

type GetBillingTrendResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBillingTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBillingTrendResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBillingTrendResponse) GoString() string {
	return s.String()
}

func (s *GetBillingTrendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBillingTrendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBillingTrendResponse) GetBody() *GetBillingTrendResponseBody {
	return s.Body
}

func (s *GetBillingTrendResponse) SetHeaders(v map[string]*string) *GetBillingTrendResponse {
	s.Headers = v
	return s
}

func (s *GetBillingTrendResponse) SetStatusCode(v int32) *GetBillingTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBillingTrendResponse) SetBody(v *GetBillingTrendResponseBody) *GetBillingTrendResponse {
	s.Body = v
	return s
}

func (s *GetBillingTrendResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
