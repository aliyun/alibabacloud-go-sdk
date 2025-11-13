// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunSearchLawQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunSearchLawQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunSearchLawQueryResponse
	GetStatusCode() *int32
	SetBody(v *RunSearchLawQueryResponseBody) *RunSearchLawQueryResponse
	GetBody() *RunSearchLawQueryResponseBody
}

type RunSearchLawQueryResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunSearchLawQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunSearchLawQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s RunSearchLawQueryResponse) GoString() string {
	return s.String()
}

func (s *RunSearchLawQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunSearchLawQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunSearchLawQueryResponse) GetBody() *RunSearchLawQueryResponseBody {
	return s.Body
}

func (s *RunSearchLawQueryResponse) SetHeaders(v map[string]*string) *RunSearchLawQueryResponse {
	s.Headers = v
	return s
}

func (s *RunSearchLawQueryResponse) SetStatusCode(v int32) *RunSearchLawQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *RunSearchLawQueryResponse) SetBody(v *RunSearchLawQueryResponseBody) *RunSearchLawQueryResponse {
	s.Body = v
	return s
}

func (s *RunSearchLawQueryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
