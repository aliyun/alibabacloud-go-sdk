// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachLabelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachLabelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachLabelResponse
	GetStatusCode() *int32
	SetBody(v *AttachLabelResponseBody) *AttachLabelResponse
	GetBody() *AttachLabelResponseBody
}

type AttachLabelResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachLabelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachLabelResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachLabelResponse) GoString() string {
	return s.String()
}

func (s *AttachLabelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachLabelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachLabelResponse) GetBody() *AttachLabelResponseBody {
	return s.Body
}

func (s *AttachLabelResponse) SetHeaders(v map[string]*string) *AttachLabelResponse {
	s.Headers = v
	return s
}

func (s *AttachLabelResponse) SetStatusCode(v int32) *AttachLabelResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachLabelResponse) SetBody(v *AttachLabelResponseBody) *AttachLabelResponse {
	s.Body = v
	return s
}

func (s *AttachLabelResponse) Validate() error {
	return dara.Validate(s)
}
