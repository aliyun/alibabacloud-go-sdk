// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachLabelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachLabelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachLabelsResponse
	GetStatusCode() *int32
	SetBody(v *AttachLabelsResponseBody) *AttachLabelsResponse
	GetBody() *AttachLabelsResponseBody
}

type AttachLabelsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachLabelsResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachLabelsResponse) GoString() string {
	return s.String()
}

func (s *AttachLabelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachLabelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachLabelsResponse) GetBody() *AttachLabelsResponseBody {
	return s.Body
}

func (s *AttachLabelsResponse) SetHeaders(v map[string]*string) *AttachLabelsResponse {
	s.Headers = v
	return s
}

func (s *AttachLabelsResponse) SetStatusCode(v int32) *AttachLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachLabelsResponse) SetBody(v *AttachLabelsResponseBody) *AttachLabelsResponse {
	s.Body = v
	return s
}

func (s *AttachLabelsResponse) Validate() error {
	return dara.Validate(s)
}
