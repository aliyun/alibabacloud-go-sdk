// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgGetVisitStatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DsgGetVisitStatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DsgGetVisitStatResponse
	GetStatusCode() *int32
	SetBody(v *DsgGetVisitStatResponseBody) *DsgGetVisitStatResponse
	GetBody() *DsgGetVisitStatResponseBody
}

type DsgGetVisitStatResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DsgGetVisitStatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DsgGetVisitStatResponse) String() string {
	return dara.Prettify(s)
}

func (s DsgGetVisitStatResponse) GoString() string {
	return s.String()
}

func (s *DsgGetVisitStatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DsgGetVisitStatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DsgGetVisitStatResponse) GetBody() *DsgGetVisitStatResponseBody {
	return s.Body
}

func (s *DsgGetVisitStatResponse) SetHeaders(v map[string]*string) *DsgGetVisitStatResponse {
	s.Headers = v
	return s
}

func (s *DsgGetVisitStatResponse) SetStatusCode(v int32) *DsgGetVisitStatResponse {
	s.StatusCode = &v
	return s
}

func (s *DsgGetVisitStatResponse) SetBody(v *DsgGetVisitStatResponseBody) *DsgGetVisitStatResponse {
	s.Body = v
	return s
}

func (s *DsgGetVisitStatResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
