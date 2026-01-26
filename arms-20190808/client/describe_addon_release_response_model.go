// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAddonReleaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAddonReleaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAddonReleaseResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAddonReleaseResponseBody) *DescribeAddonReleaseResponse
	GetBody() *DescribeAddonReleaseResponseBody
}

type DescribeAddonReleaseResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAddonReleaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAddonReleaseResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddonReleaseResponse) GoString() string {
	return s.String()
}

func (s *DescribeAddonReleaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAddonReleaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAddonReleaseResponse) GetBody() *DescribeAddonReleaseResponseBody {
	return s.Body
}

func (s *DescribeAddonReleaseResponse) SetHeaders(v map[string]*string) *DescribeAddonReleaseResponse {
	s.Headers = v
	return s
}

func (s *DescribeAddonReleaseResponse) SetStatusCode(v int32) *DescribeAddonReleaseResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAddonReleaseResponse) SetBody(v *DescribeAddonReleaseResponseBody) *DescribeAddonReleaseResponse {
	s.Body = v
	return s
}

func (s *DescribeAddonReleaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
