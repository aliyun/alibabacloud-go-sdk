// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLakebaseMountParameterGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLakebaseMountParameterGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLakebaseMountParameterGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateLakebaseMountParameterGroupResponseBody) *CreateLakebaseMountParameterGroupResponse
	GetBody() *CreateLakebaseMountParameterGroupResponseBody
}

type CreateLakebaseMountParameterGroupResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLakebaseMountParameterGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLakebaseMountParameterGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLakebaseMountParameterGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateLakebaseMountParameterGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLakebaseMountParameterGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLakebaseMountParameterGroupResponse) GetBody() *CreateLakebaseMountParameterGroupResponseBody {
	return s.Body
}

func (s *CreateLakebaseMountParameterGroupResponse) SetHeaders(v map[string]*string) *CreateLakebaseMountParameterGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateLakebaseMountParameterGroupResponse) SetStatusCode(v int32) *CreateLakebaseMountParameterGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLakebaseMountParameterGroupResponse) SetBody(v *CreateLakebaseMountParameterGroupResponseBody) *CreateLakebaseMountParameterGroupResponse {
	s.Body = v
	return s
}

func (s *CreateLakebaseMountParameterGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
