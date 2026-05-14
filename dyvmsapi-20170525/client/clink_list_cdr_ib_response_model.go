// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkListCdrIbResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClinkListCdrIbResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClinkListCdrIbResponse
	GetStatusCode() *int32
	SetBody(v *ClinkListCdrIbResponseBody) *ClinkListCdrIbResponse
	GetBody() *ClinkListCdrIbResponseBody
}

type ClinkListCdrIbResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClinkListCdrIbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClinkListCdrIbResponse) String() string {
	return dara.Prettify(s)
}

func (s ClinkListCdrIbResponse) GoString() string {
	return s.String()
}

func (s *ClinkListCdrIbResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClinkListCdrIbResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClinkListCdrIbResponse) GetBody() *ClinkListCdrIbResponseBody {
	return s.Body
}

func (s *ClinkListCdrIbResponse) SetHeaders(v map[string]*string) *ClinkListCdrIbResponse {
	s.Headers = v
	return s
}

func (s *ClinkListCdrIbResponse) SetStatusCode(v int32) *ClinkListCdrIbResponse {
	s.StatusCode = &v
	return s
}

func (s *ClinkListCdrIbResponse) SetBody(v *ClinkListCdrIbResponseBody) *ClinkListCdrIbResponse {
	s.Body = v
	return s
}

func (s *ClinkListCdrIbResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
