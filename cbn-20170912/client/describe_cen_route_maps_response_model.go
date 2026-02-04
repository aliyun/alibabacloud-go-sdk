// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCenRouteMapsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCenRouteMapsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCenRouteMapsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCenRouteMapsResponseBody) *DescribeCenRouteMapsResponse
	GetBody() *DescribeCenRouteMapsResponseBody
}

type DescribeCenRouteMapsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCenRouteMapsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCenRouteMapsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenRouteMapsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCenRouteMapsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCenRouteMapsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCenRouteMapsResponse) GetBody() *DescribeCenRouteMapsResponseBody {
	return s.Body
}

func (s *DescribeCenRouteMapsResponse) SetHeaders(v map[string]*string) *DescribeCenRouteMapsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCenRouteMapsResponse) SetStatusCode(v int32) *DescribeCenRouteMapsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCenRouteMapsResponse) SetBody(v *DescribeCenRouteMapsResponseBody) *DescribeCenRouteMapsResponse {
	s.Body = v
	return s
}

func (s *DescribeCenRouteMapsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
