// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecallManagementServiceVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRecallManagementServiceVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRecallManagementServiceVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListRecallManagementServiceVersionsResponseBody) *ListRecallManagementServiceVersionsResponse
	GetBody() *ListRecallManagementServiceVersionsResponseBody
}

type ListRecallManagementServiceVersionsResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRecallManagementServiceVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRecallManagementServiceVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRecallManagementServiceVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListRecallManagementServiceVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRecallManagementServiceVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRecallManagementServiceVersionsResponse) GetBody() *ListRecallManagementServiceVersionsResponseBody {
	return s.Body
}

func (s *ListRecallManagementServiceVersionsResponse) SetHeaders(v map[string]*string) *ListRecallManagementServiceVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListRecallManagementServiceVersionsResponse) SetStatusCode(v int32) *ListRecallManagementServiceVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRecallManagementServiceVersionsResponse) SetBody(v *ListRecallManagementServiceVersionsResponseBody) *ListRecallManagementServiceVersionsResponse {
	s.Body = v
	return s
}

func (s *ListRecallManagementServiceVersionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
