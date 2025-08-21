// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEpnInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEpnInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEpnInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreateEpnInstanceResponseBody) *CreateEpnInstanceResponse
	GetBody() *CreateEpnInstanceResponseBody
}

type CreateEpnInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEpnInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEpnInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEpnInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateEpnInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEpnInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEpnInstanceResponse) GetBody() *CreateEpnInstanceResponseBody {
	return s.Body
}

func (s *CreateEpnInstanceResponse) SetHeaders(v map[string]*string) *CreateEpnInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateEpnInstanceResponse) SetStatusCode(v int32) *CreateEpnInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEpnInstanceResponse) SetBody(v *CreateEpnInstanceResponseBody) *CreateEpnInstanceResponse {
	s.Body = v
	return s
}

func (s *CreateEpnInstanceResponse) Validate() error {
	return dara.Validate(s)
}
