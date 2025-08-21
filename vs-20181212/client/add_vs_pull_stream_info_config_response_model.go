// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddVsPullStreamInfoConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddVsPullStreamInfoConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddVsPullStreamInfoConfigResponse
	GetStatusCode() *int32
	SetBody(v *AddVsPullStreamInfoConfigResponseBody) *AddVsPullStreamInfoConfigResponse
	GetBody() *AddVsPullStreamInfoConfigResponseBody
}

type AddVsPullStreamInfoConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddVsPullStreamInfoConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddVsPullStreamInfoConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s AddVsPullStreamInfoConfigResponse) GoString() string {
	return s.String()
}

func (s *AddVsPullStreamInfoConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddVsPullStreamInfoConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddVsPullStreamInfoConfigResponse) GetBody() *AddVsPullStreamInfoConfigResponseBody {
	return s.Body
}

func (s *AddVsPullStreamInfoConfigResponse) SetHeaders(v map[string]*string) *AddVsPullStreamInfoConfigResponse {
	s.Headers = v
	return s
}

func (s *AddVsPullStreamInfoConfigResponse) SetStatusCode(v int32) *AddVsPullStreamInfoConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *AddVsPullStreamInfoConfigResponse) SetBody(v *AddVsPullStreamInfoConfigResponseBody) *AddVsPullStreamInfoConfigResponse {
	s.Body = v
	return s
}

func (s *AddVsPullStreamInfoConfigResponse) Validate() error {
	return dara.Validate(s)
}
