// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustVariablePageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCustVariablePageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCustVariablePageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCustVariablePageResponseBody) *DescribeCustVariablePageResponse
	GetBody() *DescribeCustVariablePageResponseBody
}

type DescribeCustVariablePageResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustVariablePageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustVariablePageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustVariablePageResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustVariablePageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCustVariablePageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCustVariablePageResponse) GetBody() *DescribeCustVariablePageResponseBody {
	return s.Body
}

func (s *DescribeCustVariablePageResponse) SetHeaders(v map[string]*string) *DescribeCustVariablePageResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustVariablePageResponse) SetStatusCode(v int32) *DescribeCustVariablePageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustVariablePageResponse) SetBody(v *DescribeCustVariablePageResponseBody) *DescribeCustVariablePageResponse {
	s.Body = v
	return s
}

func (s *DescribeCustVariablePageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
