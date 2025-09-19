// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFrontVulPatchListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFrontVulPatchListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFrontVulPatchListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFrontVulPatchListResponseBody) *DescribeFrontVulPatchListResponse
	GetBody() *DescribeFrontVulPatchListResponseBody
}

type DescribeFrontVulPatchListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFrontVulPatchListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFrontVulPatchListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFrontVulPatchListResponse) GoString() string {
	return s.String()
}

func (s *DescribeFrontVulPatchListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFrontVulPatchListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFrontVulPatchListResponse) GetBody() *DescribeFrontVulPatchListResponseBody {
	return s.Body
}

func (s *DescribeFrontVulPatchListResponse) SetHeaders(v map[string]*string) *DescribeFrontVulPatchListResponse {
	s.Headers = v
	return s
}

func (s *DescribeFrontVulPatchListResponse) SetStatusCode(v int32) *DescribeFrontVulPatchListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFrontVulPatchListResponse) SetBody(v *DescribeFrontVulPatchListResponseBody) *DescribeFrontVulPatchListResponse {
	s.Body = v
	return s
}

func (s *DescribeFrontVulPatchListResponse) Validate() error {
	return dara.Validate(s)
}
