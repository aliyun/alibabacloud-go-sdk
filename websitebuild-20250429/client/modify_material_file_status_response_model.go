// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMaterialFileStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyMaterialFileStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyMaterialFileStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModifyMaterialFileStatusResponseBody) *ModifyMaterialFileStatusResponse
	GetBody() *ModifyMaterialFileStatusResponseBody
}

type ModifyMaterialFileStatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyMaterialFileStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyMaterialFileStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyMaterialFileStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyMaterialFileStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyMaterialFileStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyMaterialFileStatusResponse) GetBody() *ModifyMaterialFileStatusResponseBody {
	return s.Body
}

func (s *ModifyMaterialFileStatusResponse) SetHeaders(v map[string]*string) *ModifyMaterialFileStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyMaterialFileStatusResponse) SetStatusCode(v int32) *ModifyMaterialFileStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMaterialFileStatusResponse) SetBody(v *ModifyMaterialFileStatusResponseBody) *ModifyMaterialFileStatusResponse {
	s.Body = v
	return s
}

func (s *ModifyMaterialFileStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
