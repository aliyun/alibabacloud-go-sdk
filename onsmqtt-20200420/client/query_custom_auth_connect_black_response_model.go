// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCustomAuthConnectBlackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCustomAuthConnectBlackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCustomAuthConnectBlackResponse
	GetStatusCode() *int32
	SetBody(v *QueryCustomAuthConnectBlackResponseBody) *QueryCustomAuthConnectBlackResponse
	GetBody() *QueryCustomAuthConnectBlackResponseBody
}

type QueryCustomAuthConnectBlackResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCustomAuthConnectBlackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCustomAuthConnectBlackResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCustomAuthConnectBlackResponse) GoString() string {
	return s.String()
}

func (s *QueryCustomAuthConnectBlackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCustomAuthConnectBlackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCustomAuthConnectBlackResponse) GetBody() *QueryCustomAuthConnectBlackResponseBody {
	return s.Body
}

func (s *QueryCustomAuthConnectBlackResponse) SetHeaders(v map[string]*string) *QueryCustomAuthConnectBlackResponse {
	s.Headers = v
	return s
}

func (s *QueryCustomAuthConnectBlackResponse) SetStatusCode(v int32) *QueryCustomAuthConnectBlackResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCustomAuthConnectBlackResponse) SetBody(v *QueryCustomAuthConnectBlackResponseBody) *QueryCustomAuthConnectBlackResponse {
	s.Body = v
	return s
}

func (s *QueryCustomAuthConnectBlackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
