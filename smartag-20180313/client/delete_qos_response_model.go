// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteQosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteQosResponse
	GetStatusCode() *int32
	SetBody(v *DeleteQosResponseBody) *DeleteQosResponse
	GetBody() *DeleteQosResponseBody
}

type DeleteQosResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteQosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteQosResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteQosResponse) GoString() string {
	return s.String()
}

func (s *DeleteQosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteQosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteQosResponse) GetBody() *DeleteQosResponseBody {
	return s.Body
}

func (s *DeleteQosResponse) SetHeaders(v map[string]*string) *DeleteQosResponse {
	s.Headers = v
	return s
}

func (s *DeleteQosResponse) SetStatusCode(v int32) *DeleteQosResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteQosResponse) SetBody(v *DeleteQosResponseBody) *DeleteQosResponse {
	s.Body = v
	return s
}

func (s *DeleteQosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
