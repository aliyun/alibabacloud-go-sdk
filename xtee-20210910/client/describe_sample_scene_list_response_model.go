// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSampleSceneListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSampleSceneListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSampleSceneListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSampleSceneListResponseBody) *DescribeSampleSceneListResponse
	GetBody() *DescribeSampleSceneListResponseBody
}

type DescribeSampleSceneListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSampleSceneListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSampleSceneListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleSceneListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSampleSceneListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSampleSceneListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSampleSceneListResponse) GetBody() *DescribeSampleSceneListResponseBody {
	return s.Body
}

func (s *DescribeSampleSceneListResponse) SetHeaders(v map[string]*string) *DescribeSampleSceneListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSampleSceneListResponse) SetStatusCode(v int32) *DescribeSampleSceneListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSampleSceneListResponse) SetBody(v *DescribeSampleSceneListResponseBody) *DescribeSampleSceneListResponse {
	s.Body = v
	return s
}

func (s *DescribeSampleSceneListResponse) Validate() error {
	return dara.Validate(s)
}
