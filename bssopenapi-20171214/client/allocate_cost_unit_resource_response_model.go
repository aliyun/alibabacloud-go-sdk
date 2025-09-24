// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateCostUnitResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AllocateCostUnitResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AllocateCostUnitResourceResponse
	GetStatusCode() *int32
	SetBody(v *AllocateCostUnitResourceResponseBody) *AllocateCostUnitResourceResponse
	GetBody() *AllocateCostUnitResourceResponseBody
}

type AllocateCostUnitResourceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AllocateCostUnitResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AllocateCostUnitResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s AllocateCostUnitResourceResponse) GoString() string {
	return s.String()
}

func (s *AllocateCostUnitResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AllocateCostUnitResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AllocateCostUnitResourceResponse) GetBody() *AllocateCostUnitResourceResponseBody {
	return s.Body
}

func (s *AllocateCostUnitResourceResponse) SetHeaders(v map[string]*string) *AllocateCostUnitResourceResponse {
	s.Headers = v
	return s
}

func (s *AllocateCostUnitResourceResponse) SetStatusCode(v int32) *AllocateCostUnitResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *AllocateCostUnitResourceResponse) SetBody(v *AllocateCostUnitResourceResponseBody) *AllocateCostUnitResourceResponse {
	s.Body = v
	return s
}

func (s *AllocateCostUnitResourceResponse) Validate() error {
	return dara.Validate(s)
}
