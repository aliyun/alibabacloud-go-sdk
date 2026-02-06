// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCloudResourceDefaultCertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCloudResourceDefaultCertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCloudResourceDefaultCertResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCloudResourceDefaultCertResponseBody) *ModifyCloudResourceDefaultCertResponse
	GetBody() *ModifyCloudResourceDefaultCertResponseBody
}

type ModifyCloudResourceDefaultCertResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCloudResourceDefaultCertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCloudResourceDefaultCertResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudResourceDefaultCertResponse) GoString() string {
	return s.String()
}

func (s *ModifyCloudResourceDefaultCertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCloudResourceDefaultCertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCloudResourceDefaultCertResponse) GetBody() *ModifyCloudResourceDefaultCertResponseBody {
	return s.Body
}

func (s *ModifyCloudResourceDefaultCertResponse) SetHeaders(v map[string]*string) *ModifyCloudResourceDefaultCertResponse {
	s.Headers = v
	return s
}

func (s *ModifyCloudResourceDefaultCertResponse) SetStatusCode(v int32) *ModifyCloudResourceDefaultCertResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCloudResourceDefaultCertResponse) SetBody(v *ModifyCloudResourceDefaultCertResponseBody) *ModifyCloudResourceDefaultCertResponse {
	s.Body = v
	return s
}

func (s *ModifyCloudResourceDefaultCertResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
