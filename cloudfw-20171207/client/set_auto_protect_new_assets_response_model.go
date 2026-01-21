// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAutoProtectNewAssetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetAutoProtectNewAssetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetAutoProtectNewAssetsResponse
	GetStatusCode() *int32
	SetBody(v *SetAutoProtectNewAssetsResponseBody) *SetAutoProtectNewAssetsResponse
	GetBody() *SetAutoProtectNewAssetsResponseBody
}

type SetAutoProtectNewAssetsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetAutoProtectNewAssetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetAutoProtectNewAssetsResponse) String() string {
	return dara.Prettify(s)
}

func (s SetAutoProtectNewAssetsResponse) GoString() string {
	return s.String()
}

func (s *SetAutoProtectNewAssetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetAutoProtectNewAssetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetAutoProtectNewAssetsResponse) GetBody() *SetAutoProtectNewAssetsResponseBody {
	return s.Body
}

func (s *SetAutoProtectNewAssetsResponse) SetHeaders(v map[string]*string) *SetAutoProtectNewAssetsResponse {
	s.Headers = v
	return s
}

func (s *SetAutoProtectNewAssetsResponse) SetStatusCode(v int32) *SetAutoProtectNewAssetsResponse {
	s.StatusCode = &v
	return s
}

func (s *SetAutoProtectNewAssetsResponse) SetBody(v *SetAutoProtectNewAssetsResponseBody) *SetAutoProtectNewAssetsResponse {
	s.Body = v
	return s
}

func (s *SetAutoProtectNewAssetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
