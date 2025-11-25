// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetComputeResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetComputeResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetComputeResourceResponse
	GetStatusCode() *int32
	SetBody(v *GetComputeResourceResponseBody) *GetComputeResourceResponse
	GetBody() *GetComputeResourceResponseBody
}

type GetComputeResourceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetComputeResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetComputeResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetComputeResourceResponse) GoString() string {
	return s.String()
}

func (s *GetComputeResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetComputeResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetComputeResourceResponse) GetBody() *GetComputeResourceResponseBody {
	return s.Body
}

func (s *GetComputeResourceResponse) SetHeaders(v map[string]*string) *GetComputeResourceResponse {
	s.Headers = v
	return s
}

func (s *GetComputeResourceResponse) SetStatusCode(v int32) *GetComputeResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetComputeResourceResponse) SetBody(v *GetComputeResourceResponseBody) *GetComputeResourceResponse {
	s.Body = v
	return s
}

func (s *GetComputeResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
