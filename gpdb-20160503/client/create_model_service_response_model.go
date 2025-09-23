// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateModelServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateModelServiceResponse
	GetStatusCode() *int32
	SetBody(v *CreateModelServiceResponseBody) *CreateModelServiceResponse
	GetBody() *CreateModelServiceResponseBody
}

type CreateModelServiceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateModelServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateModelServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateModelServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateModelServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateModelServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateModelServiceResponse) GetBody() *CreateModelServiceResponseBody {
	return s.Body
}

func (s *CreateModelServiceResponse) SetHeaders(v map[string]*string) *CreateModelServiceResponse {
	s.Headers = v
	return s
}

func (s *CreateModelServiceResponse) SetStatusCode(v int32) *CreateModelServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateModelServiceResponse) SetBody(v *CreateModelServiceResponseBody) *CreateModelServiceResponse {
	s.Body = v
	return s
}

func (s *CreateModelServiceResponse) Validate() error {
	return dara.Validate(s)
}
