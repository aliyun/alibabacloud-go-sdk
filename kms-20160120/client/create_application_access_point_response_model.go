// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationAccessPointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateApplicationAccessPointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateApplicationAccessPointResponse
	GetStatusCode() *int32
	SetBody(v *CreateApplicationAccessPointResponseBody) *CreateApplicationAccessPointResponse
	GetBody() *CreateApplicationAccessPointResponseBody
}

type CreateApplicationAccessPointResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApplicationAccessPointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApplicationAccessPointResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationAccessPointResponse) GoString() string {
	return s.String()
}

func (s *CreateApplicationAccessPointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateApplicationAccessPointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateApplicationAccessPointResponse) GetBody() *CreateApplicationAccessPointResponseBody {
	return s.Body
}

func (s *CreateApplicationAccessPointResponse) SetHeaders(v map[string]*string) *CreateApplicationAccessPointResponse {
	s.Headers = v
	return s
}

func (s *CreateApplicationAccessPointResponse) SetStatusCode(v int32) *CreateApplicationAccessPointResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApplicationAccessPointResponse) SetBody(v *CreateApplicationAccessPointResponseBody) *CreateApplicationAccessPointResponse {
	s.Body = v
	return s
}

func (s *CreateApplicationAccessPointResponse) Validate() error {
	return dara.Validate(s)
}
