// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlterShareReceiversResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AlterShareReceiversResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AlterShareReceiversResponse
	GetStatusCode() *int32
}

type AlterShareReceiversResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s AlterShareReceiversResponse) String() string {
	return dara.Prettify(s)
}

func (s AlterShareReceiversResponse) GoString() string {
	return s.String()
}

func (s *AlterShareReceiversResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AlterShareReceiversResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AlterShareReceiversResponse) SetHeaders(v map[string]*string) *AlterShareReceiversResponse {
	s.Headers = v
	return s
}

func (s *AlterShareReceiversResponse) SetStatusCode(v int32) *AlterShareReceiversResponse {
	s.StatusCode = &v
	return s
}

func (s *AlterShareReceiversResponse) Validate() error {
	return dara.Validate(s)
}
