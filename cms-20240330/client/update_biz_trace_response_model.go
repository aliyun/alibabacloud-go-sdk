// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBizTraceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateBizTraceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateBizTraceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateBizTraceResponseBody) *UpdateBizTraceResponse
	GetBody() *UpdateBizTraceResponseBody
}

type UpdateBizTraceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBizTraceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBizTraceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateBizTraceResponse) GoString() string {
	return s.String()
}

func (s *UpdateBizTraceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateBizTraceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateBizTraceResponse) GetBody() *UpdateBizTraceResponseBody {
	return s.Body
}

func (s *UpdateBizTraceResponse) SetHeaders(v map[string]*string) *UpdateBizTraceResponse {
	s.Headers = v
	return s
}

func (s *UpdateBizTraceResponse) SetStatusCode(v int32) *UpdateBizTraceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBizTraceResponse) SetBody(v *UpdateBizTraceResponseBody) *UpdateBizTraceResponse {
	s.Body = v
	return s
}

func (s *UpdateBizTraceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
