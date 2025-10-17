// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePerspectiveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePerspectiveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePerspectiveResponse
	GetStatusCode() *int32
	SetBody(v *DescribePerspectiveResponseBody) *DescribePerspectiveResponse
	GetBody() *DescribePerspectiveResponseBody
}

type DescribePerspectiveResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePerspectiveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePerspectiveResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePerspectiveResponse) GoString() string {
	return s.String()
}

func (s *DescribePerspectiveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePerspectiveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePerspectiveResponse) GetBody() *DescribePerspectiveResponseBody {
	return s.Body
}

func (s *DescribePerspectiveResponse) SetHeaders(v map[string]*string) *DescribePerspectiveResponse {
	s.Headers = v
	return s
}

func (s *DescribePerspectiveResponse) SetStatusCode(v int32) *DescribePerspectiveResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePerspectiveResponse) SetBody(v *DescribePerspectiveResponseBody) *DescribePerspectiveResponse {
	s.Body = v
	return s
}

func (s *DescribePerspectiveResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
