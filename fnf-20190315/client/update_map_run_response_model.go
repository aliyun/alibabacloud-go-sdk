// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMapRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMapRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMapRunResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMapRunResponseBody) *UpdateMapRunResponse
	GetBody() *UpdateMapRunResponseBody
}

type UpdateMapRunResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMapRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMapRunResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMapRunResponse) GoString() string {
	return s.String()
}

func (s *UpdateMapRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMapRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMapRunResponse) GetBody() *UpdateMapRunResponseBody {
	return s.Body
}

func (s *UpdateMapRunResponse) SetHeaders(v map[string]*string) *UpdateMapRunResponse {
	s.Headers = v
	return s
}

func (s *UpdateMapRunResponse) SetStatusCode(v int32) *UpdateMapRunResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMapRunResponse) SetBody(v *UpdateMapRunResponseBody) *UpdateMapRunResponse {
	s.Body = v
	return s
}

func (s *UpdateMapRunResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
