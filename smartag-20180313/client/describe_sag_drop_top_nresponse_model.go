// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagDropTopNResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSagDropTopNResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSagDropTopNResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSagDropTopNResponseBody) *DescribeSagDropTopNResponse
	GetBody() *DescribeSagDropTopNResponseBody
}

type DescribeSagDropTopNResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSagDropTopNResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSagDropTopNResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagDropTopNResponse) GoString() string {
	return s.String()
}

func (s *DescribeSagDropTopNResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSagDropTopNResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSagDropTopNResponse) GetBody() *DescribeSagDropTopNResponseBody {
	return s.Body
}

func (s *DescribeSagDropTopNResponse) SetHeaders(v map[string]*string) *DescribeSagDropTopNResponse {
	s.Headers = v
	return s
}

func (s *DescribeSagDropTopNResponse) SetStatusCode(v int32) *DescribeSagDropTopNResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSagDropTopNResponse) SetBody(v *DescribeSagDropTopNResponseBody) *DescribeSagDropTopNResponse {
	s.Body = v
	return s
}

func (s *DescribeSagDropTopNResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
