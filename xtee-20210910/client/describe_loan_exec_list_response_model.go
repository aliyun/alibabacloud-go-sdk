// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoanExecListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLoanExecListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLoanExecListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLoanExecListResponseBody) *DescribeLoanExecListResponse
	GetBody() *DescribeLoanExecListResponseBody
}

type DescribeLoanExecListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLoanExecListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLoanExecListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoanExecListResponse) GoString() string {
	return s.String()
}

func (s *DescribeLoanExecListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLoanExecListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLoanExecListResponse) GetBody() *DescribeLoanExecListResponseBody {
	return s.Body
}

func (s *DescribeLoanExecListResponse) SetHeaders(v map[string]*string) *DescribeLoanExecListResponse {
	s.Headers = v
	return s
}

func (s *DescribeLoanExecListResponse) SetStatusCode(v int32) *DescribeLoanExecListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLoanExecListResponse) SetBody(v *DescribeLoanExecListResponseBody) *DescribeLoanExecListResponse {
	s.Body = v
	return s
}

func (s *DescribeLoanExecListResponse) Validate() error {
	return dara.Validate(s)
}
