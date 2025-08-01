// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportZookeeperDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportZookeeperDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportZookeeperDataResponse
	GetStatusCode() *int32
	SetBody(v *ImportZookeeperDataResponseBody) *ImportZookeeperDataResponse
	GetBody() *ImportZookeeperDataResponseBody
}

type ImportZookeeperDataResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportZookeeperDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportZookeeperDataResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportZookeeperDataResponse) GoString() string {
	return s.String()
}

func (s *ImportZookeeperDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportZookeeperDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportZookeeperDataResponse) GetBody() *ImportZookeeperDataResponseBody {
	return s.Body
}

func (s *ImportZookeeperDataResponse) SetHeaders(v map[string]*string) *ImportZookeeperDataResponse {
	s.Headers = v
	return s
}

func (s *ImportZookeeperDataResponse) SetStatusCode(v int32) *ImportZookeeperDataResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportZookeeperDataResponse) SetBody(v *ImportZookeeperDataResponseBody) *ImportZookeeperDataResponse {
	s.Body = v
	return s
}

func (s *ImportZookeeperDataResponse) Validate() error {
	return dara.Validate(s)
}
