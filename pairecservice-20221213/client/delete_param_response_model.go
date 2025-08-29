// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteParamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteParamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteParamResponse
	GetStatusCode() *int32
	SetBody(v *DeleteParamResponseBody) *DeleteParamResponse
	GetBody() *DeleteParamResponseBody
}

type DeleteParamResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteParamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteParamResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteParamResponse) GoString() string {
	return s.String()
}

func (s *DeleteParamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteParamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteParamResponse) GetBody() *DeleteParamResponseBody {
	return s.Body
}

func (s *DeleteParamResponse) SetHeaders(v map[string]*string) *DeleteParamResponse {
	s.Headers = v
	return s
}

func (s *DeleteParamResponse) SetStatusCode(v int32) *DeleteParamResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteParamResponse) SetBody(v *DeleteParamResponseBody) *DeleteParamResponse {
	s.Body = v
	return s
}

func (s *DeleteParamResponse) Validate() error {
	return dara.Validate(s)
}
