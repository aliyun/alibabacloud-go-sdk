// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQualityFollowerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteQualityFollowerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteQualityFollowerResponse
	GetStatusCode() *int32
	SetBody(v *DeleteQualityFollowerResponseBody) *DeleteQualityFollowerResponse
	GetBody() *DeleteQualityFollowerResponseBody
}

type DeleteQualityFollowerResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteQualityFollowerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteQualityFollowerResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteQualityFollowerResponse) GoString() string {
	return s.String()
}

func (s *DeleteQualityFollowerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteQualityFollowerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteQualityFollowerResponse) GetBody() *DeleteQualityFollowerResponseBody {
	return s.Body
}

func (s *DeleteQualityFollowerResponse) SetHeaders(v map[string]*string) *DeleteQualityFollowerResponse {
	s.Headers = v
	return s
}

func (s *DeleteQualityFollowerResponse) SetStatusCode(v int32) *DeleteQualityFollowerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteQualityFollowerResponse) SetBody(v *DeleteQualityFollowerResponseBody) *DeleteQualityFollowerResponse {
	s.Body = v
	return s
}

func (s *DeleteQualityFollowerResponse) Validate() error {
	return dara.Validate(s)
}
