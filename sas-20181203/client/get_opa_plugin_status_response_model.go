// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpaPluginStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOpaPluginStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOpaPluginStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetOpaPluginStatusResponseBody) *GetOpaPluginStatusResponse
	GetBody() *GetOpaPluginStatusResponseBody
}

type GetOpaPluginStatusResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOpaPluginStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOpaPluginStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOpaPluginStatusResponse) GoString() string {
	return s.String()
}

func (s *GetOpaPluginStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOpaPluginStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOpaPluginStatusResponse) GetBody() *GetOpaPluginStatusResponseBody {
	return s.Body
}

func (s *GetOpaPluginStatusResponse) SetHeaders(v map[string]*string) *GetOpaPluginStatusResponse {
	s.Headers = v
	return s
}

func (s *GetOpaPluginStatusResponse) SetStatusCode(v int32) *GetOpaPluginStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOpaPluginStatusResponse) SetBody(v *GetOpaPluginStatusResponseBody) *GetOpaPluginStatusResponse {
	s.Body = v
	return s
}

func (s *GetOpaPluginStatusResponse) Validate() error {
	return dara.Validate(s)
}
