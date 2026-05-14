// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkStatIbResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClinkStatIbResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClinkStatIbResponse
	GetStatusCode() *int32
	SetBody(v *ClinkStatIbResponseBody) *ClinkStatIbResponse
	GetBody() *ClinkStatIbResponseBody
}

type ClinkStatIbResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClinkStatIbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClinkStatIbResponse) String() string {
	return dara.Prettify(s)
}

func (s ClinkStatIbResponse) GoString() string {
	return s.String()
}

func (s *ClinkStatIbResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClinkStatIbResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClinkStatIbResponse) GetBody() *ClinkStatIbResponseBody {
	return s.Body
}

func (s *ClinkStatIbResponse) SetHeaders(v map[string]*string) *ClinkStatIbResponse {
	s.Headers = v
	return s
}

func (s *ClinkStatIbResponse) SetStatusCode(v int32) *ClinkStatIbResponse {
	s.StatusCode = &v
	return s
}

func (s *ClinkStatIbResponse) SetBody(v *ClinkStatIbResponseBody) *ClinkStatIbResponse {
	s.Body = v
	return s
}

func (s *ClinkStatIbResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
