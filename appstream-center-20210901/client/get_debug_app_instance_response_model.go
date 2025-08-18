// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDebugAppInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDebugAppInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDebugAppInstanceResponse
	GetStatusCode() *int32
	SetBody(v *GetDebugAppInstanceResponseBody) *GetDebugAppInstanceResponse
	GetBody() *GetDebugAppInstanceResponseBody
}

type GetDebugAppInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDebugAppInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDebugAppInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDebugAppInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetDebugAppInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDebugAppInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDebugAppInstanceResponse) GetBody() *GetDebugAppInstanceResponseBody {
	return s.Body
}

func (s *GetDebugAppInstanceResponse) SetHeaders(v map[string]*string) *GetDebugAppInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetDebugAppInstanceResponse) SetStatusCode(v int32) *GetDebugAppInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDebugAppInstanceResponse) SetBody(v *GetDebugAppInstanceResponseBody) *GetDebugAppInstanceResponse {
	s.Body = v
	return s
}

func (s *GetDebugAppInstanceResponse) Validate() error {
	return dara.Validate(s)
}
