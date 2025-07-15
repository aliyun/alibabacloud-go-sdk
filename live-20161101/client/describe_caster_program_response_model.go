// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCasterProgramResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCasterProgramResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCasterProgramResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCasterProgramResponseBody) *DescribeCasterProgramResponse
	GetBody() *DescribeCasterProgramResponseBody
}

type DescribeCasterProgramResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCasterProgramResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCasterProgramResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterProgramResponse) GoString() string {
	return s.String()
}

func (s *DescribeCasterProgramResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCasterProgramResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCasterProgramResponse) GetBody() *DescribeCasterProgramResponseBody {
	return s.Body
}

func (s *DescribeCasterProgramResponse) SetHeaders(v map[string]*string) *DescribeCasterProgramResponse {
	s.Headers = v
	return s
}

func (s *DescribeCasterProgramResponse) SetStatusCode(v int32) *DescribeCasterProgramResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCasterProgramResponse) SetBody(v *DescribeCasterProgramResponseBody) *DescribeCasterProgramResponse {
	s.Body = v
	return s
}

func (s *DescribeCasterProgramResponse) Validate() error {
	return dara.Validate(s)
}
