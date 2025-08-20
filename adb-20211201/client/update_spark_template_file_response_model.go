// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSparkTemplateFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSparkTemplateFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSparkTemplateFileResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSparkTemplateFileResponseBody) *UpdateSparkTemplateFileResponse
	GetBody() *UpdateSparkTemplateFileResponseBody
}

type UpdateSparkTemplateFileResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSparkTemplateFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSparkTemplateFileResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSparkTemplateFileResponse) GoString() string {
	return s.String()
}

func (s *UpdateSparkTemplateFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSparkTemplateFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSparkTemplateFileResponse) GetBody() *UpdateSparkTemplateFileResponseBody {
	return s.Body
}

func (s *UpdateSparkTemplateFileResponse) SetHeaders(v map[string]*string) *UpdateSparkTemplateFileResponse {
	s.Headers = v
	return s
}

func (s *UpdateSparkTemplateFileResponse) SetStatusCode(v int32) *UpdateSparkTemplateFileResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSparkTemplateFileResponse) SetBody(v *UpdateSparkTemplateFileResponseBody) *UpdateSparkTemplateFileResponse {
	s.Body = v
	return s
}

func (s *UpdateSparkTemplateFileResponse) Validate() error {
	return dara.Validate(s)
}
