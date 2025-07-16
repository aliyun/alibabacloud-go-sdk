// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocUpdateContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DocUpdateContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DocUpdateContentResponse
	GetStatusCode() *int32
	SetBody(v *DocUpdateContentResponseBody) *DocUpdateContentResponse
	GetBody() *DocUpdateContentResponseBody
}

type DocUpdateContentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DocUpdateContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DocUpdateContentResponse) String() string {
	return dara.Prettify(s)
}

func (s DocUpdateContentResponse) GoString() string {
	return s.String()
}

func (s *DocUpdateContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DocUpdateContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DocUpdateContentResponse) GetBody() *DocUpdateContentResponseBody {
	return s.Body
}

func (s *DocUpdateContentResponse) SetHeaders(v map[string]*string) *DocUpdateContentResponse {
	s.Headers = v
	return s
}

func (s *DocUpdateContentResponse) SetStatusCode(v int32) *DocUpdateContentResponse {
	s.StatusCode = &v
	return s
}

func (s *DocUpdateContentResponse) SetBody(v *DocUpdateContentResponseBody) *DocUpdateContentResponse {
	s.Body = v
	return s
}

func (s *DocUpdateContentResponse) Validate() error {
	return dara.Validate(s)
}
