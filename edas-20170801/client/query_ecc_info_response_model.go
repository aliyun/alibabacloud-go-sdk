// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEccInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryEccInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryEccInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryEccInfoResponseBody) *QueryEccInfoResponse
	GetBody() *QueryEccInfoResponseBody
}

type QueryEccInfoResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryEccInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryEccInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryEccInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryEccInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryEccInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryEccInfoResponse) GetBody() *QueryEccInfoResponseBody {
	return s.Body
}

func (s *QueryEccInfoResponse) SetHeaders(v map[string]*string) *QueryEccInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryEccInfoResponse) SetStatusCode(v int32) *QueryEccInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEccInfoResponse) SetBody(v *QueryEccInfoResponseBody) *QueryEccInfoResponse {
	s.Body = v
	return s
}

func (s *QueryEccInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
