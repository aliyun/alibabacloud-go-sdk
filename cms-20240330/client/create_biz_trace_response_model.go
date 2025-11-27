// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBizTraceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBizTraceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBizTraceResponse
	GetStatusCode() *int32
	SetBody(v *CreateBizTraceResponseBody) *CreateBizTraceResponse
	GetBody() *CreateBizTraceResponseBody
}

type CreateBizTraceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBizTraceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBizTraceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBizTraceResponse) GoString() string {
	return s.String()
}

func (s *CreateBizTraceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBizTraceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBizTraceResponse) GetBody() *CreateBizTraceResponseBody {
	return s.Body
}

func (s *CreateBizTraceResponse) SetHeaders(v map[string]*string) *CreateBizTraceResponse {
	s.Headers = v
	return s
}

func (s *CreateBizTraceResponse) SetStatusCode(v int32) *CreateBizTraceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBizTraceResponse) SetBody(v *CreateBizTraceResponseBody) *CreateBizTraceResponse {
	s.Body = v
	return s
}

func (s *CreateBizTraceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
