// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlterShareResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AlterShareResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AlterShareResourcesResponse
	GetStatusCode() *int32
}

type AlterShareResourcesResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s AlterShareResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s AlterShareResourcesResponse) GoString() string {
	return s.String()
}

func (s *AlterShareResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AlterShareResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AlterShareResourcesResponse) SetHeaders(v map[string]*string) *AlterShareResourcesResponse {
	s.Headers = v
	return s
}

func (s *AlterShareResourcesResponse) SetStatusCode(v int32) *AlterShareResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *AlterShareResourcesResponse) Validate() error {
	return dara.Validate(s)
}
