// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUdfByVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUdfByVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUdfByVersionResponse
	GetStatusCode() *int32
	SetBody(v *GetUdfByVersionResponseBody) *GetUdfByVersionResponse
	GetBody() *GetUdfByVersionResponseBody
}

type GetUdfByVersionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUdfByVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUdfByVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUdfByVersionResponse) GoString() string {
	return s.String()
}

func (s *GetUdfByVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUdfByVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUdfByVersionResponse) GetBody() *GetUdfByVersionResponseBody {
	return s.Body
}

func (s *GetUdfByVersionResponse) SetHeaders(v map[string]*string) *GetUdfByVersionResponse {
	s.Headers = v
	return s
}

func (s *GetUdfByVersionResponse) SetStatusCode(v int32) *GetUdfByVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUdfByVersionResponse) SetBody(v *GetUdfByVersionResponseBody) *GetUdfByVersionResponse {
	s.Body = v
	return s
}

func (s *GetUdfByVersionResponse) Validate() error {
	return dara.Validate(s)
}
