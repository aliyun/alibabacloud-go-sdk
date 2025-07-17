// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDropDatabaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DropDatabaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DropDatabaseResponse
	GetStatusCode() *int32
}

type DropDatabaseResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DropDatabaseResponse) String() string {
	return dara.Prettify(s)
}

func (s DropDatabaseResponse) GoString() string {
	return s.String()
}

func (s *DropDatabaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DropDatabaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DropDatabaseResponse) SetHeaders(v map[string]*string) *DropDatabaseResponse {
	s.Headers = v
	return s
}

func (s *DropDatabaseResponse) SetStatusCode(v int32) *DropDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *DropDatabaseResponse) Validate() error {
	return dara.Validate(s)
}
