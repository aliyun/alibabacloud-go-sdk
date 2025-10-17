// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkTemplateFolderTreeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSparkTemplateFolderTreeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSparkTemplateFolderTreeResponse
	GetStatusCode() *int32
	SetBody(v *GetSparkTemplateFolderTreeResponseBody) *GetSparkTemplateFolderTreeResponse
	GetBody() *GetSparkTemplateFolderTreeResponseBody
}

type GetSparkTemplateFolderTreeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSparkTemplateFolderTreeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSparkTemplateFolderTreeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSparkTemplateFolderTreeResponse) GoString() string {
	return s.String()
}

func (s *GetSparkTemplateFolderTreeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSparkTemplateFolderTreeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSparkTemplateFolderTreeResponse) GetBody() *GetSparkTemplateFolderTreeResponseBody {
	return s.Body
}

func (s *GetSparkTemplateFolderTreeResponse) SetHeaders(v map[string]*string) *GetSparkTemplateFolderTreeResponse {
	s.Headers = v
	return s
}

func (s *GetSparkTemplateFolderTreeResponse) SetStatusCode(v int32) *GetSparkTemplateFolderTreeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSparkTemplateFolderTreeResponse) SetBody(v *GetSparkTemplateFolderTreeResponseBody) *GetSparkTemplateFolderTreeResponse {
	s.Body = v
	return s
}

func (s *GetSparkTemplateFolderTreeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
