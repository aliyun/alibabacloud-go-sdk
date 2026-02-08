// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMaterialDirectoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyMaterialDirectoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyMaterialDirectoryResponse
	GetStatusCode() *int32
	SetBody(v *ModifyMaterialDirectoryResponseBody) *ModifyMaterialDirectoryResponse
	GetBody() *ModifyMaterialDirectoryResponseBody
}

type ModifyMaterialDirectoryResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyMaterialDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyMaterialDirectoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyMaterialDirectoryResponse) GoString() string {
	return s.String()
}

func (s *ModifyMaterialDirectoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyMaterialDirectoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyMaterialDirectoryResponse) GetBody() *ModifyMaterialDirectoryResponseBody {
	return s.Body
}

func (s *ModifyMaterialDirectoryResponse) SetHeaders(v map[string]*string) *ModifyMaterialDirectoryResponse {
	s.Headers = v
	return s
}

func (s *ModifyMaterialDirectoryResponse) SetStatusCode(v int32) *ModifyMaterialDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMaterialDirectoryResponse) SetBody(v *ModifyMaterialDirectoryResponseBody) *ModifyMaterialDirectoryResponse {
	s.Body = v
	return s
}

func (s *ModifyMaterialDirectoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
