// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAddonTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAddonTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAddonTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAddonTemplateResponseBody) *DescribeAddonTemplateResponse
	GetBody() *DescribeAddonTemplateResponseBody
}

type DescribeAddonTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAddonTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAddonTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddonTemplateResponse) GoString() string {
	return s.String()
}

func (s *DescribeAddonTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAddonTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAddonTemplateResponse) GetBody() *DescribeAddonTemplateResponseBody {
	return s.Body
}

func (s *DescribeAddonTemplateResponse) SetHeaders(v map[string]*string) *DescribeAddonTemplateResponse {
	s.Headers = v
	return s
}

func (s *DescribeAddonTemplateResponse) SetStatusCode(v int32) *DescribeAddonTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAddonTemplateResponse) SetBody(v *DescribeAddonTemplateResponseBody) *DescribeAddonTemplateResponse {
	s.Body = v
	return s
}

func (s *DescribeAddonTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
