// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationPlansForRegionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOperationPlansForRegionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOperationPlansForRegionResponse
	GetStatusCode() *int32
	SetBody(v *ListOperationPlansForRegionResponseBody) *ListOperationPlansForRegionResponse
	GetBody() *ListOperationPlansForRegionResponseBody
}

type ListOperationPlansForRegionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOperationPlansForRegionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOperationPlansForRegionResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOperationPlansForRegionResponse) GoString() string {
	return s.String()
}

func (s *ListOperationPlansForRegionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOperationPlansForRegionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOperationPlansForRegionResponse) GetBody() *ListOperationPlansForRegionResponseBody {
	return s.Body
}

func (s *ListOperationPlansForRegionResponse) SetHeaders(v map[string]*string) *ListOperationPlansForRegionResponse {
	s.Headers = v
	return s
}

func (s *ListOperationPlansForRegionResponse) SetStatusCode(v int32) *ListOperationPlansForRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOperationPlansForRegionResponse) SetBody(v *ListOperationPlansForRegionResponseBody) *ListOperationPlansForRegionResponse {
	s.Body = v
	return s
}

func (s *ListOperationPlansForRegionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
