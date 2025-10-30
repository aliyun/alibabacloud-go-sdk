// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCloudResourceCertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCloudResourceCertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCloudResourceCertResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCloudResourceCertResponseBody) *ModifyCloudResourceCertResponse
	GetBody() *ModifyCloudResourceCertResponseBody
}

type ModifyCloudResourceCertResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCloudResourceCertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCloudResourceCertResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudResourceCertResponse) GoString() string {
	return s.String()
}

func (s *ModifyCloudResourceCertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCloudResourceCertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCloudResourceCertResponse) GetBody() *ModifyCloudResourceCertResponseBody {
	return s.Body
}

func (s *ModifyCloudResourceCertResponse) SetHeaders(v map[string]*string) *ModifyCloudResourceCertResponse {
	s.Headers = v
	return s
}

func (s *ModifyCloudResourceCertResponse) SetStatusCode(v int32) *ModifyCloudResourceCertResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCloudResourceCertResponse) SetBody(v *ModifyCloudResourceCertResponseBody) *ModifyCloudResourceCertResponse {
	s.Body = v
	return s
}

func (s *ModifyCloudResourceCertResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
