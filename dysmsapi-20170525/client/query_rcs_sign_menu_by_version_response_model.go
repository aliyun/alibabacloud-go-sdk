// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRcsSignMenuByVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRcsSignMenuByVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRcsSignMenuByVersionResponse
	GetStatusCode() *int32
	SetBody(v *QueryRcsSignMenuByVersionResponseBody) *QueryRcsSignMenuByVersionResponse
	GetBody() *QueryRcsSignMenuByVersionResponseBody
}

type QueryRcsSignMenuByVersionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRcsSignMenuByVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRcsSignMenuByVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRcsSignMenuByVersionResponse) GoString() string {
	return s.String()
}

func (s *QueryRcsSignMenuByVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRcsSignMenuByVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRcsSignMenuByVersionResponse) GetBody() *QueryRcsSignMenuByVersionResponseBody {
	return s.Body
}

func (s *QueryRcsSignMenuByVersionResponse) SetHeaders(v map[string]*string) *QueryRcsSignMenuByVersionResponse {
	s.Headers = v
	return s
}

func (s *QueryRcsSignMenuByVersionResponse) SetStatusCode(v int32) *QueryRcsSignMenuByVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRcsSignMenuByVersionResponse) SetBody(v *QueryRcsSignMenuByVersionResponseBody) *QueryRcsSignMenuByVersionResponse {
	s.Body = v
	return s
}

func (s *QueryRcsSignMenuByVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
