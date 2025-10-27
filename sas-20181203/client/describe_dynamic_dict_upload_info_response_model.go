// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDynamicDictUploadInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDynamicDictUploadInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDynamicDictUploadInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDynamicDictUploadInfoResponseBody) *DescribeDynamicDictUploadInfoResponse
	GetBody() *DescribeDynamicDictUploadInfoResponseBody
}

type DescribeDynamicDictUploadInfoResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDynamicDictUploadInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDynamicDictUploadInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDynamicDictUploadInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDynamicDictUploadInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDynamicDictUploadInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDynamicDictUploadInfoResponse) GetBody() *DescribeDynamicDictUploadInfoResponseBody {
	return s.Body
}

func (s *DescribeDynamicDictUploadInfoResponse) SetHeaders(v map[string]*string) *DescribeDynamicDictUploadInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDynamicDictUploadInfoResponse) SetStatusCode(v int32) *DescribeDynamicDictUploadInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDynamicDictUploadInfoResponse) SetBody(v *DescribeDynamicDictUploadInfoResponseBody) *DescribeDynamicDictUploadInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeDynamicDictUploadInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
