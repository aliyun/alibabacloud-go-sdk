// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCubecardFiletokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCubecardFiletokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCubecardFiletokenResponse
	GetStatusCode() *int32
	SetBody(v *QueryCubecardFiletokenResponseBody) *QueryCubecardFiletokenResponse
	GetBody() *QueryCubecardFiletokenResponseBody
}

type QueryCubecardFiletokenResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCubecardFiletokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCubecardFiletokenResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCubecardFiletokenResponse) GoString() string {
	return s.String()
}

func (s *QueryCubecardFiletokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCubecardFiletokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCubecardFiletokenResponse) GetBody() *QueryCubecardFiletokenResponseBody {
	return s.Body
}

func (s *QueryCubecardFiletokenResponse) SetHeaders(v map[string]*string) *QueryCubecardFiletokenResponse {
	s.Headers = v
	return s
}

func (s *QueryCubecardFiletokenResponse) SetStatusCode(v int32) *QueryCubecardFiletokenResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCubecardFiletokenResponse) SetBody(v *QueryCubecardFiletokenResponseBody) *QueryCubecardFiletokenResponse {
	s.Body = v
	return s
}

func (s *QueryCubecardFiletokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
