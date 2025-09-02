// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSuccessInstanceTrendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSuccessInstanceTrendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSuccessInstanceTrendResponse
	GetStatusCode() *int32
	SetBody(v *GetSuccessInstanceTrendResponseBody) *GetSuccessInstanceTrendResponse
	GetBody() *GetSuccessInstanceTrendResponseBody
}

type GetSuccessInstanceTrendResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSuccessInstanceTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSuccessInstanceTrendResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSuccessInstanceTrendResponse) GoString() string {
	return s.String()
}

func (s *GetSuccessInstanceTrendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSuccessInstanceTrendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSuccessInstanceTrendResponse) GetBody() *GetSuccessInstanceTrendResponseBody {
	return s.Body
}

func (s *GetSuccessInstanceTrendResponse) SetHeaders(v map[string]*string) *GetSuccessInstanceTrendResponse {
	s.Headers = v
	return s
}

func (s *GetSuccessInstanceTrendResponse) SetStatusCode(v int32) *GetSuccessInstanceTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSuccessInstanceTrendResponse) SetBody(v *GetSuccessInstanceTrendResponseBody) *GetSuccessInstanceTrendResponse {
	s.Body = v
	return s
}

func (s *GetSuccessInstanceTrendResponse) Validate() error {
	return dara.Validate(s)
}
