// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTagByParamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryTagByParamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryTagByParamResponse
	GetStatusCode() *int32
	SetBody(v *QueryTagByParamResponseBody) *QueryTagByParamResponse
	GetBody() *QueryTagByParamResponseBody
}

type QueryTagByParamResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTagByParamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTagByParamResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryTagByParamResponse) GoString() string {
	return s.String()
}

func (s *QueryTagByParamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryTagByParamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryTagByParamResponse) GetBody() *QueryTagByParamResponseBody {
	return s.Body
}

func (s *QueryTagByParamResponse) SetHeaders(v map[string]*string) *QueryTagByParamResponse {
	s.Headers = v
	return s
}

func (s *QueryTagByParamResponse) SetStatusCode(v int32) *QueryTagByParamResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTagByParamResponse) SetBody(v *QueryTagByParamResponseBody) *QueryTagByParamResponse {
	s.Body = v
	return s
}

func (s *QueryTagByParamResponse) Validate() error {
	return dara.Validate(s)
}
