// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExperimentGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListExperimentGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListExperimentGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListExperimentGroupsResponseBody) *ListExperimentGroupsResponse
	GetBody() *ListExperimentGroupsResponseBody
}

type ListExperimentGroupsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExperimentGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExperimentGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListExperimentGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListExperimentGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListExperimentGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListExperimentGroupsResponse) GetBody() *ListExperimentGroupsResponseBody {
	return s.Body
}

func (s *ListExperimentGroupsResponse) SetHeaders(v map[string]*string) *ListExperimentGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListExperimentGroupsResponse) SetStatusCode(v int32) *ListExperimentGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExperimentGroupsResponse) SetBody(v *ListExperimentGroupsResponseBody) *ListExperimentGroupsResponse {
	s.Body = v
	return s
}

func (s *ListExperimentGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
