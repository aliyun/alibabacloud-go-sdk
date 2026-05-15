// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQualityProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteQualityProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteQualityProjectResponse
	GetStatusCode() *int32
	SetBody(v *DeleteQualityProjectResponseBody) *DeleteQualityProjectResponse
	GetBody() *DeleteQualityProjectResponseBody
}

type DeleteQualityProjectResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteQualityProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteQualityProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteQualityProjectResponse) GoString() string {
	return s.String()
}

func (s *DeleteQualityProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteQualityProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteQualityProjectResponse) GetBody() *DeleteQualityProjectResponseBody {
	return s.Body
}

func (s *DeleteQualityProjectResponse) SetHeaders(v map[string]*string) *DeleteQualityProjectResponse {
	s.Headers = v
	return s
}

func (s *DeleteQualityProjectResponse) SetStatusCode(v int32) *DeleteQualityProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteQualityProjectResponse) SetBody(v *DeleteQualityProjectResponseBody) *DeleteQualityProjectResponse {
	s.Body = v
	return s
}

func (s *DeleteQualityProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
