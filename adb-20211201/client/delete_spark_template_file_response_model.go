// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSparkTemplateFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSparkTemplateFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSparkTemplateFileResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSparkTemplateFileResponseBody) *DeleteSparkTemplateFileResponse
	GetBody() *DeleteSparkTemplateFileResponseBody
}

type DeleteSparkTemplateFileResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSparkTemplateFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSparkTemplateFileResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSparkTemplateFileResponse) GoString() string {
	return s.String()
}

func (s *DeleteSparkTemplateFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSparkTemplateFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSparkTemplateFileResponse) GetBody() *DeleteSparkTemplateFileResponseBody {
	return s.Body
}

func (s *DeleteSparkTemplateFileResponse) SetHeaders(v map[string]*string) *DeleteSparkTemplateFileResponse {
	s.Headers = v
	return s
}

func (s *DeleteSparkTemplateFileResponse) SetStatusCode(v int32) *DeleteSparkTemplateFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSparkTemplateFileResponse) SetBody(v *DeleteSparkTemplateFileResponseBody) *DeleteSparkTemplateFileResponse {
	s.Body = v
	return s
}

func (s *DeleteSparkTemplateFileResponse) Validate() error {
	return dara.Validate(s)
}
