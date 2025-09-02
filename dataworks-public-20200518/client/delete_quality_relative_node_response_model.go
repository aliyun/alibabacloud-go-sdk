// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQualityRelativeNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteQualityRelativeNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteQualityRelativeNodeResponse
	GetStatusCode() *int32
	SetBody(v *DeleteQualityRelativeNodeResponseBody) *DeleteQualityRelativeNodeResponse
	GetBody() *DeleteQualityRelativeNodeResponseBody
}

type DeleteQualityRelativeNodeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteQualityRelativeNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteQualityRelativeNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteQualityRelativeNodeResponse) GoString() string {
	return s.String()
}

func (s *DeleteQualityRelativeNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteQualityRelativeNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteQualityRelativeNodeResponse) GetBody() *DeleteQualityRelativeNodeResponseBody {
	return s.Body
}

func (s *DeleteQualityRelativeNodeResponse) SetHeaders(v map[string]*string) *DeleteQualityRelativeNodeResponse {
	s.Headers = v
	return s
}

func (s *DeleteQualityRelativeNodeResponse) SetStatusCode(v int32) *DeleteQualityRelativeNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteQualityRelativeNodeResponse) SetBody(v *DeleteQualityRelativeNodeResponseBody) *DeleteQualityRelativeNodeResponse {
	s.Body = v
	return s
}

func (s *DeleteQualityRelativeNodeResponse) Validate() error {
	return dara.Validate(s)
}
