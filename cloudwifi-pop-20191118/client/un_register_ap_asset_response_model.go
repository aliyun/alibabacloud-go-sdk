// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnRegisterApAssetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnRegisterApAssetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnRegisterApAssetResponse
	GetStatusCode() *int32
	SetBody(v *UnRegisterApAssetResponseBody) *UnRegisterApAssetResponse
	GetBody() *UnRegisterApAssetResponseBody
}

type UnRegisterApAssetResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnRegisterApAssetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnRegisterApAssetResponse) String() string {
	return dara.Prettify(s)
}

func (s UnRegisterApAssetResponse) GoString() string {
	return s.String()
}

func (s *UnRegisterApAssetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnRegisterApAssetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnRegisterApAssetResponse) GetBody() *UnRegisterApAssetResponseBody {
	return s.Body
}

func (s *UnRegisterApAssetResponse) SetHeaders(v map[string]*string) *UnRegisterApAssetResponse {
	s.Headers = v
	return s
}

func (s *UnRegisterApAssetResponse) SetStatusCode(v int32) *UnRegisterApAssetResponse {
	s.StatusCode = &v
	return s
}

func (s *UnRegisterApAssetResponse) SetBody(v *UnRegisterApAssetResponseBody) *UnRegisterApAssetResponse {
	s.Body = v
	return s
}

func (s *UnRegisterApAssetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
