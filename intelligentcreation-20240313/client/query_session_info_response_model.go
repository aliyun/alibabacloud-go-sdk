// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySessionInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySessionInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySessionInfoResponse
	GetStatusCode() *int32
	SetBody(v *QuerySessionInfoResponseBody) *QuerySessionInfoResponse
	GetBody() *QuerySessionInfoResponseBody
}

type QuerySessionInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySessionInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySessionInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySessionInfoResponse) GoString() string {
	return s.String()
}

func (s *QuerySessionInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySessionInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySessionInfoResponse) GetBody() *QuerySessionInfoResponseBody {
	return s.Body
}

func (s *QuerySessionInfoResponse) SetHeaders(v map[string]*string) *QuerySessionInfoResponse {
	s.Headers = v
	return s
}

func (s *QuerySessionInfoResponse) SetStatusCode(v int32) *QuerySessionInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySessionInfoResponse) SetBody(v *QuerySessionInfoResponseBody) *QuerySessionInfoResponse {
	s.Body = v
	return s
}

func (s *QuerySessionInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
