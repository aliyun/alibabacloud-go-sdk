// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOneMetaOssieModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOneMetaOssieModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOneMetaOssieModelResponse
	GetStatusCode() *int32
	SetBody(v *GetOneMetaOssieModelResponseBody) *GetOneMetaOssieModelResponse
	GetBody() *GetOneMetaOssieModelResponseBody
}

type GetOneMetaOssieModelResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOneMetaOssieModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOneMetaOssieModelResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOneMetaOssieModelResponse) GoString() string {
	return s.String()
}

func (s *GetOneMetaOssieModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOneMetaOssieModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOneMetaOssieModelResponse) GetBody() *GetOneMetaOssieModelResponseBody {
	return s.Body
}

func (s *GetOneMetaOssieModelResponse) SetHeaders(v map[string]*string) *GetOneMetaOssieModelResponse {
	s.Headers = v
	return s
}

func (s *GetOneMetaOssieModelResponse) SetStatusCode(v int32) *GetOneMetaOssieModelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOneMetaOssieModelResponse) SetBody(v *GetOneMetaOssieModelResponseBody) *GetOneMetaOssieModelResponse {
	s.Body = v
	return s
}

func (s *GetOneMetaOssieModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
