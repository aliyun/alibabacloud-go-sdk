// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceFoldersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataServiceFoldersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataServiceFoldersResponse
	GetStatusCode() *int32
	SetBody(v *ListDataServiceFoldersResponseBody) *ListDataServiceFoldersResponse
	GetBody() *ListDataServiceFoldersResponseBody
}

type ListDataServiceFoldersResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataServiceFoldersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataServiceFoldersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceFoldersResponse) GoString() string {
	return s.String()
}

func (s *ListDataServiceFoldersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataServiceFoldersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataServiceFoldersResponse) GetBody() *ListDataServiceFoldersResponseBody {
	return s.Body
}

func (s *ListDataServiceFoldersResponse) SetHeaders(v map[string]*string) *ListDataServiceFoldersResponse {
	s.Headers = v
	return s
}

func (s *ListDataServiceFoldersResponse) SetStatusCode(v int32) *ListDataServiceFoldersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataServiceFoldersResponse) SetBody(v *ListDataServiceFoldersResponseBody) *ListDataServiceFoldersResponse {
	s.Body = v
	return s
}

func (s *ListDataServiceFoldersResponse) Validate() error {
	return dara.Validate(s)
}
