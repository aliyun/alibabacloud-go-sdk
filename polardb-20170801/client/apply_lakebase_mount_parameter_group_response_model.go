// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyLakebaseMountParameterGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyLakebaseMountParameterGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyLakebaseMountParameterGroupResponse
	GetStatusCode() *int32
	SetBody(v *ApplyLakebaseMountParameterGroupResponseBody) *ApplyLakebaseMountParameterGroupResponse
	GetBody() *ApplyLakebaseMountParameterGroupResponseBody
}

type ApplyLakebaseMountParameterGroupResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyLakebaseMountParameterGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyLakebaseMountParameterGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyLakebaseMountParameterGroupResponse) GoString() string {
	return s.String()
}

func (s *ApplyLakebaseMountParameterGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyLakebaseMountParameterGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyLakebaseMountParameterGroupResponse) GetBody() *ApplyLakebaseMountParameterGroupResponseBody {
	return s.Body
}

func (s *ApplyLakebaseMountParameterGroupResponse) SetHeaders(v map[string]*string) *ApplyLakebaseMountParameterGroupResponse {
	s.Headers = v
	return s
}

func (s *ApplyLakebaseMountParameterGroupResponse) SetStatusCode(v int32) *ApplyLakebaseMountParameterGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyLakebaseMountParameterGroupResponse) SetBody(v *ApplyLakebaseMountParameterGroupResponseBody) *ApplyLakebaseMountParameterGroupResponse {
	s.Body = v
	return s
}

func (s *ApplyLakebaseMountParameterGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
