// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetManagerInfoByAuthCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetManagerInfoByAuthCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetManagerInfoByAuthCodeResponse
	GetStatusCode() *int32
	SetBody(v *GetManagerInfoByAuthCodeResponseBody) *GetManagerInfoByAuthCodeResponse
	GetBody() *GetManagerInfoByAuthCodeResponseBody
}

type GetManagerInfoByAuthCodeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetManagerInfoByAuthCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetManagerInfoByAuthCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetManagerInfoByAuthCodeResponse) GoString() string {
	return s.String()
}

func (s *GetManagerInfoByAuthCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetManagerInfoByAuthCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetManagerInfoByAuthCodeResponse) GetBody() *GetManagerInfoByAuthCodeResponseBody {
	return s.Body
}

func (s *GetManagerInfoByAuthCodeResponse) SetHeaders(v map[string]*string) *GetManagerInfoByAuthCodeResponse {
	s.Headers = v
	return s
}

func (s *GetManagerInfoByAuthCodeResponse) SetStatusCode(v int32) *GetManagerInfoByAuthCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetManagerInfoByAuthCodeResponse) SetBody(v *GetManagerInfoByAuthCodeResponseBody) *GetManagerInfoByAuthCodeResponse {
	s.Body = v
	return s
}

func (s *GetManagerInfoByAuthCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
