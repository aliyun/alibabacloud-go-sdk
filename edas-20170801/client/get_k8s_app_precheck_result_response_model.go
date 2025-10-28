// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetK8sAppPrecheckResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetK8sAppPrecheckResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetK8sAppPrecheckResultResponse
	GetStatusCode() *int32
	SetBody(v *GetK8sAppPrecheckResultResponseBody) *GetK8sAppPrecheckResultResponse
	GetBody() *GetK8sAppPrecheckResultResponseBody
}

type GetK8sAppPrecheckResultResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetK8sAppPrecheckResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetK8sAppPrecheckResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetK8sAppPrecheckResultResponse) GoString() string {
	return s.String()
}

func (s *GetK8sAppPrecheckResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetK8sAppPrecheckResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetK8sAppPrecheckResultResponse) GetBody() *GetK8sAppPrecheckResultResponseBody {
	return s.Body
}

func (s *GetK8sAppPrecheckResultResponse) SetHeaders(v map[string]*string) *GetK8sAppPrecheckResultResponse {
	s.Headers = v
	return s
}

func (s *GetK8sAppPrecheckResultResponse) SetStatusCode(v int32) *GetK8sAppPrecheckResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetK8sAppPrecheckResultResponse) SetBody(v *GetK8sAppPrecheckResultResponseBody) *GetK8sAppPrecheckResultResponse {
	s.Body = v
	return s
}

func (s *GetK8sAppPrecheckResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
