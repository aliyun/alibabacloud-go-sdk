// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RollbackTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RollbackTableResponse
	GetStatusCode() *int32
}

type RollbackTableResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s RollbackTableResponse) String() string {
	return dara.Prettify(s)
}

func (s RollbackTableResponse) GoString() string {
	return s.String()
}

func (s *RollbackTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RollbackTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RollbackTableResponse) SetHeaders(v map[string]*string) *RollbackTableResponse {
	s.Headers = v
	return s
}

func (s *RollbackTableResponse) SetStatusCode(v int32) *RollbackTableResponse {
	s.StatusCode = &v
	return s
}

func (s *RollbackTableResponse) Validate() error {
	return dara.Validate(s)
}
