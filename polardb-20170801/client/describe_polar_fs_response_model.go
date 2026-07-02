// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarFsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePolarFsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePolarFsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePolarFsResponseBody) *DescribePolarFsResponse
	GetBody() *DescribePolarFsResponseBody
}

type DescribePolarFsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePolarFsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePolarFsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsResponse) GoString() string {
	return s.String()
}

func (s *DescribePolarFsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePolarFsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePolarFsResponse) GetBody() *DescribePolarFsResponseBody {
	return s.Body
}

func (s *DescribePolarFsResponse) SetHeaders(v map[string]*string) *DescribePolarFsResponse {
	s.Headers = v
	return s
}

func (s *DescribePolarFsResponse) SetStatusCode(v int32) *DescribePolarFsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolarFsResponse) SetBody(v *DescribePolarFsResponseBody) *DescribePolarFsResponse {
	s.Body = v
	return s
}

func (s *DescribePolarFsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
