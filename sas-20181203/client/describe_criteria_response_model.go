// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCriteriaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCriteriaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCriteriaResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCriteriaResponseBody) *DescribeCriteriaResponse
	GetBody() *DescribeCriteriaResponseBody
}

type DescribeCriteriaResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCriteriaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCriteriaResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCriteriaResponse) GoString() string {
	return s.String()
}

func (s *DescribeCriteriaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCriteriaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCriteriaResponse) GetBody() *DescribeCriteriaResponseBody {
	return s.Body
}

func (s *DescribeCriteriaResponse) SetHeaders(v map[string]*string) *DescribeCriteriaResponse {
	s.Headers = v
	return s
}

func (s *DescribeCriteriaResponse) SetStatusCode(v int32) *DescribeCriteriaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCriteriaResponse) SetBody(v *DescribeCriteriaResponseBody) *DescribeCriteriaResponse {
	s.Body = v
	return s
}

func (s *DescribeCriteriaResponse) Validate() error {
	return dara.Validate(s)
}
