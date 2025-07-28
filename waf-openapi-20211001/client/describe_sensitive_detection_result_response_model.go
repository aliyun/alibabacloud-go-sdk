// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSensitiveDetectionResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSensitiveDetectionResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSensitiveDetectionResultResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSensitiveDetectionResultResponseBody) *DescribeSensitiveDetectionResultResponse
	GetBody() *DescribeSensitiveDetectionResultResponseBody
}

type DescribeSensitiveDetectionResultResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSensitiveDetectionResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSensitiveDetectionResultResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSensitiveDetectionResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeSensitiveDetectionResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSensitiveDetectionResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSensitiveDetectionResultResponse) GetBody() *DescribeSensitiveDetectionResultResponseBody {
	return s.Body
}

func (s *DescribeSensitiveDetectionResultResponse) SetHeaders(v map[string]*string) *DescribeSensitiveDetectionResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeSensitiveDetectionResultResponse) SetStatusCode(v int32) *DescribeSensitiveDetectionResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSensitiveDetectionResultResponse) SetBody(v *DescribeSensitiveDetectionResultResponseBody) *DescribeSensitiveDetectionResultResponse {
	s.Body = v
	return s
}

func (s *DescribeSensitiveDetectionResultResponse) Validate() error {
	return dara.Validate(s)
}
