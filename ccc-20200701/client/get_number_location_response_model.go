// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNumberLocationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNumberLocationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNumberLocationResponse
	GetStatusCode() *int32
	SetBody(v *GetNumberLocationResponseBody) *GetNumberLocationResponse
	GetBody() *GetNumberLocationResponseBody
}

type GetNumberLocationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNumberLocationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNumberLocationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNumberLocationResponse) GoString() string {
	return s.String()
}

func (s *GetNumberLocationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNumberLocationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNumberLocationResponse) GetBody() *GetNumberLocationResponseBody {
	return s.Body
}

func (s *GetNumberLocationResponse) SetHeaders(v map[string]*string) *GetNumberLocationResponse {
	s.Headers = v
	return s
}

func (s *GetNumberLocationResponse) SetStatusCode(v int32) *GetNumberLocationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNumberLocationResponse) SetBody(v *GetNumberLocationResponseBody) *GetNumberLocationResponse {
	s.Body = v
	return s
}

func (s *GetNumberLocationResponse) Validate() error {
	return dara.Validate(s)
}
