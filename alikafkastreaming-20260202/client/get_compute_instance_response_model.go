// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetComputeInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetComputeInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetComputeInstanceResponse
	GetStatusCode() *int32
	SetBody(v *GetComputeInstanceResponseBody) *GetComputeInstanceResponse
	GetBody() *GetComputeInstanceResponseBody
}

type GetComputeInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetComputeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetComputeInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetComputeInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetComputeInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetComputeInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetComputeInstanceResponse) GetBody() *GetComputeInstanceResponseBody {
	return s.Body
}

func (s *GetComputeInstanceResponse) SetHeaders(v map[string]*string) *GetComputeInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetComputeInstanceResponse) SetStatusCode(v int32) *GetComputeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetComputeInstanceResponse) SetBody(v *GetComputeInstanceResponseBody) *GetComputeInstanceResponse {
	s.Body = v
	return s
}

func (s *GetComputeInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
