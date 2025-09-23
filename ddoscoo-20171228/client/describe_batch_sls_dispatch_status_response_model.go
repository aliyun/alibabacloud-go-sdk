// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBatchSlsDispatchStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBatchSlsDispatchStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBatchSlsDispatchStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBatchSlsDispatchStatusResponseBody) *DescribeBatchSlsDispatchStatusResponse
	GetBody() *DescribeBatchSlsDispatchStatusResponseBody
}

type DescribeBatchSlsDispatchStatusResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBatchSlsDispatchStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBatchSlsDispatchStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBatchSlsDispatchStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeBatchSlsDispatchStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBatchSlsDispatchStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBatchSlsDispatchStatusResponse) GetBody() *DescribeBatchSlsDispatchStatusResponseBody {
	return s.Body
}

func (s *DescribeBatchSlsDispatchStatusResponse) SetHeaders(v map[string]*string) *DescribeBatchSlsDispatchStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeBatchSlsDispatchStatusResponse) SetStatusCode(v int32) *DescribeBatchSlsDispatchStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusResponse) SetBody(v *DescribeBatchSlsDispatchStatusResponseBody) *DescribeBatchSlsDispatchStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeBatchSlsDispatchStatusResponse) Validate() error {
	return dara.Validate(s)
}
