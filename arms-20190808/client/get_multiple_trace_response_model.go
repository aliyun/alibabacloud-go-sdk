// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultipleTraceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMultipleTraceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMultipleTraceResponse
	GetStatusCode() *int32
	SetBody(v *GetMultipleTraceResponseBody) *GetMultipleTraceResponse
	GetBody() *GetMultipleTraceResponseBody
}

type GetMultipleTraceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMultipleTraceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMultipleTraceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMultipleTraceResponse) GoString() string {
	return s.String()
}

func (s *GetMultipleTraceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMultipleTraceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMultipleTraceResponse) GetBody() *GetMultipleTraceResponseBody {
	return s.Body
}

func (s *GetMultipleTraceResponse) SetHeaders(v map[string]*string) *GetMultipleTraceResponse {
	s.Headers = v
	return s
}

func (s *GetMultipleTraceResponse) SetStatusCode(v int32) *GetMultipleTraceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMultipleTraceResponse) SetBody(v *GetMultipleTraceResponseBody) *GetMultipleTraceResponse {
	s.Body = v
	return s
}

func (s *GetMultipleTraceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
