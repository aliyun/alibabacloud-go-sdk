// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyAdviceByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyAdviceByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyAdviceByIdResponse
	GetStatusCode() *int32
	SetBody(v *ApplyAdviceByIdResponseBody) *ApplyAdviceByIdResponse
	GetBody() *ApplyAdviceByIdResponseBody
}

type ApplyAdviceByIdResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyAdviceByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyAdviceByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyAdviceByIdResponse) GoString() string {
	return s.String()
}

func (s *ApplyAdviceByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyAdviceByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyAdviceByIdResponse) GetBody() *ApplyAdviceByIdResponseBody {
	return s.Body
}

func (s *ApplyAdviceByIdResponse) SetHeaders(v map[string]*string) *ApplyAdviceByIdResponse {
	s.Headers = v
	return s
}

func (s *ApplyAdviceByIdResponse) SetStatusCode(v int32) *ApplyAdviceByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyAdviceByIdResponse) SetBody(v *ApplyAdviceByIdResponseBody) *ApplyAdviceByIdResponse {
	s.Body = v
	return s
}

func (s *ApplyAdviceByIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
