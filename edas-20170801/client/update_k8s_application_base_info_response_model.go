// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateK8sApplicationBaseInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateK8sApplicationBaseInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateK8sApplicationBaseInfoResponse
	GetStatusCode() *int32
	SetBody(v *UpdateK8sApplicationBaseInfoResponseBody) *UpdateK8sApplicationBaseInfoResponse
	GetBody() *UpdateK8sApplicationBaseInfoResponseBody
}

type UpdateK8sApplicationBaseInfoResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateK8sApplicationBaseInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateK8sApplicationBaseInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateK8sApplicationBaseInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateK8sApplicationBaseInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateK8sApplicationBaseInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateK8sApplicationBaseInfoResponse) GetBody() *UpdateK8sApplicationBaseInfoResponseBody {
	return s.Body
}

func (s *UpdateK8sApplicationBaseInfoResponse) SetHeaders(v map[string]*string) *UpdateK8sApplicationBaseInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateK8sApplicationBaseInfoResponse) SetStatusCode(v int32) *UpdateK8sApplicationBaseInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateK8sApplicationBaseInfoResponse) SetBody(v *UpdateK8sApplicationBaseInfoResponseBody) *UpdateK8sApplicationBaseInfoResponse {
	s.Body = v
	return s
}

func (s *UpdateK8sApplicationBaseInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
