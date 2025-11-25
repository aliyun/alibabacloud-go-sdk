// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebCacheModeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyWebCacheModeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyWebCacheModeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyWebCacheModeResponseBody) *ModifyWebCacheModeResponse
	GetBody() *ModifyWebCacheModeResponseBody
}

type ModifyWebCacheModeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyWebCacheModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyWebCacheModeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebCacheModeResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebCacheModeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyWebCacheModeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyWebCacheModeResponse) GetBody() *ModifyWebCacheModeResponseBody {
	return s.Body
}

func (s *ModifyWebCacheModeResponse) SetHeaders(v map[string]*string) *ModifyWebCacheModeResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebCacheModeResponse) SetStatusCode(v int32) *ModifyWebCacheModeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyWebCacheModeResponse) SetBody(v *ModifyWebCacheModeResponseBody) *ModifyWebCacheModeResponse {
	s.Body = v
	return s
}

func (s *ModifyWebCacheModeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
