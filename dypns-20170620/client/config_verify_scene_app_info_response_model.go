// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigVerifySceneAppInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigVerifySceneAppInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigVerifySceneAppInfoResponse
	GetStatusCode() *int32
	SetBody(v *ConfigVerifySceneAppInfoResponseBody) *ConfigVerifySceneAppInfoResponse
	GetBody() *ConfigVerifySceneAppInfoResponseBody
}

type ConfigVerifySceneAppInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigVerifySceneAppInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigVerifySceneAppInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigVerifySceneAppInfoResponse) GoString() string {
	return s.String()
}

func (s *ConfigVerifySceneAppInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigVerifySceneAppInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigVerifySceneAppInfoResponse) GetBody() *ConfigVerifySceneAppInfoResponseBody {
	return s.Body
}

func (s *ConfigVerifySceneAppInfoResponse) SetHeaders(v map[string]*string) *ConfigVerifySceneAppInfoResponse {
	s.Headers = v
	return s
}

func (s *ConfigVerifySceneAppInfoResponse) SetStatusCode(v int32) *ConfigVerifySceneAppInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigVerifySceneAppInfoResponse) SetBody(v *ConfigVerifySceneAppInfoResponseBody) *ConfigVerifySceneAppInfoResponse {
	s.Body = v
	return s
}

func (s *ConfigVerifySceneAppInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
