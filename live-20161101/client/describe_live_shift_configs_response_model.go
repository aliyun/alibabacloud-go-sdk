// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveShiftConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveShiftConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveShiftConfigsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveShiftConfigsResponseBody) *DescribeLiveShiftConfigsResponse
	GetBody() *DescribeLiveShiftConfigsResponseBody
}

type DescribeLiveShiftConfigsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveShiftConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveShiftConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveShiftConfigsResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveShiftConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveShiftConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveShiftConfigsResponse) GetBody() *DescribeLiveShiftConfigsResponseBody {
	return s.Body
}

func (s *DescribeLiveShiftConfigsResponse) SetHeaders(v map[string]*string) *DescribeLiveShiftConfigsResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveShiftConfigsResponse) SetStatusCode(v int32) *DescribeLiveShiftConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveShiftConfigsResponse) SetBody(v *DescribeLiveShiftConfigsResponseBody) *DescribeLiveShiftConfigsResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveShiftConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
