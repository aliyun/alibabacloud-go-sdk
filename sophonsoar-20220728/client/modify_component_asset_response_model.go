// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyComponentAssetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyComponentAssetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyComponentAssetResponse
	GetStatusCode() *int32
	SetBody(v *ModifyComponentAssetResponseBody) *ModifyComponentAssetResponse
	GetBody() *ModifyComponentAssetResponseBody
}

type ModifyComponentAssetResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyComponentAssetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyComponentAssetResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyComponentAssetResponse) GoString() string {
	return s.String()
}

func (s *ModifyComponentAssetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyComponentAssetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyComponentAssetResponse) GetBody() *ModifyComponentAssetResponseBody {
	return s.Body
}

func (s *ModifyComponentAssetResponse) SetHeaders(v map[string]*string) *ModifyComponentAssetResponse {
	s.Headers = v
	return s
}

func (s *ModifyComponentAssetResponse) SetStatusCode(v int32) *ModifyComponentAssetResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyComponentAssetResponse) SetBody(v *ModifyComponentAssetResponseBody) *ModifyComponentAssetResponse {
	s.Body = v
	return s
}

func (s *ModifyComponentAssetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
