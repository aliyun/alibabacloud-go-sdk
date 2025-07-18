// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachApplication2ConnectorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachApplication2ConnectorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachApplication2ConnectorResponse
	GetStatusCode() *int32
	SetBody(v *AttachApplication2ConnectorResponseBody) *AttachApplication2ConnectorResponse
	GetBody() *AttachApplication2ConnectorResponseBody
}

type AttachApplication2ConnectorResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachApplication2ConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachApplication2ConnectorResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachApplication2ConnectorResponse) GoString() string {
	return s.String()
}

func (s *AttachApplication2ConnectorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachApplication2ConnectorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachApplication2ConnectorResponse) GetBody() *AttachApplication2ConnectorResponseBody {
	return s.Body
}

func (s *AttachApplication2ConnectorResponse) SetHeaders(v map[string]*string) *AttachApplication2ConnectorResponse {
	s.Headers = v
	return s
}

func (s *AttachApplication2ConnectorResponse) SetStatusCode(v int32) *AttachApplication2ConnectorResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachApplication2ConnectorResponse) SetBody(v *AttachApplication2ConnectorResponseBody) *AttachApplication2ConnectorResponse {
	s.Body = v
	return s
}

func (s *AttachApplication2ConnectorResponse) Validate() error {
	return dara.Validate(s)
}
