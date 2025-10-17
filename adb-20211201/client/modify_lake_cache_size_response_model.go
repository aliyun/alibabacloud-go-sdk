// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLakeCacheSizeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyLakeCacheSizeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyLakeCacheSizeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyLakeCacheSizeResponseBody) *ModifyLakeCacheSizeResponse
	GetBody() *ModifyLakeCacheSizeResponseBody
}

type ModifyLakeCacheSizeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyLakeCacheSizeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyLakeCacheSizeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyLakeCacheSizeResponse) GoString() string {
	return s.String()
}

func (s *ModifyLakeCacheSizeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyLakeCacheSizeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyLakeCacheSizeResponse) GetBody() *ModifyLakeCacheSizeResponseBody {
	return s.Body
}

func (s *ModifyLakeCacheSizeResponse) SetHeaders(v map[string]*string) *ModifyLakeCacheSizeResponse {
	s.Headers = v
	return s
}

func (s *ModifyLakeCacheSizeResponse) SetStatusCode(v int32) *ModifyLakeCacheSizeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLakeCacheSizeResponse) SetBody(v *ModifyLakeCacheSizeResponseBody) *ModifyLakeCacheSizeResponse {
	s.Body = v
	return s
}

func (s *ModifyLakeCacheSizeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
