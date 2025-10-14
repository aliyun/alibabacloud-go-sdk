// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSupportLinesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSupportLinesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSupportLinesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSupportLinesResponseBody) *DescribeSupportLinesResponse
	GetBody() *DescribeSupportLinesResponseBody
}

type DescribeSupportLinesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSupportLinesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSupportLinesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupportLinesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSupportLinesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSupportLinesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSupportLinesResponse) GetBody() *DescribeSupportLinesResponseBody {
	return s.Body
}

func (s *DescribeSupportLinesResponse) SetHeaders(v map[string]*string) *DescribeSupportLinesResponse {
	s.Headers = v
	return s
}

func (s *DescribeSupportLinesResponse) SetStatusCode(v int32) *DescribeSupportLinesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSupportLinesResponse) SetBody(v *DescribeSupportLinesResponseBody) *DescribeSupportLinesResponse {
	s.Body = v
	return s
}

func (s *DescribeSupportLinesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
