// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMaterialFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyMaterialFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyMaterialFileResponse
	GetStatusCode() *int32
	SetBody(v *ModifyMaterialFileResponseBody) *ModifyMaterialFileResponse
	GetBody() *ModifyMaterialFileResponseBody
}

type ModifyMaterialFileResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyMaterialFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyMaterialFileResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyMaterialFileResponse) GoString() string {
	return s.String()
}

func (s *ModifyMaterialFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyMaterialFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyMaterialFileResponse) GetBody() *ModifyMaterialFileResponseBody {
	return s.Body
}

func (s *ModifyMaterialFileResponse) SetHeaders(v map[string]*string) *ModifyMaterialFileResponse {
	s.Headers = v
	return s
}

func (s *ModifyMaterialFileResponse) SetStatusCode(v int32) *ModifyMaterialFileResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMaterialFileResponse) SetBody(v *ModifyMaterialFileResponseBody) *ModifyMaterialFileResponse {
	s.Body = v
	return s
}

func (s *ModifyMaterialFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
