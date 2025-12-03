// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEaiJupyterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEaiJupyterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEaiJupyterResponse
	GetStatusCode() *int32
	SetBody(v *CreateEaiJupyterResponseBody) *CreateEaiJupyterResponse
	GetBody() *CreateEaiJupyterResponseBody
}

type CreateEaiJupyterResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEaiJupyterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEaiJupyterResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEaiJupyterResponse) GoString() string {
	return s.String()
}

func (s *CreateEaiJupyterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEaiJupyterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEaiJupyterResponse) GetBody() *CreateEaiJupyterResponseBody {
	return s.Body
}

func (s *CreateEaiJupyterResponse) SetHeaders(v map[string]*string) *CreateEaiJupyterResponse {
	s.Headers = v
	return s
}

func (s *CreateEaiJupyterResponse) SetStatusCode(v int32) *CreateEaiJupyterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEaiJupyterResponse) SetBody(v *CreateEaiJupyterResponseBody) *CreateEaiJupyterResponse {
	s.Body = v
	return s
}

func (s *CreateEaiJupyterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
