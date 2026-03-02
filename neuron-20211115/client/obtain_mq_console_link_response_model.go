// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObtainMqConsoleLinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ObtainMqConsoleLinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ObtainMqConsoleLinkResponse
	GetStatusCode() *int32
	SetBody(v *MqConsoleLinkResult) *ObtainMqConsoleLinkResponse
	GetBody() *MqConsoleLinkResult
}

type ObtainMqConsoleLinkResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MqConsoleLinkResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ObtainMqConsoleLinkResponse) String() string {
	return dara.Prettify(s)
}

func (s ObtainMqConsoleLinkResponse) GoString() string {
	return s.String()
}

func (s *ObtainMqConsoleLinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ObtainMqConsoleLinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ObtainMqConsoleLinkResponse) GetBody() *MqConsoleLinkResult {
	return s.Body
}

func (s *ObtainMqConsoleLinkResponse) SetHeaders(v map[string]*string) *ObtainMqConsoleLinkResponse {
	s.Headers = v
	return s
}

func (s *ObtainMqConsoleLinkResponse) SetStatusCode(v int32) *ObtainMqConsoleLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *ObtainMqConsoleLinkResponse) SetBody(v *MqConsoleLinkResult) *ObtainMqConsoleLinkResponse {
	s.Body = v
	return s
}

func (s *ObtainMqConsoleLinkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
