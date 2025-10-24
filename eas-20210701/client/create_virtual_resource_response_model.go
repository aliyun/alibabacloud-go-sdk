// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVirtualResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVirtualResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVirtualResourceResponse
	GetStatusCode() *int32
	SetBody(v *CreateVirtualResourceResponseBody) *CreateVirtualResourceResponse
	GetBody() *CreateVirtualResourceResponseBody
}

type CreateVirtualResourceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVirtualResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVirtualResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVirtualResourceResponse) GoString() string {
	return s.String()
}

func (s *CreateVirtualResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVirtualResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVirtualResourceResponse) GetBody() *CreateVirtualResourceResponseBody {
	return s.Body
}

func (s *CreateVirtualResourceResponse) SetHeaders(v map[string]*string) *CreateVirtualResourceResponse {
	s.Headers = v
	return s
}

func (s *CreateVirtualResourceResponse) SetStatusCode(v int32) *CreateVirtualResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVirtualResourceResponse) SetBody(v *CreateVirtualResourceResponseBody) *CreateVirtualResourceResponse {
	s.Body = v
	return s
}

func (s *CreateVirtualResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
