// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPlayerConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPlayerConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPlayerConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetPlayerConfigResponseBody) *GetPlayerConfigResponse
	GetBody() *GetPlayerConfigResponseBody
}

type GetPlayerConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPlayerConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPlayerConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPlayerConfigResponse) GoString() string {
	return s.String()
}

func (s *GetPlayerConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPlayerConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPlayerConfigResponse) GetBody() *GetPlayerConfigResponseBody {
	return s.Body
}

func (s *GetPlayerConfigResponse) SetHeaders(v map[string]*string) *GetPlayerConfigResponse {
	s.Headers = v
	return s
}

func (s *GetPlayerConfigResponse) SetStatusCode(v int32) *GetPlayerConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPlayerConfigResponse) SetBody(v *GetPlayerConfigResponseBody) *GetPlayerConfigResponse {
	s.Body = v
	return s
}

func (s *GetPlayerConfigResponse) Validate() error {
	return dara.Validate(s)
}
