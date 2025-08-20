// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkTemplateFileContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSparkTemplateFileContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSparkTemplateFileContentResponse
	GetStatusCode() *int32
	SetBody(v *GetSparkTemplateFileContentResponseBody) *GetSparkTemplateFileContentResponse
	GetBody() *GetSparkTemplateFileContentResponseBody
}

type GetSparkTemplateFileContentResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSparkTemplateFileContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSparkTemplateFileContentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSparkTemplateFileContentResponse) GoString() string {
	return s.String()
}

func (s *GetSparkTemplateFileContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSparkTemplateFileContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSparkTemplateFileContentResponse) GetBody() *GetSparkTemplateFileContentResponseBody {
	return s.Body
}

func (s *GetSparkTemplateFileContentResponse) SetHeaders(v map[string]*string) *GetSparkTemplateFileContentResponse {
	s.Headers = v
	return s
}

func (s *GetSparkTemplateFileContentResponse) SetStatusCode(v int32) *GetSparkTemplateFileContentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSparkTemplateFileContentResponse) SetBody(v *GetSparkTemplateFileContentResponseBody) *GetSparkTemplateFileContentResponse {
	s.Body = v
	return s
}

func (s *GetSparkTemplateFileContentResponse) Validate() error {
	return dara.Validate(s)
}
