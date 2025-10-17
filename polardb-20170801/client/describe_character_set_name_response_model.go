// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCharacterSetNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCharacterSetNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCharacterSetNameResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCharacterSetNameResponseBody) *DescribeCharacterSetNameResponse
	GetBody() *DescribeCharacterSetNameResponseBody
}

type DescribeCharacterSetNameResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCharacterSetNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCharacterSetNameResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCharacterSetNameResponse) GoString() string {
	return s.String()
}

func (s *DescribeCharacterSetNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCharacterSetNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCharacterSetNameResponse) GetBody() *DescribeCharacterSetNameResponseBody {
	return s.Body
}

func (s *DescribeCharacterSetNameResponse) SetHeaders(v map[string]*string) *DescribeCharacterSetNameResponse {
	s.Headers = v
	return s
}

func (s *DescribeCharacterSetNameResponse) SetStatusCode(v int32) *DescribeCharacterSetNameResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCharacterSetNameResponse) SetBody(v *DescribeCharacterSetNameResponseBody) *DescribeCharacterSetNameResponse {
	s.Body = v
	return s
}

func (s *DescribeCharacterSetNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
