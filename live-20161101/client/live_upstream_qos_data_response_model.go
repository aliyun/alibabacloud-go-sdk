// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLiveUpstreamQosDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LiveUpstreamQosDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LiveUpstreamQosDataResponse
	GetStatusCode() *int32
	SetBody(v *LiveUpstreamQosDataResponseBody) *LiveUpstreamQosDataResponse
	GetBody() *LiveUpstreamQosDataResponseBody
}

type LiveUpstreamQosDataResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LiveUpstreamQosDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LiveUpstreamQosDataResponse) String() string {
	return dara.Prettify(s)
}

func (s LiveUpstreamQosDataResponse) GoString() string {
	return s.String()
}

func (s *LiveUpstreamQosDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LiveUpstreamQosDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LiveUpstreamQosDataResponse) GetBody() *LiveUpstreamQosDataResponseBody {
	return s.Body
}

func (s *LiveUpstreamQosDataResponse) SetHeaders(v map[string]*string) *LiveUpstreamQosDataResponse {
	s.Headers = v
	return s
}

func (s *LiveUpstreamQosDataResponse) SetStatusCode(v int32) *LiveUpstreamQosDataResponse {
	s.StatusCode = &v
	return s
}

func (s *LiveUpstreamQosDataResponse) SetBody(v *LiveUpstreamQosDataResponseBody) *LiveUpstreamQosDataResponse {
	s.Body = v
	return s
}

func (s *LiveUpstreamQosDataResponse) Validate() error {
	return dara.Validate(s)
}
