// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIntranetAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyIntranetAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyIntranetAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyIntranetAttributeResponseBody) *ModifyIntranetAttributeResponse
	GetBody() *ModifyIntranetAttributeResponseBody
}

type ModifyIntranetAttributeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyIntranetAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyIntranetAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyIntranetAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyIntranetAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyIntranetAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyIntranetAttributeResponse) GetBody() *ModifyIntranetAttributeResponseBody {
	return s.Body
}

func (s *ModifyIntranetAttributeResponse) SetHeaders(v map[string]*string) *ModifyIntranetAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyIntranetAttributeResponse) SetStatusCode(v int32) *ModifyIntranetAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyIntranetAttributeResponse) SetBody(v *ModifyIntranetAttributeResponseBody) *ModifyIntranetAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyIntranetAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
