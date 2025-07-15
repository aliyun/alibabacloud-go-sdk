// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveAIStudioResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveAIStudioResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveAIStudioResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveAIStudioResponseBody) *DescribeLiveAIStudioResponse
	GetBody() *DescribeLiveAIStudioResponseBody
}

type DescribeLiveAIStudioResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveAIStudioResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveAIStudioResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveAIStudioResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveAIStudioResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveAIStudioResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveAIStudioResponse) GetBody() *DescribeLiveAIStudioResponseBody {
	return s.Body
}

func (s *DescribeLiveAIStudioResponse) SetHeaders(v map[string]*string) *DescribeLiveAIStudioResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveAIStudioResponse) SetStatusCode(v int32) *DescribeLiveAIStudioResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveAIStudioResponse) SetBody(v *DescribeLiveAIStudioResponseBody) *DescribeLiveAIStudioResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveAIStudioResponse) Validate() error {
	return dara.Validate(s)
}
