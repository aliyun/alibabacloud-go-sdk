// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDetectPornDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveDetectPornDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveDetectPornDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveDetectPornDataResponseBody) *DescribeLiveDetectPornDataResponse
	GetBody() *DescribeLiveDetectPornDataResponseBody
}

type DescribeLiveDetectPornDataResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveDetectPornDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveDetectPornDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDetectPornDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDetectPornDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveDetectPornDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveDetectPornDataResponse) GetBody() *DescribeLiveDetectPornDataResponseBody {
	return s.Body
}

func (s *DescribeLiveDetectPornDataResponse) SetHeaders(v map[string]*string) *DescribeLiveDetectPornDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveDetectPornDataResponse) SetStatusCode(v int32) *DescribeLiveDetectPornDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveDetectPornDataResponse) SetBody(v *DescribeLiveDetectPornDataResponseBody) *DescribeLiveDetectPornDataResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveDetectPornDataResponse) Validate() error {
	return dara.Validate(s)
}
