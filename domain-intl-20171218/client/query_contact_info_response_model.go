// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryContactInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryContactInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryContactInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryContactInfoResponseBody) *QueryContactInfoResponse
	GetBody() *QueryContactInfoResponseBody
}

type QueryContactInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryContactInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryContactInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryContactInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryContactInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryContactInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryContactInfoResponse) GetBody() *QueryContactInfoResponseBody {
	return s.Body
}

func (s *QueryContactInfoResponse) SetHeaders(v map[string]*string) *QueryContactInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryContactInfoResponse) SetStatusCode(v int32) *QueryContactInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryContactInfoResponse) SetBody(v *QueryContactInfoResponseBody) *QueryContactInfoResponse {
	s.Body = v
	return s
}

func (s *QueryContactInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
