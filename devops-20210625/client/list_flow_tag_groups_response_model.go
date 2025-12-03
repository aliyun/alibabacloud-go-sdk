// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlowTagGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFlowTagGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFlowTagGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListFlowTagGroupsResponseBody) *ListFlowTagGroupsResponse
	GetBody() *ListFlowTagGroupsResponseBody
}

type ListFlowTagGroupsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFlowTagGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFlowTagGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFlowTagGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListFlowTagGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFlowTagGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFlowTagGroupsResponse) GetBody() *ListFlowTagGroupsResponseBody {
	return s.Body
}

func (s *ListFlowTagGroupsResponse) SetHeaders(v map[string]*string) *ListFlowTagGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListFlowTagGroupsResponse) SetStatusCode(v int32) *ListFlowTagGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFlowTagGroupsResponse) SetBody(v *ListFlowTagGroupsResponseBody) *ListFlowTagGroupsResponse {
	s.Body = v
	return s
}

func (s *ListFlowTagGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
