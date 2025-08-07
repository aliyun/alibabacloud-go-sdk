// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTempModifyDBNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TempModifyDBNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TempModifyDBNodeResponse
	GetStatusCode() *int32
	SetBody(v *TempModifyDBNodeResponseBody) *TempModifyDBNodeResponse
	GetBody() *TempModifyDBNodeResponseBody
}

type TempModifyDBNodeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TempModifyDBNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TempModifyDBNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s TempModifyDBNodeResponse) GoString() string {
	return s.String()
}

func (s *TempModifyDBNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TempModifyDBNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TempModifyDBNodeResponse) GetBody() *TempModifyDBNodeResponseBody {
	return s.Body
}

func (s *TempModifyDBNodeResponse) SetHeaders(v map[string]*string) *TempModifyDBNodeResponse {
	s.Headers = v
	return s
}

func (s *TempModifyDBNodeResponse) SetStatusCode(v int32) *TempModifyDBNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *TempModifyDBNodeResponse) SetBody(v *TempModifyDBNodeResponseBody) *TempModifyDBNodeResponse {
	s.Body = v
	return s
}

func (s *TempModifyDBNodeResponse) Validate() error {
	return dara.Validate(s)
}
