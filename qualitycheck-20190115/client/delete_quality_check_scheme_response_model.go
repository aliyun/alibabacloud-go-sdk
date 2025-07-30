// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQualityCheckSchemeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteQualityCheckSchemeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteQualityCheckSchemeResponse
	GetStatusCode() *int32
	SetBody(v *DeleteQualityCheckSchemeResponseBody) *DeleteQualityCheckSchemeResponse
	GetBody() *DeleteQualityCheckSchemeResponseBody
}

type DeleteQualityCheckSchemeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteQualityCheckSchemeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteQualityCheckSchemeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteQualityCheckSchemeResponse) GoString() string {
	return s.String()
}

func (s *DeleteQualityCheckSchemeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteQualityCheckSchemeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteQualityCheckSchemeResponse) GetBody() *DeleteQualityCheckSchemeResponseBody {
	return s.Body
}

func (s *DeleteQualityCheckSchemeResponse) SetHeaders(v map[string]*string) *DeleteQualityCheckSchemeResponse {
	s.Headers = v
	return s
}

func (s *DeleteQualityCheckSchemeResponse) SetStatusCode(v int32) *DeleteQualityCheckSchemeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteQualityCheckSchemeResponse) SetBody(v *DeleteQualityCheckSchemeResponseBody) *DeleteQualityCheckSchemeResponse {
	s.Body = v
	return s
}

func (s *DeleteQualityCheckSchemeResponse) Validate() error {
	return dara.Validate(s)
}
