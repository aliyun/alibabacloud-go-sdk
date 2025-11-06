// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMgsTestreqbodyautogenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMgsTestreqbodyautogenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMgsTestreqbodyautogenResponse
	GetStatusCode() *int32
	SetBody(v *QueryMgsTestreqbodyautogenResponseBody) *QueryMgsTestreqbodyautogenResponse
	GetBody() *QueryMgsTestreqbodyautogenResponseBody
}

type QueryMgsTestreqbodyautogenResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMgsTestreqbodyautogenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMgsTestreqbodyautogenResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMgsTestreqbodyautogenResponse) GoString() string {
	return s.String()
}

func (s *QueryMgsTestreqbodyautogenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMgsTestreqbodyautogenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMgsTestreqbodyautogenResponse) GetBody() *QueryMgsTestreqbodyautogenResponseBody {
	return s.Body
}

func (s *QueryMgsTestreqbodyautogenResponse) SetHeaders(v map[string]*string) *QueryMgsTestreqbodyautogenResponse {
	s.Headers = v
	return s
}

func (s *QueryMgsTestreqbodyautogenResponse) SetStatusCode(v int32) *QueryMgsTestreqbodyautogenResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMgsTestreqbodyautogenResponse) SetBody(v *QueryMgsTestreqbodyautogenResponseBody) *QueryMgsTestreqbodyautogenResponse {
	s.Body = v
	return s
}

func (s *QueryMgsTestreqbodyautogenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
