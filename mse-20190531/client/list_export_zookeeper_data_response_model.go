// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExportZookeeperDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListExportZookeeperDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListExportZookeeperDataResponse
	GetStatusCode() *int32
	SetBody(v *ListExportZookeeperDataResponseBody) *ListExportZookeeperDataResponse
	GetBody() *ListExportZookeeperDataResponseBody
}

type ListExportZookeeperDataResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExportZookeeperDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExportZookeeperDataResponse) String() string {
	return dara.Prettify(s)
}

func (s ListExportZookeeperDataResponse) GoString() string {
	return s.String()
}

func (s *ListExportZookeeperDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListExportZookeeperDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListExportZookeeperDataResponse) GetBody() *ListExportZookeeperDataResponseBody {
	return s.Body
}

func (s *ListExportZookeeperDataResponse) SetHeaders(v map[string]*string) *ListExportZookeeperDataResponse {
	s.Headers = v
	return s
}

func (s *ListExportZookeeperDataResponse) SetStatusCode(v int32) *ListExportZookeeperDataResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExportZookeeperDataResponse) SetBody(v *ListExportZookeeperDataResponseBody) *ListExportZookeeperDataResponse {
	s.Body = v
	return s
}

func (s *ListExportZookeeperDataResponse) Validate() error {
	return dara.Validate(s)
}
