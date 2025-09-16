// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConfigFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteConfigFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteConfigFileResponse
	GetStatusCode() *int32
	SetBody(v *DeleteConfigFileResponseBody) *DeleteConfigFileResponse
	GetBody() *DeleteConfigFileResponseBody
}

type DeleteConfigFileResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteConfigFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteConfigFileResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteConfigFileResponse) GoString() string {
	return s.String()
}

func (s *DeleteConfigFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteConfigFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteConfigFileResponse) GetBody() *DeleteConfigFileResponseBody {
	return s.Body
}

func (s *DeleteConfigFileResponse) SetHeaders(v map[string]*string) *DeleteConfigFileResponse {
	s.Headers = v
	return s
}

func (s *DeleteConfigFileResponse) SetStatusCode(v int32) *DeleteConfigFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteConfigFileResponse) SetBody(v *DeleteConfigFileResponseBody) *DeleteConfigFileResponse {
	s.Body = v
	return s
}

func (s *DeleteConfigFileResponse) Validate() error {
	return dara.Validate(s)
}
