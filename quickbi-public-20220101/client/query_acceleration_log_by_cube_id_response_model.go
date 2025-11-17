// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAccelerationLogByCubeIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAccelerationLogByCubeIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAccelerationLogByCubeIdResponse
	GetStatusCode() *int32
	SetBody(v *QueryAccelerationLogByCubeIdResponseBody) *QueryAccelerationLogByCubeIdResponse
	GetBody() *QueryAccelerationLogByCubeIdResponseBody
}

type QueryAccelerationLogByCubeIdResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAccelerationLogByCubeIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAccelerationLogByCubeIdResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAccelerationLogByCubeIdResponse) GoString() string {
	return s.String()
}

func (s *QueryAccelerationLogByCubeIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAccelerationLogByCubeIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAccelerationLogByCubeIdResponse) GetBody() *QueryAccelerationLogByCubeIdResponseBody {
	return s.Body
}

func (s *QueryAccelerationLogByCubeIdResponse) SetHeaders(v map[string]*string) *QueryAccelerationLogByCubeIdResponse {
	s.Headers = v
	return s
}

func (s *QueryAccelerationLogByCubeIdResponse) SetStatusCode(v int32) *QueryAccelerationLogByCubeIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAccelerationLogByCubeIdResponse) SetBody(v *QueryAccelerationLogByCubeIdResponseBody) *QueryAccelerationLogByCubeIdResponse {
	s.Body = v
	return s
}

func (s *QueryAccelerationLogByCubeIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
