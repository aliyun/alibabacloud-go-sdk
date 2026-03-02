// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGrantRolesToDeveloperResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchGrantRolesToDeveloperResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchGrantRolesToDeveloperResponse
	GetStatusCode() *int32
}

type BatchGrantRolesToDeveloperResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s BatchGrantRolesToDeveloperResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchGrantRolesToDeveloperResponse) GoString() string {
	return s.String()
}

func (s *BatchGrantRolesToDeveloperResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchGrantRolesToDeveloperResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchGrantRolesToDeveloperResponse) SetHeaders(v map[string]*string) *BatchGrantRolesToDeveloperResponse {
	s.Headers = v
	return s
}

func (s *BatchGrantRolesToDeveloperResponse) SetStatusCode(v int32) *BatchGrantRolesToDeveloperResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchGrantRolesToDeveloperResponse) Validate() error {
	return dara.Validate(s)
}
