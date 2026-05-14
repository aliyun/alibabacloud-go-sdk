// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkListCdrObResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClinkListCdrObResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClinkListCdrObResponse
	GetStatusCode() *int32
	SetBody(v *ClinkListCdrObResponseBody) *ClinkListCdrObResponse
	GetBody() *ClinkListCdrObResponseBody
}

type ClinkListCdrObResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClinkListCdrObResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClinkListCdrObResponse) String() string {
	return dara.Prettify(s)
}

func (s ClinkListCdrObResponse) GoString() string {
	return s.String()
}

func (s *ClinkListCdrObResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClinkListCdrObResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClinkListCdrObResponse) GetBody() *ClinkListCdrObResponseBody {
	return s.Body
}

func (s *ClinkListCdrObResponse) SetHeaders(v map[string]*string) *ClinkListCdrObResponse {
	s.Headers = v
	return s
}

func (s *ClinkListCdrObResponse) SetStatusCode(v int32) *ClinkListCdrObResponse {
	s.StatusCode = &v
	return s
}

func (s *ClinkListCdrObResponse) SetBody(v *ClinkListCdrObResponseBody) *ClinkListCdrObResponse {
	s.Body = v
	return s
}

func (s *ClinkListCdrObResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
