// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConsoleFuncGrayStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetConsoleFuncGrayStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetConsoleFuncGrayStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetConsoleFuncGrayStatusResponseBody) *GetConsoleFuncGrayStatusResponse
	GetBody() *GetConsoleFuncGrayStatusResponseBody
}

type GetConsoleFuncGrayStatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConsoleFuncGrayStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConsoleFuncGrayStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetConsoleFuncGrayStatusResponse) GoString() string {
	return s.String()
}

func (s *GetConsoleFuncGrayStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetConsoleFuncGrayStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetConsoleFuncGrayStatusResponse) GetBody() *GetConsoleFuncGrayStatusResponseBody {
	return s.Body
}

func (s *GetConsoleFuncGrayStatusResponse) SetHeaders(v map[string]*string) *GetConsoleFuncGrayStatusResponse {
	s.Headers = v
	return s
}

func (s *GetConsoleFuncGrayStatusResponse) SetStatusCode(v int32) *GetConsoleFuncGrayStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConsoleFuncGrayStatusResponse) SetBody(v *GetConsoleFuncGrayStatusResponseBody) *GetConsoleFuncGrayStatusResponse {
	s.Body = v
	return s
}

func (s *GetConsoleFuncGrayStatusResponse) Validate() error {
	return dara.Validate(s)
}
