// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportOneMetaOssieModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportOneMetaOssieModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportOneMetaOssieModelResponse
	GetStatusCode() *int32
	SetBody(v *ImportOneMetaOssieModelResponseBody) *ImportOneMetaOssieModelResponse
	GetBody() *ImportOneMetaOssieModelResponseBody
}

type ImportOneMetaOssieModelResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportOneMetaOssieModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportOneMetaOssieModelResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportOneMetaOssieModelResponse) GoString() string {
	return s.String()
}

func (s *ImportOneMetaOssieModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportOneMetaOssieModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportOneMetaOssieModelResponse) GetBody() *ImportOneMetaOssieModelResponseBody {
	return s.Body
}

func (s *ImportOneMetaOssieModelResponse) SetHeaders(v map[string]*string) *ImportOneMetaOssieModelResponse {
	s.Headers = v
	return s
}

func (s *ImportOneMetaOssieModelResponse) SetStatusCode(v int32) *ImportOneMetaOssieModelResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportOneMetaOssieModelResponse) SetBody(v *ImportOneMetaOssieModelResponseBody) *ImportOneMetaOssieModelResponse {
	s.Body = v
	return s
}

func (s *ImportOneMetaOssieModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
