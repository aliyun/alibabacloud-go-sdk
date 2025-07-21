// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTaskByParamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryTaskByParamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryTaskByParamResponse
	GetStatusCode() *int32
	SetBody(v *QueryTaskByParamResponseBody) *QueryTaskByParamResponse
	GetBody() *QueryTaskByParamResponseBody
}

type QueryTaskByParamResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTaskByParamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTaskByParamResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskByParamResponse) GoString() string {
	return s.String()
}

func (s *QueryTaskByParamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryTaskByParamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryTaskByParamResponse) GetBody() *QueryTaskByParamResponseBody {
	return s.Body
}

func (s *QueryTaskByParamResponse) SetHeaders(v map[string]*string) *QueryTaskByParamResponse {
	s.Headers = v
	return s
}

func (s *QueryTaskByParamResponse) SetStatusCode(v int32) *QueryTaskByParamResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTaskByParamResponse) SetBody(v *QueryTaskByParamResponseBody) *QueryTaskByParamResponse {
	s.Body = v
	return s
}

func (s *QueryTaskByParamResponse) Validate() error {
	return dara.Validate(s)
}
