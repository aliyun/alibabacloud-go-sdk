// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSafTagListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSafTagListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSafTagListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSafTagListResponseBody) *DescribeSafTagListResponse
	GetBody() *DescribeSafTagListResponseBody
}

type DescribeSafTagListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSafTagListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSafTagListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSafTagListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSafTagListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSafTagListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSafTagListResponse) GetBody() *DescribeSafTagListResponseBody {
	return s.Body
}

func (s *DescribeSafTagListResponse) SetHeaders(v map[string]*string) *DescribeSafTagListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSafTagListResponse) SetStatusCode(v int32) *DescribeSafTagListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSafTagListResponse) SetBody(v *DescribeSafTagListResponseBody) *DescribeSafTagListResponse {
	s.Body = v
	return s
}

func (s *DescribeSafTagListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
