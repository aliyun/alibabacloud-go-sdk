// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iViewInstanceRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ViewInstanceRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ViewInstanceRecordsResponse
	GetStatusCode() *int32
	SetBody(v *ViewInstanceRecordsResponseBody) *ViewInstanceRecordsResponse
	GetBody() *ViewInstanceRecordsResponseBody
}

type ViewInstanceRecordsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ViewInstanceRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ViewInstanceRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s ViewInstanceRecordsResponse) GoString() string {
	return s.String()
}

func (s *ViewInstanceRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ViewInstanceRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ViewInstanceRecordsResponse) GetBody() *ViewInstanceRecordsResponseBody {
	return s.Body
}

func (s *ViewInstanceRecordsResponse) SetHeaders(v map[string]*string) *ViewInstanceRecordsResponse {
	s.Headers = v
	return s
}

func (s *ViewInstanceRecordsResponse) SetStatusCode(v int32) *ViewInstanceRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ViewInstanceRecordsResponse) SetBody(v *ViewInstanceRecordsResponseBody) *ViewInstanceRecordsResponse {
	s.Body = v
	return s
}

func (s *ViewInstanceRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
