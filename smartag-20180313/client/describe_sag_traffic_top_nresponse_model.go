// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagTrafficTopNResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSagTrafficTopNResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSagTrafficTopNResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSagTrafficTopNResponseBody) *DescribeSagTrafficTopNResponse
	GetBody() *DescribeSagTrafficTopNResponseBody
}

type DescribeSagTrafficTopNResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSagTrafficTopNResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSagTrafficTopNResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagTrafficTopNResponse) GoString() string {
	return s.String()
}

func (s *DescribeSagTrafficTopNResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSagTrafficTopNResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSagTrafficTopNResponse) GetBody() *DescribeSagTrafficTopNResponseBody {
	return s.Body
}

func (s *DescribeSagTrafficTopNResponse) SetHeaders(v map[string]*string) *DescribeSagTrafficTopNResponse {
	s.Headers = v
	return s
}

func (s *DescribeSagTrafficTopNResponse) SetStatusCode(v int32) *DescribeSagTrafficTopNResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSagTrafficTopNResponse) SetBody(v *DescribeSagTrafficTopNResponseBody) *DescribeSagTrafficTopNResponse {
	s.Body = v
	return s
}

func (s *DescribeSagTrafficTopNResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
