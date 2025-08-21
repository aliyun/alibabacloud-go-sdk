// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachInstanceSDGResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachInstanceSDGResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachInstanceSDGResponse
	GetStatusCode() *int32
	SetBody(v *AttachInstanceSDGResponseBody) *AttachInstanceSDGResponse
	GetBody() *AttachInstanceSDGResponseBody
}

type AttachInstanceSDGResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachInstanceSDGResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachInstanceSDGResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachInstanceSDGResponse) GoString() string {
	return s.String()
}

func (s *AttachInstanceSDGResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachInstanceSDGResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachInstanceSDGResponse) GetBody() *AttachInstanceSDGResponseBody {
	return s.Body
}

func (s *AttachInstanceSDGResponse) SetHeaders(v map[string]*string) *AttachInstanceSDGResponse {
	s.Headers = v
	return s
}

func (s *AttachInstanceSDGResponse) SetStatusCode(v int32) *AttachInstanceSDGResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachInstanceSDGResponse) SetBody(v *AttachInstanceSDGResponseBody) *AttachInstanceSDGResponse {
	s.Body = v
	return s
}

func (s *AttachInstanceSDGResponse) Validate() error {
	return dara.Validate(s)
}
