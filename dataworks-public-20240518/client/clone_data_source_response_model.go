// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneDataSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloneDataSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloneDataSourceResponse
	GetStatusCode() *int32
	SetBody(v *CloneDataSourceResponseBody) *CloneDataSourceResponse
	GetBody() *CloneDataSourceResponseBody
}

type CloneDataSourceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloneDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloneDataSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s CloneDataSourceResponse) GoString() string {
	return s.String()
}

func (s *CloneDataSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloneDataSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloneDataSourceResponse) GetBody() *CloneDataSourceResponseBody {
	return s.Body
}

func (s *CloneDataSourceResponse) SetHeaders(v map[string]*string) *CloneDataSourceResponse {
	s.Headers = v
	return s
}

func (s *CloneDataSourceResponse) SetStatusCode(v int32) *CloneDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CloneDataSourceResponse) SetBody(v *CloneDataSourceResponseBody) *CloneDataSourceResponse {
	s.Body = v
	return s
}

func (s *CloneDataSourceResponse) Validate() error {
	return dara.Validate(s)
}
