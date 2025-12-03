// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryHBaseHaDBResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryHBaseHaDBResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryHBaseHaDBResponse
	GetStatusCode() *int32
	SetBody(v *QueryHBaseHaDBResponseBody) *QueryHBaseHaDBResponse
	GetBody() *QueryHBaseHaDBResponseBody
}

type QueryHBaseHaDBResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryHBaseHaDBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryHBaseHaDBResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryHBaseHaDBResponse) GoString() string {
	return s.String()
}

func (s *QueryHBaseHaDBResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryHBaseHaDBResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryHBaseHaDBResponse) GetBody() *QueryHBaseHaDBResponseBody {
	return s.Body
}

func (s *QueryHBaseHaDBResponse) SetHeaders(v map[string]*string) *QueryHBaseHaDBResponse {
	s.Headers = v
	return s
}

func (s *QueryHBaseHaDBResponse) SetStatusCode(v int32) *QueryHBaseHaDBResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryHBaseHaDBResponse) SetBody(v *QueryHBaseHaDBResponseBody) *QueryHBaseHaDBResponse {
	s.Body = v
	return s
}

func (s *QueryHBaseHaDBResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
