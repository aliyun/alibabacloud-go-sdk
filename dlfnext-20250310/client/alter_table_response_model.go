// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlterTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AlterTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AlterTableResponse
	GetStatusCode() *int32
}

type AlterTableResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s AlterTableResponse) String() string {
	return dara.Prettify(s)
}

func (s AlterTableResponse) GoString() string {
	return s.String()
}

func (s *AlterTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AlterTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AlterTableResponse) SetHeaders(v map[string]*string) *AlterTableResponse {
	s.Headers = v
	return s
}

func (s *AlterTableResponse) SetStatusCode(v int32) *AlterTableResponse {
	s.StatusCode = &v
	return s
}

func (s *AlterTableResponse) Validate() error {
	return dara.Validate(s)
}
