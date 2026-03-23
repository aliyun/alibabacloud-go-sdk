// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApAssetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApAssetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApAssetResponse
	GetStatusCode() *int32
	SetBody(v *GetApAssetResponseBody) *GetApAssetResponse
	GetBody() *GetApAssetResponseBody
}

type GetApAssetResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApAssetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApAssetResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApAssetResponse) GoString() string {
	return s.String()
}

func (s *GetApAssetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApAssetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApAssetResponse) GetBody() *GetApAssetResponseBody {
	return s.Body
}

func (s *GetApAssetResponse) SetHeaders(v map[string]*string) *GetApAssetResponse {
	s.Headers = v
	return s
}

func (s *GetApAssetResponse) SetStatusCode(v int32) *GetApAssetResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApAssetResponse) SetBody(v *GetApAssetResponseBody) *GetApAssetResponse {
	s.Body = v
	return s
}

func (s *GetApAssetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
