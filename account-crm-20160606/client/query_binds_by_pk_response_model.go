// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryBindsByPkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryBindsByPkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryBindsByPkResponse
	GetStatusCode() *int32
	SetBody(v *QueryBindsByPkResponseBody) *QueryBindsByPkResponse
	GetBody() *QueryBindsByPkResponseBody
}

type QueryBindsByPkResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryBindsByPkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryBindsByPkResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryBindsByPkResponse) GoString() string {
	return s.String()
}

func (s *QueryBindsByPkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryBindsByPkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryBindsByPkResponse) GetBody() *QueryBindsByPkResponseBody {
	return s.Body
}

func (s *QueryBindsByPkResponse) SetHeaders(v map[string]*string) *QueryBindsByPkResponse {
	s.Headers = v
	return s
}

func (s *QueryBindsByPkResponse) SetStatusCode(v int32) *QueryBindsByPkResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBindsByPkResponse) SetBody(v *QueryBindsByPkResponseBody) *QueryBindsByPkResponse {
	s.Body = v
	return s
}

func (s *QueryBindsByPkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
