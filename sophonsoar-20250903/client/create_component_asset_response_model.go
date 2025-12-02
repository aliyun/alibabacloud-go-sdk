// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComponentAssetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateComponentAssetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateComponentAssetResponse
	GetStatusCode() *int32
	SetBody(v *CreateComponentAssetResponseBody) *CreateComponentAssetResponse
	GetBody() *CreateComponentAssetResponseBody
}

type CreateComponentAssetResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateComponentAssetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateComponentAssetResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateComponentAssetResponse) GoString() string {
	return s.String()
}

func (s *CreateComponentAssetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateComponentAssetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateComponentAssetResponse) GetBody() *CreateComponentAssetResponseBody {
	return s.Body
}

func (s *CreateComponentAssetResponse) SetHeaders(v map[string]*string) *CreateComponentAssetResponse {
	s.Headers = v
	return s
}

func (s *CreateComponentAssetResponse) SetStatusCode(v int32) *CreateComponentAssetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateComponentAssetResponse) SetBody(v *CreateComponentAssetResponseBody) *CreateComponentAssetResponse {
	s.Body = v
	return s
}

func (s *CreateComponentAssetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
