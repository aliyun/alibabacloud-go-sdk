// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduledReportsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetScheduledReportsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetScheduledReportsResponse
	GetStatusCode() *int32
	SetBody(v *GetScheduledReportsResponseBody) *GetScheduledReportsResponse
	GetBody() *GetScheduledReportsResponseBody
}

type GetScheduledReportsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetScheduledReportsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetScheduledReportsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledReportsResponse) GoString() string {
	return s.String()
}

func (s *GetScheduledReportsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetScheduledReportsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetScheduledReportsResponse) GetBody() *GetScheduledReportsResponseBody {
	return s.Body
}

func (s *GetScheduledReportsResponse) SetHeaders(v map[string]*string) *GetScheduledReportsResponse {
	s.Headers = v
	return s
}

func (s *GetScheduledReportsResponse) SetStatusCode(v int32) *GetScheduledReportsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScheduledReportsResponse) SetBody(v *GetScheduledReportsResponseBody) *GetScheduledReportsResponse {
	s.Body = v
	return s
}

func (s *GetScheduledReportsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
