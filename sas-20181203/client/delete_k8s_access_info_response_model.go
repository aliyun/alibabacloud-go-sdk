// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteK8sAccessInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteK8sAccessInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteK8sAccessInfoResponse
	GetStatusCode() *int32
	SetBody(v *DeleteK8sAccessInfoResponseBody) *DeleteK8sAccessInfoResponse
	GetBody() *DeleteK8sAccessInfoResponseBody
}

type DeleteK8sAccessInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteK8sAccessInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteK8sAccessInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteK8sAccessInfoResponse) GoString() string {
	return s.String()
}

func (s *DeleteK8sAccessInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteK8sAccessInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteK8sAccessInfoResponse) GetBody() *DeleteK8sAccessInfoResponseBody {
	return s.Body
}

func (s *DeleteK8sAccessInfoResponse) SetHeaders(v map[string]*string) *DeleteK8sAccessInfoResponse {
	s.Headers = v
	return s
}

func (s *DeleteK8sAccessInfoResponse) SetStatusCode(v int32) *DeleteK8sAccessInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteK8sAccessInfoResponse) SetBody(v *DeleteK8sAccessInfoResponseBody) *DeleteK8sAccessInfoResponse {
	s.Body = v
	return s
}

func (s *DeleteK8sAccessInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
