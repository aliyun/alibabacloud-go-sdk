// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachPluginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachPluginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachPluginResponse
	GetStatusCode() *int32
	SetBody(v *DetachPluginResponseBody) *DetachPluginResponse
	GetBody() *DetachPluginResponseBody
}

type DetachPluginResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachPluginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachPluginResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachPluginResponse) GoString() string {
	return s.String()
}

func (s *DetachPluginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachPluginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachPluginResponse) GetBody() *DetachPluginResponseBody {
	return s.Body
}

func (s *DetachPluginResponse) SetHeaders(v map[string]*string) *DetachPluginResponse {
	s.Headers = v
	return s
}

func (s *DetachPluginResponse) SetStatusCode(v int32) *DetachPluginResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachPluginResponse) SetBody(v *DetachPluginResponseBody) *DetachPluginResponse {
	s.Body = v
	return s
}

func (s *DetachPluginResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
