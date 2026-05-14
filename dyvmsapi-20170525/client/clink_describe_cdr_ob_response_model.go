// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkDescribeCdrObResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClinkDescribeCdrObResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClinkDescribeCdrObResponse
	GetStatusCode() *int32
	SetBody(v *ClinkDescribeCdrObResponseBody) *ClinkDescribeCdrObResponse
	GetBody() *ClinkDescribeCdrObResponseBody
}

type ClinkDescribeCdrObResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClinkDescribeCdrObResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClinkDescribeCdrObResponse) String() string {
	return dara.Prettify(s)
}

func (s ClinkDescribeCdrObResponse) GoString() string {
	return s.String()
}

func (s *ClinkDescribeCdrObResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClinkDescribeCdrObResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClinkDescribeCdrObResponse) GetBody() *ClinkDescribeCdrObResponseBody {
	return s.Body
}

func (s *ClinkDescribeCdrObResponse) SetHeaders(v map[string]*string) *ClinkDescribeCdrObResponse {
	s.Headers = v
	return s
}

func (s *ClinkDescribeCdrObResponse) SetStatusCode(v int32) *ClinkDescribeCdrObResponse {
	s.StatusCode = &v
	return s
}

func (s *ClinkDescribeCdrObResponse) SetBody(v *ClinkDescribeCdrObResponseBody) *ClinkDescribeCdrObResponse {
	s.Body = v
	return s
}

func (s *ClinkDescribeCdrObResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
