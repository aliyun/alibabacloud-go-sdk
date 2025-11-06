// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMgsApirestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMgsApirestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMgsApirestResponse
	GetStatusCode() *int32
	SetBody(v *QueryMgsApirestResponseBody) *QueryMgsApirestResponse
	GetBody() *QueryMgsApirestResponseBody
}

type QueryMgsApirestResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMgsApirestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMgsApirestResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMgsApirestResponse) GoString() string {
	return s.String()
}

func (s *QueryMgsApirestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMgsApirestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMgsApirestResponse) GetBody() *QueryMgsApirestResponseBody {
	return s.Body
}

func (s *QueryMgsApirestResponse) SetHeaders(v map[string]*string) *QueryMgsApirestResponse {
	s.Headers = v
	return s
}

func (s *QueryMgsApirestResponse) SetStatusCode(v int32) *QueryMgsApirestResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMgsApirestResponse) SetBody(v *QueryMgsApirestResponseBody) *QueryMgsApirestResponse {
	s.Body = v
	return s
}

func (s *QueryMgsApirestResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
