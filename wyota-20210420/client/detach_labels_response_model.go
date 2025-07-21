// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachLabelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachLabelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachLabelsResponse
	GetStatusCode() *int32
	SetBody(v *DetachLabelsResponseBody) *DetachLabelsResponse
	GetBody() *DetachLabelsResponseBody
}

type DetachLabelsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachLabelsResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachLabelsResponse) GoString() string {
	return s.String()
}

func (s *DetachLabelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachLabelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachLabelsResponse) GetBody() *DetachLabelsResponseBody {
	return s.Body
}

func (s *DetachLabelsResponse) SetHeaders(v map[string]*string) *DetachLabelsResponse {
	s.Headers = v
	return s
}

func (s *DetachLabelsResponse) SetStatusCode(v int32) *DetachLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachLabelsResponse) SetBody(v *DetachLabelsResponseBody) *DetachLabelsResponse {
	s.Body = v
	return s
}

func (s *DetachLabelsResponse) Validate() error {
	return dara.Validate(s)
}
