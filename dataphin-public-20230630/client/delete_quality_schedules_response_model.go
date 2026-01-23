// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQualitySchedulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteQualitySchedulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteQualitySchedulesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteQualitySchedulesResponseBody) *DeleteQualitySchedulesResponse
	GetBody() *DeleteQualitySchedulesResponseBody
}

type DeleteQualitySchedulesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteQualitySchedulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteQualitySchedulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteQualitySchedulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteQualitySchedulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteQualitySchedulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteQualitySchedulesResponse) GetBody() *DeleteQualitySchedulesResponseBody {
	return s.Body
}

func (s *DeleteQualitySchedulesResponse) SetHeaders(v map[string]*string) *DeleteQualitySchedulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteQualitySchedulesResponse) SetStatusCode(v int32) *DeleteQualitySchedulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteQualitySchedulesResponse) SetBody(v *DeleteQualitySchedulesResponseBody) *DeleteQualitySchedulesResponse {
	s.Body = v
	return s
}

func (s *DeleteQualitySchedulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
