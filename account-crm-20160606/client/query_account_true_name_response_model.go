// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAccountTrueNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAccountTrueNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAccountTrueNameResponse
	GetStatusCode() *int32
	SetBody(v *QueryAccountTrueNameResponseBody) *QueryAccountTrueNameResponse
	GetBody() *QueryAccountTrueNameResponseBody
}

type QueryAccountTrueNameResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAccountTrueNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAccountTrueNameResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountTrueNameResponse) GoString() string {
	return s.String()
}

func (s *QueryAccountTrueNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAccountTrueNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAccountTrueNameResponse) GetBody() *QueryAccountTrueNameResponseBody {
	return s.Body
}

func (s *QueryAccountTrueNameResponse) SetHeaders(v map[string]*string) *QueryAccountTrueNameResponse {
	s.Headers = v
	return s
}

func (s *QueryAccountTrueNameResponse) SetStatusCode(v int32) *QueryAccountTrueNameResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAccountTrueNameResponse) SetBody(v *QueryAccountTrueNameResponseBody) *QueryAccountTrueNameResponse {
	s.Body = v
	return s
}

func (s *QueryAccountTrueNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
