// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLakebaseMountParameterGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLakebaseMountParameterGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLakebaseMountParameterGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLakebaseMountParameterGroupResponseBody) *DeleteLakebaseMountParameterGroupResponse
	GetBody() *DeleteLakebaseMountParameterGroupResponseBody
}

type DeleteLakebaseMountParameterGroupResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLakebaseMountParameterGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLakebaseMountParameterGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLakebaseMountParameterGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteLakebaseMountParameterGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLakebaseMountParameterGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLakebaseMountParameterGroupResponse) GetBody() *DeleteLakebaseMountParameterGroupResponseBody {
	return s.Body
}

func (s *DeleteLakebaseMountParameterGroupResponse) SetHeaders(v map[string]*string) *DeleteLakebaseMountParameterGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteLakebaseMountParameterGroupResponse) SetStatusCode(v int32) *DeleteLakebaseMountParameterGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLakebaseMountParameterGroupResponse) SetBody(v *DeleteLakebaseMountParameterGroupResponseBody) *DeleteLakebaseMountParameterGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteLakebaseMountParameterGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
