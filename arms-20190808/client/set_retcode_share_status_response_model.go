// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetRetcodeShareStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetRetcodeShareStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetRetcodeShareStatusResponse
	GetStatusCode() *int32
	SetBody(v *SetRetcodeShareStatusResponseBody) *SetRetcodeShareStatusResponse
	GetBody() *SetRetcodeShareStatusResponseBody
}

type SetRetcodeShareStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetRetcodeShareStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetRetcodeShareStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s SetRetcodeShareStatusResponse) GoString() string {
	return s.String()
}

func (s *SetRetcodeShareStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetRetcodeShareStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetRetcodeShareStatusResponse) GetBody() *SetRetcodeShareStatusResponseBody {
	return s.Body
}

func (s *SetRetcodeShareStatusResponse) SetHeaders(v map[string]*string) *SetRetcodeShareStatusResponse {
	s.Headers = v
	return s
}

func (s *SetRetcodeShareStatusResponse) SetStatusCode(v int32) *SetRetcodeShareStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetRetcodeShareStatusResponse) SetBody(v *SetRetcodeShareStatusResponseBody) *SetRetcodeShareStatusResponse {
	s.Body = v
	return s
}

func (s *SetRetcodeShareStatusResponse) Validate() error {
	return dara.Validate(s)
}
