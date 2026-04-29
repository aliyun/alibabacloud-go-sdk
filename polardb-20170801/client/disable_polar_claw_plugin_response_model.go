// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisablePolarClawPluginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisablePolarClawPluginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisablePolarClawPluginResponse
	GetStatusCode() *int32
	SetBody(v *DisablePolarClawPluginResponseBody) *DisablePolarClawPluginResponse
	GetBody() *DisablePolarClawPluginResponseBody
}

type DisablePolarClawPluginResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisablePolarClawPluginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisablePolarClawPluginResponse) String() string {
	return dara.Prettify(s)
}

func (s DisablePolarClawPluginResponse) GoString() string {
	return s.String()
}

func (s *DisablePolarClawPluginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisablePolarClawPluginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisablePolarClawPluginResponse) GetBody() *DisablePolarClawPluginResponseBody {
	return s.Body
}

func (s *DisablePolarClawPluginResponse) SetHeaders(v map[string]*string) *DisablePolarClawPluginResponse {
	s.Headers = v
	return s
}

func (s *DisablePolarClawPluginResponse) SetStatusCode(v int32) *DisablePolarClawPluginResponse {
	s.StatusCode = &v
	return s
}

func (s *DisablePolarClawPluginResponse) SetBody(v *DisablePolarClawPluginResponseBody) *DisablePolarClawPluginResponse {
	s.Body = v
	return s
}

func (s *DisablePolarClawPluginResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
