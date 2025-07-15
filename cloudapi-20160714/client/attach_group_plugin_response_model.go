// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachGroupPluginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachGroupPluginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachGroupPluginResponse
	GetStatusCode() *int32
	SetBody(v *AttachGroupPluginResponseBody) *AttachGroupPluginResponse
	GetBody() *AttachGroupPluginResponseBody
}

type AttachGroupPluginResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachGroupPluginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachGroupPluginResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachGroupPluginResponse) GoString() string {
	return s.String()
}

func (s *AttachGroupPluginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachGroupPluginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachGroupPluginResponse) GetBody() *AttachGroupPluginResponseBody {
	return s.Body
}

func (s *AttachGroupPluginResponse) SetHeaders(v map[string]*string) *AttachGroupPluginResponse {
	s.Headers = v
	return s
}

func (s *AttachGroupPluginResponse) SetStatusCode(v int32) *AttachGroupPluginResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachGroupPluginResponse) SetBody(v *AttachGroupPluginResponseBody) *AttachGroupPluginResponse {
	s.Body = v
	return s
}

func (s *AttachGroupPluginResponse) Validate() error {
	return dara.Validate(s)
}
