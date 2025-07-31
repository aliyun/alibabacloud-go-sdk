// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceByVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourceByVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourceByVersionResponse
	GetStatusCode() *int32
	SetBody(v *GetResourceByVersionResponseBody) *GetResourceByVersionResponse
	GetBody() *GetResourceByVersionResponseBody
}

type GetResourceByVersionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceByVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceByVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourceByVersionResponse) GoString() string {
	return s.String()
}

func (s *GetResourceByVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourceByVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourceByVersionResponse) GetBody() *GetResourceByVersionResponseBody {
	return s.Body
}

func (s *GetResourceByVersionResponse) SetHeaders(v map[string]*string) *GetResourceByVersionResponse {
	s.Headers = v
	return s
}

func (s *GetResourceByVersionResponse) SetStatusCode(v int32) *GetResourceByVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceByVersionResponse) SetBody(v *GetResourceByVersionResponseBody) *GetResourceByVersionResponse {
	s.Body = v
	return s
}

func (s *GetResourceByVersionResponse) Validate() error {
	return dara.Validate(s)
}
