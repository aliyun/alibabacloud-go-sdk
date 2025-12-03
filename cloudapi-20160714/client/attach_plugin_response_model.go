// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachPluginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachPluginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachPluginResponse
	GetStatusCode() *int32
	SetBody(v *AttachPluginResponseBody) *AttachPluginResponse
	GetBody() *AttachPluginResponseBody
}

type AttachPluginResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachPluginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachPluginResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachPluginResponse) GoString() string {
	return s.String()
}

func (s *AttachPluginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachPluginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachPluginResponse) GetBody() *AttachPluginResponseBody {
	return s.Body
}

func (s *AttachPluginResponse) SetHeaders(v map[string]*string) *AttachPluginResponse {
	s.Headers = v
	return s
}

func (s *AttachPluginResponse) SetStatusCode(v int32) *AttachPluginResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachPluginResponse) SetBody(v *AttachPluginResponseBody) *AttachPluginResponse {
	s.Body = v
	return s
}

func (s *AttachPluginResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
