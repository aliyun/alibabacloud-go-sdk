// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAssetCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAssetCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAssetCountResponse
	GetStatusCode() *int32
	SetBody(v *GetAssetCountResponseBody) *GetAssetCountResponse
	GetBody() *GetAssetCountResponseBody
}

type GetAssetCountResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAssetCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAssetCountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAssetCountResponse) GoString() string {
	return s.String()
}

func (s *GetAssetCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAssetCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAssetCountResponse) GetBody() *GetAssetCountResponseBody {
	return s.Body
}

func (s *GetAssetCountResponse) SetHeaders(v map[string]*string) *GetAssetCountResponse {
	s.Headers = v
	return s
}

func (s *GetAssetCountResponse) SetStatusCode(v int32) *GetAssetCountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAssetCountResponse) SetBody(v *GetAssetCountResponseBody) *GetAssetCountResponse {
	s.Body = v
	return s
}

func (s *GetAssetCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
