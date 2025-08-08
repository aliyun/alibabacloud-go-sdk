// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlterShareResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AlterShareResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AlterShareResponse
	GetStatusCode() *int32
}

type AlterShareResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s AlterShareResponse) String() string {
	return dara.Prettify(s)
}

func (s AlterShareResponse) GoString() string {
	return s.String()
}

func (s *AlterShareResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AlterShareResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AlterShareResponse) SetHeaders(v map[string]*string) *AlterShareResponse {
	s.Headers = v
	return s
}

func (s *AlterShareResponse) SetStatusCode(v int32) *AlterShareResponse {
	s.StatusCode = &v
	return s
}

func (s *AlterShareResponse) Validate() error {
	return dara.Validate(s)
}
