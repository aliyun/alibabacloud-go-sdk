// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOneMetaOssieModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateOneMetaOssieModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateOneMetaOssieModelResponse
	GetStatusCode() *int32
	SetBody(v *UpdateOneMetaOssieModelResponseBody) *UpdateOneMetaOssieModelResponse
	GetBody() *UpdateOneMetaOssieModelResponseBody
}

type UpdateOneMetaOssieModelResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateOneMetaOssieModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateOneMetaOssieModelResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateOneMetaOssieModelResponse) GoString() string {
	return s.String()
}

func (s *UpdateOneMetaOssieModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateOneMetaOssieModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateOneMetaOssieModelResponse) GetBody() *UpdateOneMetaOssieModelResponseBody {
	return s.Body
}

func (s *UpdateOneMetaOssieModelResponse) SetHeaders(v map[string]*string) *UpdateOneMetaOssieModelResponse {
	s.Headers = v
	return s
}

func (s *UpdateOneMetaOssieModelResponse) SetStatusCode(v int32) *UpdateOneMetaOssieModelResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOneMetaOssieModelResponse) SetBody(v *UpdateOneMetaOssieModelResponseBody) *UpdateOneMetaOssieModelResponse {
	s.Body = v
	return s
}

func (s *UpdateOneMetaOssieModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
