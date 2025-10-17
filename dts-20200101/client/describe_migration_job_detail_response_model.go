// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMigrationJobDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMigrationJobDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMigrationJobDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMigrationJobDetailResponseBody) *DescribeMigrationJobDetailResponse
	GetBody() *DescribeMigrationJobDetailResponseBody
}

type DescribeMigrationJobDetailResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMigrationJobDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMigrationJobDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMigrationJobDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMigrationJobDetailResponse) GetBody() *DescribeMigrationJobDetailResponseBody {
	return s.Body
}

func (s *DescribeMigrationJobDetailResponse) SetHeaders(v map[string]*string) *DescribeMigrationJobDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeMigrationJobDetailResponse) SetStatusCode(v int32) *DescribeMigrationJobDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMigrationJobDetailResponse) SetBody(v *DescribeMigrationJobDetailResponseBody) *DescribeMigrationJobDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeMigrationJobDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
