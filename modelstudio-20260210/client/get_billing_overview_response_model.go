// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBillingOverviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBillingOverviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBillingOverviewResponse
	GetStatusCode() *int32
	SetBody(v *GetBillingOverviewResponseBody) *GetBillingOverviewResponse
	GetBody() *GetBillingOverviewResponseBody
}

type GetBillingOverviewResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBillingOverviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBillingOverviewResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBillingOverviewResponse) GoString() string {
	return s.String()
}

func (s *GetBillingOverviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBillingOverviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBillingOverviewResponse) GetBody() *GetBillingOverviewResponseBody {
	return s.Body
}

func (s *GetBillingOverviewResponse) SetHeaders(v map[string]*string) *GetBillingOverviewResponse {
	s.Headers = v
	return s
}

func (s *GetBillingOverviewResponse) SetStatusCode(v int32) *GetBillingOverviewResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBillingOverviewResponse) SetBody(v *GetBillingOverviewResponseBody) *GetBillingOverviewResponse {
	s.Body = v
	return s
}

func (s *GetBillingOverviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
