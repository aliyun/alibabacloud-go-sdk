// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGrafanaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddGrafanaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddGrafanaResponse
	GetStatusCode() *int32
	SetBody(v *AddGrafanaResponseBody) *AddGrafanaResponse
	GetBody() *AddGrafanaResponseBody
}

type AddGrafanaResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddGrafanaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddGrafanaResponse) String() string {
	return dara.Prettify(s)
}

func (s AddGrafanaResponse) GoString() string {
	return s.String()
}

func (s *AddGrafanaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddGrafanaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddGrafanaResponse) GetBody() *AddGrafanaResponseBody {
	return s.Body
}

func (s *AddGrafanaResponse) SetHeaders(v map[string]*string) *AddGrafanaResponse {
	s.Headers = v
	return s
}

func (s *AddGrafanaResponse) SetStatusCode(v int32) *AddGrafanaResponse {
	s.StatusCode = &v
	return s
}

func (s *AddGrafanaResponse) SetBody(v *AddGrafanaResponseBody) *AddGrafanaResponse {
	s.Body = v
	return s
}

func (s *AddGrafanaResponse) Validate() error {
	return dara.Validate(s)
}
