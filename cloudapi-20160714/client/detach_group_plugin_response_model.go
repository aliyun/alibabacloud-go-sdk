// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachGroupPluginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachGroupPluginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachGroupPluginResponse
	GetStatusCode() *int32
	SetBody(v *DetachGroupPluginResponseBody) *DetachGroupPluginResponse
	GetBody() *DetachGroupPluginResponseBody
}

type DetachGroupPluginResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachGroupPluginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachGroupPluginResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachGroupPluginResponse) GoString() string {
	return s.String()
}

func (s *DetachGroupPluginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachGroupPluginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachGroupPluginResponse) GetBody() *DetachGroupPluginResponseBody {
	return s.Body
}

func (s *DetachGroupPluginResponse) SetHeaders(v map[string]*string) *DetachGroupPluginResponse {
	s.Headers = v
	return s
}

func (s *DetachGroupPluginResponse) SetStatusCode(v int32) *DetachGroupPluginResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachGroupPluginResponse) SetBody(v *DetachGroupPluginResponseBody) *DetachGroupPluginResponse {
	s.Body = v
	return s
}

func (s *DetachGroupPluginResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
