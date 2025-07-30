// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWarningConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWarningConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWarningConfigResponse
	GetStatusCode() *int32
	SetBody(v *CreateWarningConfigResponseBody) *CreateWarningConfigResponse
	GetBody() *CreateWarningConfigResponseBody
}

type CreateWarningConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWarningConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWarningConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWarningConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateWarningConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWarningConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWarningConfigResponse) GetBody() *CreateWarningConfigResponseBody {
	return s.Body
}

func (s *CreateWarningConfigResponse) SetHeaders(v map[string]*string) *CreateWarningConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateWarningConfigResponse) SetStatusCode(v int32) *CreateWarningConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWarningConfigResponse) SetBody(v *CreateWarningConfigResponseBody) *CreateWarningConfigResponse {
	s.Body = v
	return s
}

func (s *CreateWarningConfigResponse) Validate() error {
	return dara.Validate(s)
}
