// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQualityEntityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteQualityEntityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteQualityEntityResponse
	GetStatusCode() *int32
	SetBody(v *DeleteQualityEntityResponseBody) *DeleteQualityEntityResponse
	GetBody() *DeleteQualityEntityResponseBody
}

type DeleteQualityEntityResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteQualityEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteQualityEntityResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteQualityEntityResponse) GoString() string {
	return s.String()
}

func (s *DeleteQualityEntityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteQualityEntityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteQualityEntityResponse) GetBody() *DeleteQualityEntityResponseBody {
	return s.Body
}

func (s *DeleteQualityEntityResponse) SetHeaders(v map[string]*string) *DeleteQualityEntityResponse {
	s.Headers = v
	return s
}

func (s *DeleteQualityEntityResponse) SetStatusCode(v int32) *DeleteQualityEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteQualityEntityResponse) SetBody(v *DeleteQualityEntityResponseBody) *DeleteQualityEntityResponse {
	s.Body = v
	return s
}

func (s *DeleteQualityEntityResponse) Validate() error {
	return dara.Validate(s)
}
