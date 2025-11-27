// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBizTraceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBizTraceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBizTraceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBizTraceResponseBody) *DeleteBizTraceResponse
	GetBody() *DeleteBizTraceResponseBody
}

type DeleteBizTraceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBizTraceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBizTraceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBizTraceResponse) GoString() string {
	return s.String()
}

func (s *DeleteBizTraceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBizTraceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBizTraceResponse) GetBody() *DeleteBizTraceResponseBody {
	return s.Body
}

func (s *DeleteBizTraceResponse) SetHeaders(v map[string]*string) *DeleteBizTraceResponse {
	s.Headers = v
	return s
}

func (s *DeleteBizTraceResponse) SetStatusCode(v int32) *DeleteBizTraceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBizTraceResponse) SetBody(v *DeleteBizTraceResponseBody) *DeleteBizTraceResponse {
	s.Body = v
	return s
}

func (s *DeleteBizTraceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
