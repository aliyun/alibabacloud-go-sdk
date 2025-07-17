// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRetcodeAppByPidResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRetcodeAppByPidResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRetcodeAppByPidResponse
	GetStatusCode() *int32
	SetBody(v *GetRetcodeAppByPidResponseBody) *GetRetcodeAppByPidResponse
	GetBody() *GetRetcodeAppByPidResponseBody
}

type GetRetcodeAppByPidResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRetcodeAppByPidResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRetcodeAppByPidResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRetcodeAppByPidResponse) GoString() string {
	return s.String()
}

func (s *GetRetcodeAppByPidResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRetcodeAppByPidResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRetcodeAppByPidResponse) GetBody() *GetRetcodeAppByPidResponseBody {
	return s.Body
}

func (s *GetRetcodeAppByPidResponse) SetHeaders(v map[string]*string) *GetRetcodeAppByPidResponse {
	s.Headers = v
	return s
}

func (s *GetRetcodeAppByPidResponse) SetStatusCode(v int32) *GetRetcodeAppByPidResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRetcodeAppByPidResponse) SetBody(v *GetRetcodeAppByPidResponseBody) *GetRetcodeAppByPidResponse {
	s.Body = v
	return s
}

func (s *GetRetcodeAppByPidResponse) Validate() error {
	return dara.Validate(s)
}
