// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkListCdrObsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClinkListCdrObsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClinkListCdrObsResponse
	GetStatusCode() *int32
	SetBody(v *ClinkListCdrObsResponseBody) *ClinkListCdrObsResponse
	GetBody() *ClinkListCdrObsResponseBody
}

type ClinkListCdrObsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClinkListCdrObsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClinkListCdrObsResponse) String() string {
	return dara.Prettify(s)
}

func (s ClinkListCdrObsResponse) GoString() string {
	return s.String()
}

func (s *ClinkListCdrObsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClinkListCdrObsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClinkListCdrObsResponse) GetBody() *ClinkListCdrObsResponseBody {
	return s.Body
}

func (s *ClinkListCdrObsResponse) SetHeaders(v map[string]*string) *ClinkListCdrObsResponse {
	s.Headers = v
	return s
}

func (s *ClinkListCdrObsResponse) SetStatusCode(v int32) *ClinkListCdrObsResponse {
	s.StatusCode = &v
	return s
}

func (s *ClinkListCdrObsResponse) SetBody(v *ClinkListCdrObsResponseBody) *ClinkListCdrObsResponse {
	s.Body = v
	return s
}

func (s *ClinkListCdrObsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
