// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryConfigResponse
	GetStatusCode() *int32
	SetBody(v *QueryConfigResponseBody) *QueryConfigResponse
	GetBody() *QueryConfigResponseBody
}

type QueryConfigResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryConfigResponse) GoString() string {
	return s.String()
}

func (s *QueryConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryConfigResponse) GetBody() *QueryConfigResponseBody {
	return s.Body
}

func (s *QueryConfigResponse) SetHeaders(v map[string]*string) *QueryConfigResponse {
	s.Headers = v
	return s
}

func (s *QueryConfigResponse) SetStatusCode(v int32) *QueryConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryConfigResponse) SetBody(v *QueryConfigResponseBody) *QueryConfigResponse {
	s.Body = v
	return s
}

func (s *QueryConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
