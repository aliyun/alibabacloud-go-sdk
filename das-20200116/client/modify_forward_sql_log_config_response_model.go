// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyForwardSqlLogConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyForwardSqlLogConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyForwardSqlLogConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyForwardSqlLogConfigResponseBody) *ModifyForwardSqlLogConfigResponse
	GetBody() *ModifyForwardSqlLogConfigResponseBody
}

type ModifyForwardSqlLogConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyForwardSqlLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyForwardSqlLogConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyForwardSqlLogConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyForwardSqlLogConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyForwardSqlLogConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyForwardSqlLogConfigResponse) GetBody() *ModifyForwardSqlLogConfigResponseBody {
	return s.Body
}

func (s *ModifyForwardSqlLogConfigResponse) SetHeaders(v map[string]*string) *ModifyForwardSqlLogConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyForwardSqlLogConfigResponse) SetStatusCode(v int32) *ModifyForwardSqlLogConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyForwardSqlLogConfigResponse) SetBody(v *ModifyForwardSqlLogConfigResponseBody) *ModifyForwardSqlLogConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyForwardSqlLogConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
