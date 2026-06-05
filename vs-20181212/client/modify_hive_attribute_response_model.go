// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHiveAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyHiveAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyHiveAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyHiveAttributeResponseBody) *ModifyHiveAttributeResponse
	GetBody() *ModifyHiveAttributeResponseBody
}

type ModifyHiveAttributeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHiveAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyHiveAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyHiveAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyHiveAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyHiveAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyHiveAttributeResponse) GetBody() *ModifyHiveAttributeResponseBody {
	return s.Body
}

func (s *ModifyHiveAttributeResponse) SetHeaders(v map[string]*string) *ModifyHiveAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyHiveAttributeResponse) SetStatusCode(v int32) *ModifyHiveAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHiveAttributeResponse) SetBody(v *ModifyHiveAttributeResponseBody) *ModifyHiveAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyHiveAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
