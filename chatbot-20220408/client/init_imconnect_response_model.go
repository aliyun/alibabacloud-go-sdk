// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitIMConnectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InitIMConnectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InitIMConnectResponse
	GetStatusCode() *int32
	SetBody(v *InitIMConnectResponseBody) *InitIMConnectResponse
	GetBody() *InitIMConnectResponseBody
}

type InitIMConnectResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitIMConnectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitIMConnectResponse) String() string {
	return dara.Prettify(s)
}

func (s InitIMConnectResponse) GoString() string {
	return s.String()
}

func (s *InitIMConnectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InitIMConnectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InitIMConnectResponse) GetBody() *InitIMConnectResponseBody {
	return s.Body
}

func (s *InitIMConnectResponse) SetHeaders(v map[string]*string) *InitIMConnectResponse {
	s.Headers = v
	return s
}

func (s *InitIMConnectResponse) SetStatusCode(v int32) *InitIMConnectResponse {
	s.StatusCode = &v
	return s
}

func (s *InitIMConnectResponse) SetBody(v *InitIMConnectResponseBody) *InitIMConnectResponse {
	s.Body = v
	return s
}

func (s *InitIMConnectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
