// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySessionByClientIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySessionByClientIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySessionByClientIdResponse
	GetStatusCode() *int32
	SetBody(v *QuerySessionByClientIdResponseBody) *QuerySessionByClientIdResponse
	GetBody() *QuerySessionByClientIdResponseBody
}

type QuerySessionByClientIdResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySessionByClientIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySessionByClientIdResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySessionByClientIdResponse) GoString() string {
	return s.String()
}

func (s *QuerySessionByClientIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySessionByClientIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySessionByClientIdResponse) GetBody() *QuerySessionByClientIdResponseBody {
	return s.Body
}

func (s *QuerySessionByClientIdResponse) SetHeaders(v map[string]*string) *QuerySessionByClientIdResponse {
	s.Headers = v
	return s
}

func (s *QuerySessionByClientIdResponse) SetStatusCode(v int32) *QuerySessionByClientIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySessionByClientIdResponse) SetBody(v *QuerySessionByClientIdResponseBody) *QuerySessionByClientIdResponse {
	s.Body = v
	return s
}

func (s *QuerySessionByClientIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
