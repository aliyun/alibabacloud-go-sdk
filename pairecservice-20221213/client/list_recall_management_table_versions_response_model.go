// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecallManagementTableVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRecallManagementTableVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRecallManagementTableVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListRecallManagementTableVersionsResponseBody) *ListRecallManagementTableVersionsResponse
	GetBody() *ListRecallManagementTableVersionsResponseBody
}

type ListRecallManagementTableVersionsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRecallManagementTableVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRecallManagementTableVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRecallManagementTableVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListRecallManagementTableVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRecallManagementTableVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRecallManagementTableVersionsResponse) GetBody() *ListRecallManagementTableVersionsResponseBody {
	return s.Body
}

func (s *ListRecallManagementTableVersionsResponse) SetHeaders(v map[string]*string) *ListRecallManagementTableVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListRecallManagementTableVersionsResponse) SetStatusCode(v int32) *ListRecallManagementTableVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRecallManagementTableVersionsResponse) SetBody(v *ListRecallManagementTableVersionsResponseBody) *ListRecallManagementTableVersionsResponse {
	s.Body = v
	return s
}

func (s *ListRecallManagementTableVersionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
