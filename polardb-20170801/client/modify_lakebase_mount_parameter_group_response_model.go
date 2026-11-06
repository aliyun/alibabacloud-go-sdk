// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLakebaseMountParameterGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyLakebaseMountParameterGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyLakebaseMountParameterGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModifyLakebaseMountParameterGroupResponseBody) *ModifyLakebaseMountParameterGroupResponse
	GetBody() *ModifyLakebaseMountParameterGroupResponseBody
}

type ModifyLakebaseMountParameterGroupResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyLakebaseMountParameterGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyLakebaseMountParameterGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyLakebaseMountParameterGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyLakebaseMountParameterGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyLakebaseMountParameterGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyLakebaseMountParameterGroupResponse) GetBody() *ModifyLakebaseMountParameterGroupResponseBody {
	return s.Body
}

func (s *ModifyLakebaseMountParameterGroupResponse) SetHeaders(v map[string]*string) *ModifyLakebaseMountParameterGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyLakebaseMountParameterGroupResponse) SetStatusCode(v int32) *ModifyLakebaseMountParameterGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLakebaseMountParameterGroupResponse) SetBody(v *ModifyLakebaseMountParameterGroupResponseBody) *ModifyLakebaseMountParameterGroupResponse {
	s.Body = v
	return s
}

func (s *ModifyLakebaseMountParameterGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
