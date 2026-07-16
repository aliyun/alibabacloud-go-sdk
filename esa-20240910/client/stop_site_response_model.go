// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopSiteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopSiteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopSiteResponse
	GetStatusCode() *int32
	SetBody(v *StopSiteResponseBody) *StopSiteResponse
	GetBody() *StopSiteResponseBody
}

type StopSiteResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopSiteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopSiteResponse) String() string {
	return dara.Prettify(s)
}

func (s StopSiteResponse) GoString() string {
	return s.String()
}

func (s *StopSiteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopSiteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopSiteResponse) GetBody() *StopSiteResponseBody {
	return s.Body
}

func (s *StopSiteResponse) SetHeaders(v map[string]*string) *StopSiteResponse {
	s.Headers = v
	return s
}

func (s *StopSiteResponse) SetStatusCode(v int32) *StopSiteResponse {
	s.StatusCode = &v
	return s
}

func (s *StopSiteResponse) SetBody(v *StopSiteResponseBody) *StopSiteResponse {
	s.Body = v
	return s
}

func (s *StopSiteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
