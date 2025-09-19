// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAffectedMaliciousFileImagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAffectedMaliciousFileImagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAffectedMaliciousFileImagesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAffectedMaliciousFileImagesResponseBody) *DescribeAffectedMaliciousFileImagesResponse
	GetBody() *DescribeAffectedMaliciousFileImagesResponseBody
}

type DescribeAffectedMaliciousFileImagesResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAffectedMaliciousFileImagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAffectedMaliciousFileImagesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAffectedMaliciousFileImagesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAffectedMaliciousFileImagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAffectedMaliciousFileImagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAffectedMaliciousFileImagesResponse) GetBody() *DescribeAffectedMaliciousFileImagesResponseBody {
	return s.Body
}

func (s *DescribeAffectedMaliciousFileImagesResponse) SetHeaders(v map[string]*string) *DescribeAffectedMaliciousFileImagesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponse) SetStatusCode(v int32) *DescribeAffectedMaliciousFileImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponse) SetBody(v *DescribeAffectedMaliciousFileImagesResponseBody) *DescribeAffectedMaliciousFileImagesResponse {
	s.Body = v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponse) Validate() error {
	return dara.Validate(s)
}
