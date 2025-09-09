// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackInstanceVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RollbackInstanceVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RollbackInstanceVersionResponse
	GetStatusCode() *int32
	SetBody(v *RollbackInstanceVersionResponseBody) *RollbackInstanceVersionResponse
	GetBody() *RollbackInstanceVersionResponseBody
}

type RollbackInstanceVersionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RollbackInstanceVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RollbackInstanceVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s RollbackInstanceVersionResponse) GoString() string {
	return s.String()
}

func (s *RollbackInstanceVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RollbackInstanceVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RollbackInstanceVersionResponse) GetBody() *RollbackInstanceVersionResponseBody {
	return s.Body
}

func (s *RollbackInstanceVersionResponse) SetHeaders(v map[string]*string) *RollbackInstanceVersionResponse {
	s.Headers = v
	return s
}

func (s *RollbackInstanceVersionResponse) SetStatusCode(v int32) *RollbackInstanceVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *RollbackInstanceVersionResponse) SetBody(v *RollbackInstanceVersionResponseBody) *RollbackInstanceVersionResponse {
	s.Body = v
	return s
}

func (s *RollbackInstanceVersionResponse) Validate() error {
	return dara.Validate(s)
}
