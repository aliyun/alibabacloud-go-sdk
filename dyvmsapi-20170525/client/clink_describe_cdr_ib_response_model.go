// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkDescribeCdrIbResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClinkDescribeCdrIbResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClinkDescribeCdrIbResponse
	GetStatusCode() *int32
	SetBody(v *ClinkDescribeCdrIbResponseBody) *ClinkDescribeCdrIbResponse
	GetBody() *ClinkDescribeCdrIbResponseBody
}

type ClinkDescribeCdrIbResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClinkDescribeCdrIbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClinkDescribeCdrIbResponse) String() string {
	return dara.Prettify(s)
}

func (s ClinkDescribeCdrIbResponse) GoString() string {
	return s.String()
}

func (s *ClinkDescribeCdrIbResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClinkDescribeCdrIbResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClinkDescribeCdrIbResponse) GetBody() *ClinkDescribeCdrIbResponseBody {
	return s.Body
}

func (s *ClinkDescribeCdrIbResponse) SetHeaders(v map[string]*string) *ClinkDescribeCdrIbResponse {
	s.Headers = v
	return s
}

func (s *ClinkDescribeCdrIbResponse) SetStatusCode(v int32) *ClinkDescribeCdrIbResponse {
	s.StatusCode = &v
	return s
}

func (s *ClinkDescribeCdrIbResponse) SetBody(v *ClinkDescribeCdrIbResponseBody) *ClinkDescribeCdrIbResponse {
	s.Body = v
	return s
}

func (s *ClinkDescribeCdrIbResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
