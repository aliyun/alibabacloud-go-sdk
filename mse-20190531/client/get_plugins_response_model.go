// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPluginsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPluginsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPluginsResponse
	GetStatusCode() *int32
	SetBody(v *GetPluginsResponseBody) *GetPluginsResponse
	GetBody() *GetPluginsResponseBody
}

type GetPluginsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPluginsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPluginsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPluginsResponse) GoString() string {
	return s.String()
}

func (s *GetPluginsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPluginsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPluginsResponse) GetBody() *GetPluginsResponseBody {
	return s.Body
}

func (s *GetPluginsResponse) SetHeaders(v map[string]*string) *GetPluginsResponse {
	s.Headers = v
	return s
}

func (s *GetPluginsResponse) SetStatusCode(v int32) *GetPluginsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPluginsResponse) SetBody(v *GetPluginsResponseBody) *GetPluginsResponse {
	s.Body = v
	return s
}

func (s *GetPluginsResponse) Validate() error {
	return dara.Validate(s)
}
