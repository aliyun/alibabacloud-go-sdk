// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOneMetaOssieModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteOneMetaOssieModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteOneMetaOssieModelResponse
	GetStatusCode() *int32
	SetBody(v *DeleteOneMetaOssieModelResponseBody) *DeleteOneMetaOssieModelResponse
	GetBody() *DeleteOneMetaOssieModelResponseBody
}

type DeleteOneMetaOssieModelResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteOneMetaOssieModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteOneMetaOssieModelResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteOneMetaOssieModelResponse) GoString() string {
	return s.String()
}

func (s *DeleteOneMetaOssieModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteOneMetaOssieModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteOneMetaOssieModelResponse) GetBody() *DeleteOneMetaOssieModelResponseBody {
	return s.Body
}

func (s *DeleteOneMetaOssieModelResponse) SetHeaders(v map[string]*string) *DeleteOneMetaOssieModelResponse {
	s.Headers = v
	return s
}

func (s *DeleteOneMetaOssieModelResponse) SetStatusCode(v int32) *DeleteOneMetaOssieModelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteOneMetaOssieModelResponse) SetBody(v *DeleteOneMetaOssieModelResponseBody) *DeleteOneMetaOssieModelResponse {
	s.Body = v
	return s
}

func (s *DeleteOneMetaOssieModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
