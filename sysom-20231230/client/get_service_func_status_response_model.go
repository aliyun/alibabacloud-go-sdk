// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceFuncStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetServiceFuncStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetServiceFuncStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetServiceFuncStatusResponseBody) *GetServiceFuncStatusResponse
	GetBody() *GetServiceFuncStatusResponseBody
}

type GetServiceFuncStatusResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceFuncStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceFuncStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetServiceFuncStatusResponse) GoString() string {
	return s.String()
}

func (s *GetServiceFuncStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetServiceFuncStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetServiceFuncStatusResponse) GetBody() *GetServiceFuncStatusResponseBody {
	return s.Body
}

func (s *GetServiceFuncStatusResponse) SetHeaders(v map[string]*string) *GetServiceFuncStatusResponse {
	s.Headers = v
	return s
}

func (s *GetServiceFuncStatusResponse) SetStatusCode(v int32) *GetServiceFuncStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceFuncStatusResponse) SetBody(v *GetServiceFuncStatusResponseBody) *GetServiceFuncStatusResponse {
	s.Body = v
	return s
}

func (s *GetServiceFuncStatusResponse) Validate() error {
	return dara.Validate(s)
}
