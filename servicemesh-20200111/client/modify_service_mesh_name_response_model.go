// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyServiceMeshNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyServiceMeshNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyServiceMeshNameResponse
	GetStatusCode() *int32
	SetBody(v *ModifyServiceMeshNameResponseBody) *ModifyServiceMeshNameResponse
	GetBody() *ModifyServiceMeshNameResponseBody
}

type ModifyServiceMeshNameResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyServiceMeshNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyServiceMeshNameResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyServiceMeshNameResponse) GoString() string {
	return s.String()
}

func (s *ModifyServiceMeshNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyServiceMeshNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyServiceMeshNameResponse) GetBody() *ModifyServiceMeshNameResponseBody {
	return s.Body
}

func (s *ModifyServiceMeshNameResponse) SetHeaders(v map[string]*string) *ModifyServiceMeshNameResponse {
	s.Headers = v
	return s
}

func (s *ModifyServiceMeshNameResponse) SetStatusCode(v int32) *ModifyServiceMeshNameResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyServiceMeshNameResponse) SetBody(v *ModifyServiceMeshNameResponseBody) *ModifyServiceMeshNameResponse {
	s.Body = v
	return s
}

func (s *ModifyServiceMeshNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
