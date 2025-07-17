// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRetcodeAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRetcodeAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRetcodeAppResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRetcodeAppResponseBody) *DeleteRetcodeAppResponse
	GetBody() *DeleteRetcodeAppResponseBody
}

type DeleteRetcodeAppResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRetcodeAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRetcodeAppResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRetcodeAppResponse) GoString() string {
	return s.String()
}

func (s *DeleteRetcodeAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRetcodeAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRetcodeAppResponse) GetBody() *DeleteRetcodeAppResponseBody {
	return s.Body
}

func (s *DeleteRetcodeAppResponse) SetHeaders(v map[string]*string) *DeleteRetcodeAppResponse {
	s.Headers = v
	return s
}

func (s *DeleteRetcodeAppResponse) SetStatusCode(v int32) *DeleteRetcodeAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRetcodeAppResponse) SetBody(v *DeleteRetcodeAppResponseBody) *DeleteRetcodeAppResponse {
	s.Body = v
	return s
}

func (s *DeleteRetcodeAppResponse) Validate() error {
	return dara.Validate(s)
}
