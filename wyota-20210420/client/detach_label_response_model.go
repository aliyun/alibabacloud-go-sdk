// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachLabelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachLabelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachLabelResponse
	GetStatusCode() *int32
	SetBody(v *DetachLabelResponseBody) *DetachLabelResponse
	GetBody() *DetachLabelResponseBody
}

type DetachLabelResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachLabelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachLabelResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachLabelResponse) GoString() string {
	return s.String()
}

func (s *DetachLabelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachLabelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachLabelResponse) GetBody() *DetachLabelResponseBody {
	return s.Body
}

func (s *DetachLabelResponse) SetHeaders(v map[string]*string) *DetachLabelResponse {
	s.Headers = v
	return s
}

func (s *DetachLabelResponse) SetStatusCode(v int32) *DetachLabelResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachLabelResponse) SetBody(v *DetachLabelResponseBody) *DetachLabelResponse {
	s.Body = v
	return s
}

func (s *DetachLabelResponse) Validate() error {
	return dara.Validate(s)
}
