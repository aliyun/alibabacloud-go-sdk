// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTemplateByParamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryTemplateByParamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryTemplateByParamResponse
	GetStatusCode() *int32
	SetBody(v *QueryTemplateByParamResponseBody) *QueryTemplateByParamResponse
	GetBody() *QueryTemplateByParamResponseBody
}

type QueryTemplateByParamResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTemplateByParamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTemplateByParamResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryTemplateByParamResponse) GoString() string {
	return s.String()
}

func (s *QueryTemplateByParamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryTemplateByParamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryTemplateByParamResponse) GetBody() *QueryTemplateByParamResponseBody {
	return s.Body
}

func (s *QueryTemplateByParamResponse) SetHeaders(v map[string]*string) *QueryTemplateByParamResponse {
	s.Headers = v
	return s
}

func (s *QueryTemplateByParamResponse) SetStatusCode(v int32) *QueryTemplateByParamResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTemplateByParamResponse) SetBody(v *QueryTemplateByParamResponseBody) *QueryTemplateByParamResponse {
	s.Body = v
	return s
}

func (s *QueryTemplateByParamResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
