// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBundleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyBundleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyBundleResponse
	GetStatusCode() *int32
	SetBody(v *ModifyBundleResponseBody) *ModifyBundleResponse
	GetBody() *ModifyBundleResponseBody
}

type ModifyBundleResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyBundleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyBundleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyBundleResponse) GoString() string {
	return s.String()
}

func (s *ModifyBundleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyBundleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyBundleResponse) GetBody() *ModifyBundleResponseBody {
	return s.Body
}

func (s *ModifyBundleResponse) SetHeaders(v map[string]*string) *ModifyBundleResponse {
	s.Headers = v
	return s
}

func (s *ModifyBundleResponse) SetStatusCode(v int32) *ModifyBundleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyBundleResponse) SetBody(v *ModifyBundleResponseBody) *ModifyBundleResponse {
	s.Body = v
	return s
}

func (s *ModifyBundleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
