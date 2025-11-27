// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBizTraceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBizTraceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBizTraceResponse
	GetStatusCode() *int32
	SetBody(v *GetBizTraceResponseBody) *GetBizTraceResponse
	GetBody() *GetBizTraceResponseBody
}

type GetBizTraceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBizTraceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBizTraceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBizTraceResponse) GoString() string {
	return s.String()
}

func (s *GetBizTraceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBizTraceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBizTraceResponse) GetBody() *GetBizTraceResponseBody {
	return s.Body
}

func (s *GetBizTraceResponse) SetHeaders(v map[string]*string) *GetBizTraceResponse {
	s.Headers = v
	return s
}

func (s *GetBizTraceResponse) SetStatusCode(v int32) *GetBizTraceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBizTraceResponse) SetBody(v *GetBizTraceResponseBody) *GetBizTraceResponse {
	s.Body = v
	return s
}

func (s *GetBizTraceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
