// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQualityWatchesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteQualityWatchesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteQualityWatchesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteQualityWatchesResponseBody) *DeleteQualityWatchesResponse
	GetBody() *DeleteQualityWatchesResponseBody
}

type DeleteQualityWatchesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteQualityWatchesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteQualityWatchesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteQualityWatchesResponse) GoString() string {
	return s.String()
}

func (s *DeleteQualityWatchesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteQualityWatchesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteQualityWatchesResponse) GetBody() *DeleteQualityWatchesResponseBody {
	return s.Body
}

func (s *DeleteQualityWatchesResponse) SetHeaders(v map[string]*string) *DeleteQualityWatchesResponse {
	s.Headers = v
	return s
}

func (s *DeleteQualityWatchesResponse) SetStatusCode(v int32) *DeleteQualityWatchesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteQualityWatchesResponse) SetBody(v *DeleteQualityWatchesResponseBody) *DeleteQualityWatchesResponse {
	s.Body = v
	return s
}

func (s *DeleteQualityWatchesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
