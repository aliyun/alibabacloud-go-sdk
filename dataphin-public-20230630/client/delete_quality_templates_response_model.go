// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQualityTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteQualityTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteQualityTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteQualityTemplatesResponseBody) *DeleteQualityTemplatesResponse
	GetBody() *DeleteQualityTemplatesResponseBody
}

type DeleteQualityTemplatesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteQualityTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteQualityTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteQualityTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DeleteQualityTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteQualityTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteQualityTemplatesResponse) GetBody() *DeleteQualityTemplatesResponseBody {
	return s.Body
}

func (s *DeleteQualityTemplatesResponse) SetHeaders(v map[string]*string) *DeleteQualityTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DeleteQualityTemplatesResponse) SetStatusCode(v int32) *DeleteQualityTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteQualityTemplatesResponse) SetBody(v *DeleteQualityTemplatesResponseBody) *DeleteQualityTemplatesResponse {
	s.Body = v
	return s
}

func (s *DeleteQualityTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
