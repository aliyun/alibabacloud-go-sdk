// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableInsightResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableInsightResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableInsightResponse
	GetStatusCode() *int32
	SetBody(v *DisableInsightResponseBody) *DisableInsightResponse
	GetBody() *DisableInsightResponseBody
}

type DisableInsightResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableInsightResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableInsightResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableInsightResponse) GoString() string {
	return s.String()
}

func (s *DisableInsightResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableInsightResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableInsightResponse) GetBody() *DisableInsightResponseBody {
	return s.Body
}

func (s *DisableInsightResponse) SetHeaders(v map[string]*string) *DisableInsightResponse {
	s.Headers = v
	return s
}

func (s *DisableInsightResponse) SetStatusCode(v int32) *DisableInsightResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableInsightResponse) SetBody(v *DisableInsightResponseBody) *DisableInsightResponse {
	s.Body = v
	return s
}

func (s *DisableInsightResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
