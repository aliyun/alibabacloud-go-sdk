// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCharacterSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCharacterSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCharacterSetResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCharacterSetResponseBody) *DescribeCharacterSetResponse
	GetBody() *DescribeCharacterSetResponseBody
}

type DescribeCharacterSetResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCharacterSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCharacterSetResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCharacterSetResponse) GoString() string {
	return s.String()
}

func (s *DescribeCharacterSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCharacterSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCharacterSetResponse) GetBody() *DescribeCharacterSetResponseBody {
	return s.Body
}

func (s *DescribeCharacterSetResponse) SetHeaders(v map[string]*string) *DescribeCharacterSetResponse {
	s.Headers = v
	return s
}

func (s *DescribeCharacterSetResponse) SetStatusCode(v int32) *DescribeCharacterSetResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCharacterSetResponse) SetBody(v *DescribeCharacterSetResponseBody) *DescribeCharacterSetResponse {
	s.Body = v
	return s
}

func (s *DescribeCharacterSetResponse) Validate() error {
	return dara.Validate(s)
}
