// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateCostCenterResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AllocateCostCenterResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AllocateCostCenterResourceResponse
	GetStatusCode() *int32
	SetBody(v *AllocateCostCenterResourceResponseBody) *AllocateCostCenterResourceResponse
	GetBody() *AllocateCostCenterResourceResponseBody
}

type AllocateCostCenterResourceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AllocateCostCenterResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AllocateCostCenterResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s AllocateCostCenterResourceResponse) GoString() string {
	return s.String()
}

func (s *AllocateCostCenterResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AllocateCostCenterResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AllocateCostCenterResourceResponse) GetBody() *AllocateCostCenterResourceResponseBody {
	return s.Body
}

func (s *AllocateCostCenterResourceResponse) SetHeaders(v map[string]*string) *AllocateCostCenterResourceResponse {
	s.Headers = v
	return s
}

func (s *AllocateCostCenterResourceResponse) SetStatusCode(v int32) *AllocateCostCenterResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *AllocateCostCenterResourceResponse) SetBody(v *AllocateCostCenterResourceResponseBody) *AllocateCostCenterResourceResponse {
	s.Body = v
	return s
}

func (s *AllocateCostCenterResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
