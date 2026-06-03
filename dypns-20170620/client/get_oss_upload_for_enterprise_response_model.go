// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssUploadForEnterpriseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOssUploadForEnterpriseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOssUploadForEnterpriseResponse
	GetStatusCode() *int32
	SetBody(v *GetOssUploadForEnterpriseResponseBody) *GetOssUploadForEnterpriseResponse
	GetBody() *GetOssUploadForEnterpriseResponseBody
}

type GetOssUploadForEnterpriseResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOssUploadForEnterpriseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOssUploadForEnterpriseResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOssUploadForEnterpriseResponse) GoString() string {
	return s.String()
}

func (s *GetOssUploadForEnterpriseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOssUploadForEnterpriseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOssUploadForEnterpriseResponse) GetBody() *GetOssUploadForEnterpriseResponseBody {
	return s.Body
}

func (s *GetOssUploadForEnterpriseResponse) SetHeaders(v map[string]*string) *GetOssUploadForEnterpriseResponse {
	s.Headers = v
	return s
}

func (s *GetOssUploadForEnterpriseResponse) SetStatusCode(v int32) *GetOssUploadForEnterpriseResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOssUploadForEnterpriseResponse) SetBody(v *GetOssUploadForEnterpriseResponseBody) *GetOssUploadForEnterpriseResponse {
	s.Body = v
	return s
}

func (s *GetOssUploadForEnterpriseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
