// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTagStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryTagStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryTagStatusResponse
	GetStatusCode() *int32
	SetBody(v *QueryTagStatusResponseBody) *QueryTagStatusResponse
	GetBody() *QueryTagStatusResponseBody
}

type QueryTagStatusResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTagStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTagStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryTagStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryTagStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryTagStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryTagStatusResponse) GetBody() *QueryTagStatusResponseBody {
	return s.Body
}

func (s *QueryTagStatusResponse) SetHeaders(v map[string]*string) *QueryTagStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryTagStatusResponse) SetStatusCode(v int32) *QueryTagStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTagStatusResponse) SetBody(v *QueryTagStatusResponseBody) *QueryTagStatusResponse {
	s.Body = v
	return s
}

func (s *QueryTagStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
