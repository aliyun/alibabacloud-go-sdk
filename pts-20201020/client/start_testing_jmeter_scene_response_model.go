// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartTestingJMeterSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartTestingJMeterSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartTestingJMeterSceneResponse
	GetStatusCode() *int32
	SetBody(v *StartTestingJMeterSceneResponseBody) *StartTestingJMeterSceneResponse
	GetBody() *StartTestingJMeterSceneResponseBody
}

type StartTestingJMeterSceneResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartTestingJMeterSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartTestingJMeterSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s StartTestingJMeterSceneResponse) GoString() string {
	return s.String()
}

func (s *StartTestingJMeterSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartTestingJMeterSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartTestingJMeterSceneResponse) GetBody() *StartTestingJMeterSceneResponseBody {
	return s.Body
}

func (s *StartTestingJMeterSceneResponse) SetHeaders(v map[string]*string) *StartTestingJMeterSceneResponse {
	s.Headers = v
	return s
}

func (s *StartTestingJMeterSceneResponse) SetStatusCode(v int32) *StartTestingJMeterSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *StartTestingJMeterSceneResponse) SetBody(v *StartTestingJMeterSceneResponseBody) *StartTestingJMeterSceneResponse {
	s.Body = v
	return s
}

func (s *StartTestingJMeterSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
