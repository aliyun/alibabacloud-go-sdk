// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryNotifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryNotifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryNotifyResponse
	GetStatusCode() *int32
	SetBody(v *QueryNotifyResponseBody) *QueryNotifyResponse
	GetBody() *QueryNotifyResponseBody
}

type QueryNotifyResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryNotifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryNotifyResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryNotifyResponse) GoString() string {
	return s.String()
}

func (s *QueryNotifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryNotifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryNotifyResponse) GetBody() *QueryNotifyResponseBody {
	return s.Body
}

func (s *QueryNotifyResponse) SetHeaders(v map[string]*string) *QueryNotifyResponse {
	s.Headers = v
	return s
}

func (s *QueryNotifyResponse) SetStatusCode(v int32) *QueryNotifyResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryNotifyResponse) SetBody(v *QueryNotifyResponseBody) *QueryNotifyResponse {
	s.Body = v
	return s
}

func (s *QueryNotifyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
