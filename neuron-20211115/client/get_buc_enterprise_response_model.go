// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucEnterpriseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucEnterpriseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucEnterpriseResponse
	GetStatusCode() *int32
	SetBody(v *BucEnterprise) *GetBucEnterpriseResponse
	GetBody() *BucEnterprise
}

type GetBucEnterpriseResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BucEnterprise     `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucEnterpriseResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucEnterpriseResponse) GoString() string {
	return s.String()
}

func (s *GetBucEnterpriseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucEnterpriseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucEnterpriseResponse) GetBody() *BucEnterprise {
	return s.Body
}

func (s *GetBucEnterpriseResponse) SetHeaders(v map[string]*string) *GetBucEnterpriseResponse {
	s.Headers = v
	return s
}

func (s *GetBucEnterpriseResponse) SetStatusCode(v int32) *GetBucEnterpriseResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucEnterpriseResponse) SetBody(v *BucEnterprise) *GetBucEnterpriseResponse {
	s.Body = v
	return s
}

func (s *GetBucEnterpriseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
