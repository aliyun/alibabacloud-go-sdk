// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomAuthConnectBlackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddCustomAuthConnectBlackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddCustomAuthConnectBlackResponse
	GetStatusCode() *int32
	SetBody(v *AddCustomAuthConnectBlackResponseBody) *AddCustomAuthConnectBlackResponse
	GetBody() *AddCustomAuthConnectBlackResponseBody
}

type AddCustomAuthConnectBlackResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCustomAuthConnectBlackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCustomAuthConnectBlackResponse) String() string {
	return dara.Prettify(s)
}

func (s AddCustomAuthConnectBlackResponse) GoString() string {
	return s.String()
}

func (s *AddCustomAuthConnectBlackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddCustomAuthConnectBlackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddCustomAuthConnectBlackResponse) GetBody() *AddCustomAuthConnectBlackResponseBody {
	return s.Body
}

func (s *AddCustomAuthConnectBlackResponse) SetHeaders(v map[string]*string) *AddCustomAuthConnectBlackResponse {
	s.Headers = v
	return s
}

func (s *AddCustomAuthConnectBlackResponse) SetStatusCode(v int32) *AddCustomAuthConnectBlackResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCustomAuthConnectBlackResponse) SetBody(v *AddCustomAuthConnectBlackResponseBody) *AddCustomAuthConnectBlackResponse {
	s.Body = v
	return s
}

func (s *AddCustomAuthConnectBlackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
