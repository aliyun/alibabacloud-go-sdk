// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyResourceLogFieldConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyResourceLogFieldConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyResourceLogFieldConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyResourceLogFieldConfigResponseBody) *ModifyResourceLogFieldConfigResponse
	GetBody() *ModifyResourceLogFieldConfigResponseBody
}

type ModifyResourceLogFieldConfigResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyResourceLogFieldConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyResourceLogFieldConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyResourceLogFieldConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyResourceLogFieldConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyResourceLogFieldConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyResourceLogFieldConfigResponse) GetBody() *ModifyResourceLogFieldConfigResponseBody {
	return s.Body
}

func (s *ModifyResourceLogFieldConfigResponse) SetHeaders(v map[string]*string) *ModifyResourceLogFieldConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyResourceLogFieldConfigResponse) SetStatusCode(v int32) *ModifyResourceLogFieldConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyResourceLogFieldConfigResponse) SetBody(v *ModifyResourceLogFieldConfigResponseBody) *ModifyResourceLogFieldConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyResourceLogFieldConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
