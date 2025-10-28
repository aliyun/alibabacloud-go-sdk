// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetK8sApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetK8sApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetK8sApplicationResponse
	GetStatusCode() *int32
	SetBody(v *GetK8sApplicationResponseBody) *GetK8sApplicationResponse
	GetBody() *GetK8sApplicationResponseBody
}

type GetK8sApplicationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetK8sApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetK8sApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetK8sApplicationResponse) GoString() string {
	return s.String()
}

func (s *GetK8sApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetK8sApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetK8sApplicationResponse) GetBody() *GetK8sApplicationResponseBody {
	return s.Body
}

func (s *GetK8sApplicationResponse) SetHeaders(v map[string]*string) *GetK8sApplicationResponse {
	s.Headers = v
	return s
}

func (s *GetK8sApplicationResponse) SetStatusCode(v int32) *GetK8sApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetK8sApplicationResponse) SetBody(v *GetK8sApplicationResponseBody) *GetK8sApplicationResponse {
	s.Body = v
	return s
}

func (s *GetK8sApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
