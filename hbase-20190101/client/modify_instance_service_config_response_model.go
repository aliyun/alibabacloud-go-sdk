// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceServiceConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceServiceConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceServiceConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceServiceConfigResponseBody) *ModifyInstanceServiceConfigResponse
	GetBody() *ModifyInstanceServiceConfigResponseBody
}

type ModifyInstanceServiceConfigResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceServiceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceServiceConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceServiceConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceServiceConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceServiceConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceServiceConfigResponse) GetBody() *ModifyInstanceServiceConfigResponseBody {
	return s.Body
}

func (s *ModifyInstanceServiceConfigResponse) SetHeaders(v map[string]*string) *ModifyInstanceServiceConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceServiceConfigResponse) SetStatusCode(v int32) *ModifyInstanceServiceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceServiceConfigResponse) SetBody(v *ModifyInstanceServiceConfigResponseBody) *ModifyInstanceServiceConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceServiceConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
