// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRetcodeShareUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRetcodeShareUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRetcodeShareUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetRetcodeShareUrlResponseBody) *GetRetcodeShareUrlResponse
	GetBody() *GetRetcodeShareUrlResponseBody
}

type GetRetcodeShareUrlResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRetcodeShareUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRetcodeShareUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRetcodeShareUrlResponse) GoString() string {
	return s.String()
}

func (s *GetRetcodeShareUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRetcodeShareUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRetcodeShareUrlResponse) GetBody() *GetRetcodeShareUrlResponseBody {
	return s.Body
}

func (s *GetRetcodeShareUrlResponse) SetHeaders(v map[string]*string) *GetRetcodeShareUrlResponse {
	s.Headers = v
	return s
}

func (s *GetRetcodeShareUrlResponse) SetStatusCode(v int32) *GetRetcodeShareUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRetcodeShareUrlResponse) SetBody(v *GetRetcodeShareUrlResponseBody) *GetRetcodeShareUrlResponse {
	s.Body = v
	return s
}

func (s *GetRetcodeShareUrlResponse) Validate() error {
	return dara.Validate(s)
}
