// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkListCdrIbsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClinkListCdrIbsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClinkListCdrIbsResponse
	GetStatusCode() *int32
	SetBody(v *ClinkListCdrIbsResponseBody) *ClinkListCdrIbsResponse
	GetBody() *ClinkListCdrIbsResponseBody
}

type ClinkListCdrIbsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClinkListCdrIbsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClinkListCdrIbsResponse) String() string {
	return dara.Prettify(s)
}

func (s ClinkListCdrIbsResponse) GoString() string {
	return s.String()
}

func (s *ClinkListCdrIbsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClinkListCdrIbsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClinkListCdrIbsResponse) GetBody() *ClinkListCdrIbsResponseBody {
	return s.Body
}

func (s *ClinkListCdrIbsResponse) SetHeaders(v map[string]*string) *ClinkListCdrIbsResponse {
	s.Headers = v
	return s
}

func (s *ClinkListCdrIbsResponse) SetStatusCode(v int32) *ClinkListCdrIbsResponse {
	s.StatusCode = &v
	return s
}

func (s *ClinkListCdrIbsResponse) SetBody(v *ClinkListCdrIbsResponseBody) *ClinkListCdrIbsResponse {
	s.Body = v
	return s
}

func (s *ClinkListCdrIbsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
