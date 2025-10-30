// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachApplication2ConnectorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachApplication2ConnectorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachApplication2ConnectorResponse
	GetStatusCode() *int32
	SetBody(v *DetachApplication2ConnectorResponseBody) *DetachApplication2ConnectorResponse
	GetBody() *DetachApplication2ConnectorResponseBody
}

type DetachApplication2ConnectorResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachApplication2ConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachApplication2ConnectorResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachApplication2ConnectorResponse) GoString() string {
	return s.String()
}

func (s *DetachApplication2ConnectorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachApplication2ConnectorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachApplication2ConnectorResponse) GetBody() *DetachApplication2ConnectorResponseBody {
	return s.Body
}

func (s *DetachApplication2ConnectorResponse) SetHeaders(v map[string]*string) *DetachApplication2ConnectorResponse {
	s.Headers = v
	return s
}

func (s *DetachApplication2ConnectorResponse) SetStatusCode(v int32) *DetachApplication2ConnectorResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachApplication2ConnectorResponse) SetBody(v *DetachApplication2ConnectorResponseBody) *DetachApplication2ConnectorResponse {
	s.Body = v
	return s
}

func (s *DetachApplication2ConnectorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
