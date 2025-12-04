// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestartInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestartInstanceResponse
	GetStatusCode() *int32
	SetBody(v *RestartInstanceResponseBody) *RestartInstanceResponse
	GetBody() *RestartInstanceResponseBody
}

type RestartInstanceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s RestartInstanceResponse) GoString() string {
	return s.String()
}

func (s *RestartInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestartInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestartInstanceResponse) GetBody() *RestartInstanceResponseBody {
	return s.Body
}

func (s *RestartInstanceResponse) SetHeaders(v map[string]*string) *RestartInstanceResponse {
	s.Headers = v
	return s
}

func (s *RestartInstanceResponse) SetStatusCode(v int32) *RestartInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartInstanceResponse) SetBody(v *RestartInstanceResponseBody) *RestartInstanceResponse {
	s.Body = v
	return s
}

func (s *RestartInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
