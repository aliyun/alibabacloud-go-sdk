// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyComputeBurstConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyComputeBurstConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyComputeBurstConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyComputeBurstConfigResponseBody) *ModifyComputeBurstConfigResponse
	GetBody() *ModifyComputeBurstConfigResponseBody
}

type ModifyComputeBurstConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyComputeBurstConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyComputeBurstConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyComputeBurstConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyComputeBurstConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyComputeBurstConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyComputeBurstConfigResponse) GetBody() *ModifyComputeBurstConfigResponseBody {
	return s.Body
}

func (s *ModifyComputeBurstConfigResponse) SetHeaders(v map[string]*string) *ModifyComputeBurstConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyComputeBurstConfigResponse) SetStatusCode(v int32) *ModifyComputeBurstConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyComputeBurstConfigResponse) SetBody(v *ModifyComputeBurstConfigResponseBody) *ModifyComputeBurstConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyComputeBurstConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
