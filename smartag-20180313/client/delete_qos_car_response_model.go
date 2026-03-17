// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQosCarResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteQosCarResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteQosCarResponse
	GetStatusCode() *int32
	SetBody(v *DeleteQosCarResponseBody) *DeleteQosCarResponse
	GetBody() *DeleteQosCarResponseBody
}

type DeleteQosCarResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteQosCarResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteQosCarResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteQosCarResponse) GoString() string {
	return s.String()
}

func (s *DeleteQosCarResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteQosCarResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteQosCarResponse) GetBody() *DeleteQosCarResponseBody {
	return s.Body
}

func (s *DeleteQosCarResponse) SetHeaders(v map[string]*string) *DeleteQosCarResponse {
	s.Headers = v
	return s
}

func (s *DeleteQosCarResponse) SetStatusCode(v int32) *DeleteQosCarResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteQosCarResponse) SetBody(v *DeleteQosCarResponseBody) *DeleteQosCarResponse {
	s.Body = v
	return s
}

func (s *DeleteQosCarResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
