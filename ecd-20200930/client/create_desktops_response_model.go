// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDesktopsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDesktopsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDesktopsResponse
	GetStatusCode() *int32
	SetBody(v *CreateDesktopsResponseBody) *CreateDesktopsResponse
	GetBody() *CreateDesktopsResponseBody
}

type CreateDesktopsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDesktopsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDesktopsResponse) GoString() string {
	return s.String()
}

func (s *CreateDesktopsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDesktopsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDesktopsResponse) GetBody() *CreateDesktopsResponseBody {
	return s.Body
}

func (s *CreateDesktopsResponse) SetHeaders(v map[string]*string) *CreateDesktopsResponse {
	s.Headers = v
	return s
}

func (s *CreateDesktopsResponse) SetStatusCode(v int32) *CreateDesktopsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDesktopsResponse) SetBody(v *CreateDesktopsResponseBody) *CreateDesktopsResponse {
	s.Body = v
	return s
}

func (s *CreateDesktopsResponse) Validate() error {
	return dara.Validate(s)
}
