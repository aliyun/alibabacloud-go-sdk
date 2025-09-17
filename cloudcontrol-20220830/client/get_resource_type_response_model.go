// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourceTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourceTypeResponse
	GetStatusCode() *int32
	SetBody(v *GetResourceTypeResponseBody) *GetResourceTypeResponse
	GetBody() *GetResourceTypeResponseBody
}

type GetResourceTypeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourceTypeResponse) GoString() string {
	return s.String()
}

func (s *GetResourceTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourceTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourceTypeResponse) GetBody() *GetResourceTypeResponseBody {
	return s.Body
}

func (s *GetResourceTypeResponse) SetHeaders(v map[string]*string) *GetResourceTypeResponse {
	s.Headers = v
	return s
}

func (s *GetResourceTypeResponse) SetStatusCode(v int32) *GetResourceTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceTypeResponse) SetBody(v *GetResourceTypeResponseBody) *GetResourceTypeResponse {
	s.Body = v
	return s
}

func (s *GetResourceTypeResponse) Validate() error {
	return dara.Validate(s)
}
