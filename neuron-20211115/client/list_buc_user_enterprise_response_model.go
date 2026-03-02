// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBucUserEnterpriseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBucUserEnterpriseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBucUserEnterpriseResponse
	GetStatusCode() *int32
	SetBody(v *BucUserEnterpriseListResult) *ListBucUserEnterpriseResponse
	GetBody() *BucUserEnterpriseListResult
}

type ListBucUserEnterpriseResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BucUserEnterpriseListResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBucUserEnterpriseResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBucUserEnterpriseResponse) GoString() string {
	return s.String()
}

func (s *ListBucUserEnterpriseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBucUserEnterpriseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBucUserEnterpriseResponse) GetBody() *BucUserEnterpriseListResult {
	return s.Body
}

func (s *ListBucUserEnterpriseResponse) SetHeaders(v map[string]*string) *ListBucUserEnterpriseResponse {
	s.Headers = v
	return s
}

func (s *ListBucUserEnterpriseResponse) SetStatusCode(v int32) *ListBucUserEnterpriseResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBucUserEnterpriseResponse) SetBody(v *BucUserEnterpriseListResult) *ListBucUserEnterpriseResponse {
	s.Body = v
	return s
}

func (s *ListBucUserEnterpriseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
