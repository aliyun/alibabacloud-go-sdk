// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComponentAssetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateComponentAssetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateComponentAssetResponse
	GetStatusCode() *int32
	SetBody(v *UpdateComponentAssetResponseBody) *UpdateComponentAssetResponse
	GetBody() *UpdateComponentAssetResponseBody
}

type UpdateComponentAssetResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateComponentAssetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateComponentAssetResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateComponentAssetResponse) GoString() string {
	return s.String()
}

func (s *UpdateComponentAssetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateComponentAssetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateComponentAssetResponse) GetBody() *UpdateComponentAssetResponseBody {
	return s.Body
}

func (s *UpdateComponentAssetResponse) SetHeaders(v map[string]*string) *UpdateComponentAssetResponse {
	s.Headers = v
	return s
}

func (s *UpdateComponentAssetResponse) SetStatusCode(v int32) *UpdateComponentAssetResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateComponentAssetResponse) SetBody(v *UpdateComponentAssetResponseBody) *UpdateComponentAssetResponse {
	s.Body = v
	return s
}

func (s *UpdateComponentAssetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
