// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhysicalInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPhysicalInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPhysicalInstanceResponse
	GetStatusCode() *int32
	SetBody(v *GetPhysicalInstanceResponseBody) *GetPhysicalInstanceResponse
	GetBody() *GetPhysicalInstanceResponseBody
}

type GetPhysicalInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPhysicalInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPhysicalInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetPhysicalInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPhysicalInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPhysicalInstanceResponse) GetBody() *GetPhysicalInstanceResponseBody {
	return s.Body
}

func (s *GetPhysicalInstanceResponse) SetHeaders(v map[string]*string) *GetPhysicalInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetPhysicalInstanceResponse) SetStatusCode(v int32) *GetPhysicalInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPhysicalInstanceResponse) SetBody(v *GetPhysicalInstanceResponseBody) *GetPhysicalInstanceResponse {
	s.Body = v
	return s
}

func (s *GetPhysicalInstanceResponse) Validate() error {
	return dara.Validate(s)
}
