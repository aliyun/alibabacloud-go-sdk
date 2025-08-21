// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSelfImagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSelfImagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSelfImagesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSelfImagesResponseBody) *DescribeSelfImagesResponse
	GetBody() *DescribeSelfImagesResponseBody
}

type DescribeSelfImagesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSelfImagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSelfImagesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSelfImagesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSelfImagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSelfImagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSelfImagesResponse) GetBody() *DescribeSelfImagesResponseBody {
	return s.Body
}

func (s *DescribeSelfImagesResponse) SetHeaders(v map[string]*string) *DescribeSelfImagesResponse {
	s.Headers = v
	return s
}

func (s *DescribeSelfImagesResponse) SetStatusCode(v int32) *DescribeSelfImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSelfImagesResponse) SetBody(v *DescribeSelfImagesResponseBody) *DescribeSelfImagesResponse {
	s.Body = v
	return s
}

func (s *DescribeSelfImagesResponse) Validate() error {
	return dara.Validate(s)
}
