// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProgressControlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ProgressControlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ProgressControlResponse
	GetStatusCode() *int32
	SetBody(v *ProgressControlResponseBody) *ProgressControlResponse
	GetBody() *ProgressControlResponseBody
}

type ProgressControlResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ProgressControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ProgressControlResponse) String() string {
	return dara.Prettify(s)
}

func (s ProgressControlResponse) GoString() string {
	return s.String()
}

func (s *ProgressControlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ProgressControlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ProgressControlResponse) GetBody() *ProgressControlResponseBody {
	return s.Body
}

func (s *ProgressControlResponse) SetHeaders(v map[string]*string) *ProgressControlResponse {
	s.Headers = v
	return s
}

func (s *ProgressControlResponse) SetStatusCode(v int32) *ProgressControlResponse {
	s.StatusCode = &v
	return s
}

func (s *ProgressControlResponse) SetBody(v *ProgressControlResponseBody) *ProgressControlResponse {
	s.Body = v
	return s
}

func (s *ProgressControlResponse) Validate() error {
	return dara.Validate(s)
}
