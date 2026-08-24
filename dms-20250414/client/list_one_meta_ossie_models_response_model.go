// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOneMetaOssieModelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOneMetaOssieModelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOneMetaOssieModelsResponse
	GetStatusCode() *int32
	SetBody(v *ListOneMetaOssieModelsResponseBody) *ListOneMetaOssieModelsResponse
	GetBody() *ListOneMetaOssieModelsResponseBody
}

type ListOneMetaOssieModelsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOneMetaOssieModelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOneMetaOssieModelsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOneMetaOssieModelsResponse) GoString() string {
	return s.String()
}

func (s *ListOneMetaOssieModelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOneMetaOssieModelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOneMetaOssieModelsResponse) GetBody() *ListOneMetaOssieModelsResponseBody {
	return s.Body
}

func (s *ListOneMetaOssieModelsResponse) SetHeaders(v map[string]*string) *ListOneMetaOssieModelsResponse {
	s.Headers = v
	return s
}

func (s *ListOneMetaOssieModelsResponse) SetStatusCode(v int32) *ListOneMetaOssieModelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOneMetaOssieModelsResponse) SetBody(v *ListOneMetaOssieModelsResponseBody) *ListOneMetaOssieModelsResponse {
	s.Body = v
	return s
}

func (s *ListOneMetaOssieModelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
