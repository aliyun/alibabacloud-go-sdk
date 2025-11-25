// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComputeResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateComputeResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateComputeResourceResponse
	GetStatusCode() *int32
	SetBody(v *CreateComputeResourceResponseBody) *CreateComputeResourceResponse
	GetBody() *CreateComputeResourceResponseBody
}

type CreateComputeResourceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateComputeResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateComputeResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeResourceResponse) GoString() string {
	return s.String()
}

func (s *CreateComputeResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateComputeResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateComputeResourceResponse) GetBody() *CreateComputeResourceResponseBody {
	return s.Body
}

func (s *CreateComputeResourceResponse) SetHeaders(v map[string]*string) *CreateComputeResourceResponse {
	s.Headers = v
	return s
}

func (s *CreateComputeResourceResponse) SetStatusCode(v int32) *CreateComputeResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateComputeResourceResponse) SetBody(v *CreateComputeResourceResponseBody) *CreateComputeResourceResponse {
	s.Body = v
	return s
}

func (s *CreateComputeResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
