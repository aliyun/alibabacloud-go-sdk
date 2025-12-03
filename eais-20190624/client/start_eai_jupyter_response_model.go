// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartEaiJupyterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartEaiJupyterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartEaiJupyterResponse
	GetStatusCode() *int32
	SetBody(v *StartEaiJupyterResponseBody) *StartEaiJupyterResponse
	GetBody() *StartEaiJupyterResponseBody
}

type StartEaiJupyterResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartEaiJupyterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartEaiJupyterResponse) String() string {
	return dara.Prettify(s)
}

func (s StartEaiJupyterResponse) GoString() string {
	return s.String()
}

func (s *StartEaiJupyterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartEaiJupyterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartEaiJupyterResponse) GetBody() *StartEaiJupyterResponseBody {
	return s.Body
}

func (s *StartEaiJupyterResponse) SetHeaders(v map[string]*string) *StartEaiJupyterResponse {
	s.Headers = v
	return s
}

func (s *StartEaiJupyterResponse) SetStatusCode(v int32) *StartEaiJupyterResponse {
	s.StatusCode = &v
	return s
}

func (s *StartEaiJupyterResponse) SetBody(v *StartEaiJupyterResponseBody) *StartEaiJupyterResponse {
	s.Body = v
	return s
}

func (s *StartEaiJupyterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
