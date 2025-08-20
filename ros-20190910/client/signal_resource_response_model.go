// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSignalResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SignalResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SignalResourceResponse
	GetStatusCode() *int32
	SetBody(v *SignalResourceResponseBody) *SignalResourceResponse
	GetBody() *SignalResourceResponseBody
}

type SignalResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SignalResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SignalResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s SignalResourceResponse) GoString() string {
	return s.String()
}

func (s *SignalResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SignalResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SignalResourceResponse) GetBody() *SignalResourceResponseBody {
	return s.Body
}

func (s *SignalResourceResponse) SetHeaders(v map[string]*string) *SignalResourceResponse {
	s.Headers = v
	return s
}

func (s *SignalResourceResponse) SetStatusCode(v int32) *SignalResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *SignalResourceResponse) SetBody(v *SignalResourceResponseBody) *SignalResourceResponse {
	s.Body = v
	return s
}

func (s *SignalResourceResponse) Validate() error {
	return dara.Validate(s)
}
