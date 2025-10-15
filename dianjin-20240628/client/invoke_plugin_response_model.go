// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokePluginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InvokePluginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InvokePluginResponse
	GetStatusCode() *int32
	SetBody(v *InvokePluginResponseBody) *InvokePluginResponse
	GetBody() *InvokePluginResponseBody
}

type InvokePluginResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InvokePluginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvokePluginResponse) String() string {
	return dara.Prettify(s)
}

func (s InvokePluginResponse) GoString() string {
	return s.String()
}

func (s *InvokePluginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InvokePluginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InvokePluginResponse) GetBody() *InvokePluginResponseBody {
	return s.Body
}

func (s *InvokePluginResponse) SetHeaders(v map[string]*string) *InvokePluginResponse {
	s.Headers = v
	return s
}

func (s *InvokePluginResponse) SetStatusCode(v int32) *InvokePluginResponse {
	s.StatusCode = &v
	return s
}

func (s *InvokePluginResponse) SetBody(v *InvokePluginResponseBody) *InvokePluginResponse {
	s.Body = v
	return s
}

func (s *InvokePluginResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
