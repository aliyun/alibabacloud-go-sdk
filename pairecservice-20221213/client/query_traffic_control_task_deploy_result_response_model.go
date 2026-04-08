// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTrafficControlTaskDeployResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryTrafficControlTaskDeployResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryTrafficControlTaskDeployResultResponse
	GetStatusCode() *int32
	SetBody(v *QueryTrafficControlTaskDeployResultResponseBody) *QueryTrafficControlTaskDeployResultResponse
	GetBody() *QueryTrafficControlTaskDeployResultResponseBody
}

type QueryTrafficControlTaskDeployResultResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTrafficControlTaskDeployResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTrafficControlTaskDeployResultResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryTrafficControlTaskDeployResultResponse) GoString() string {
	return s.String()
}

func (s *QueryTrafficControlTaskDeployResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryTrafficControlTaskDeployResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryTrafficControlTaskDeployResultResponse) GetBody() *QueryTrafficControlTaskDeployResultResponseBody {
	return s.Body
}

func (s *QueryTrafficControlTaskDeployResultResponse) SetHeaders(v map[string]*string) *QueryTrafficControlTaskDeployResultResponse {
	s.Headers = v
	return s
}

func (s *QueryTrafficControlTaskDeployResultResponse) SetStatusCode(v int32) *QueryTrafficControlTaskDeployResultResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTrafficControlTaskDeployResultResponse) SetBody(v *QueryTrafficControlTaskDeployResultResponseBody) *QueryTrafficControlTaskDeployResultResponse {
	s.Body = v
	return s
}

func (s *QueryTrafficControlTaskDeployResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
