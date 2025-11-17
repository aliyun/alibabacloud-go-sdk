// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLastAccelerationEngineJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryLastAccelerationEngineJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryLastAccelerationEngineJobResponse
	GetStatusCode() *int32
	SetBody(v *QueryLastAccelerationEngineJobResponseBody) *QueryLastAccelerationEngineJobResponse
	GetBody() *QueryLastAccelerationEngineJobResponseBody
}

type QueryLastAccelerationEngineJobResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryLastAccelerationEngineJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryLastAccelerationEngineJobResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryLastAccelerationEngineJobResponse) GoString() string {
	return s.String()
}

func (s *QueryLastAccelerationEngineJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryLastAccelerationEngineJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryLastAccelerationEngineJobResponse) GetBody() *QueryLastAccelerationEngineJobResponseBody {
	return s.Body
}

func (s *QueryLastAccelerationEngineJobResponse) SetHeaders(v map[string]*string) *QueryLastAccelerationEngineJobResponse {
	s.Headers = v
	return s
}

func (s *QueryLastAccelerationEngineJobResponse) SetStatusCode(v int32) *QueryLastAccelerationEngineJobResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryLastAccelerationEngineJobResponse) SetBody(v *QueryLastAccelerationEngineJobResponseBody) *QueryLastAccelerationEngineJobResponse {
	s.Body = v
	return s
}

func (s *QueryLastAccelerationEngineJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
