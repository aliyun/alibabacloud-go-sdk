// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAliyunPKByAliyunIDResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAliyunPKByAliyunIDResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAliyunPKByAliyunIDResponse
	GetStatusCode() *int32
	SetBody(v *GetAliyunPKByAliyunIDResponseBody) *GetAliyunPKByAliyunIDResponse
	GetBody() *GetAliyunPKByAliyunIDResponseBody
}

type GetAliyunPKByAliyunIDResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAliyunPKByAliyunIDResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAliyunPKByAliyunIDResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAliyunPKByAliyunIDResponse) GoString() string {
	return s.String()
}

func (s *GetAliyunPKByAliyunIDResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAliyunPKByAliyunIDResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAliyunPKByAliyunIDResponse) GetBody() *GetAliyunPKByAliyunIDResponseBody {
	return s.Body
}

func (s *GetAliyunPKByAliyunIDResponse) SetHeaders(v map[string]*string) *GetAliyunPKByAliyunIDResponse {
	s.Headers = v
	return s
}

func (s *GetAliyunPKByAliyunIDResponse) SetStatusCode(v int32) *GetAliyunPKByAliyunIDResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAliyunPKByAliyunIDResponse) SetBody(v *GetAliyunPKByAliyunIDResponseBody) *GetAliyunPKByAliyunIDResponse {
	s.Body = v
	return s
}

func (s *GetAliyunPKByAliyunIDResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
