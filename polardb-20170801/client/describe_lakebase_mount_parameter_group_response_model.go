// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLakebaseMountParameterGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLakebaseMountParameterGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLakebaseMountParameterGroupResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLakebaseMountParameterGroupResponseBody) *DescribeLakebaseMountParameterGroupResponse
	GetBody() *DescribeLakebaseMountParameterGroupResponseBody
}

type DescribeLakebaseMountParameterGroupResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLakebaseMountParameterGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLakebaseMountParameterGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLakebaseMountParameterGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeLakebaseMountParameterGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLakebaseMountParameterGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLakebaseMountParameterGroupResponse) GetBody() *DescribeLakebaseMountParameterGroupResponseBody {
	return s.Body
}

func (s *DescribeLakebaseMountParameterGroupResponse) SetHeaders(v map[string]*string) *DescribeLakebaseMountParameterGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeLakebaseMountParameterGroupResponse) SetStatusCode(v int32) *DescribeLakebaseMountParameterGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupResponse) SetBody(v *DescribeLakebaseMountParameterGroupResponseBody) *DescribeLakebaseMountParameterGroupResponse {
	s.Body = v
	return s
}

func (s *DescribeLakebaseMountParameterGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
