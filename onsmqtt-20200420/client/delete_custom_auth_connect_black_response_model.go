// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomAuthConnectBlackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomAuthConnectBlackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomAuthConnectBlackResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomAuthConnectBlackResponseBody) *DeleteCustomAuthConnectBlackResponse
	GetBody() *DeleteCustomAuthConnectBlackResponseBody
}

type DeleteCustomAuthConnectBlackResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomAuthConnectBlackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomAuthConnectBlackResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomAuthConnectBlackResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomAuthConnectBlackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomAuthConnectBlackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomAuthConnectBlackResponse) GetBody() *DeleteCustomAuthConnectBlackResponseBody {
	return s.Body
}

func (s *DeleteCustomAuthConnectBlackResponse) SetHeaders(v map[string]*string) *DeleteCustomAuthConnectBlackResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomAuthConnectBlackResponse) SetStatusCode(v int32) *DeleteCustomAuthConnectBlackResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomAuthConnectBlackResponse) SetBody(v *DeleteCustomAuthConnectBlackResponseBody) *DeleteCustomAuthConnectBlackResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomAuthConnectBlackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
