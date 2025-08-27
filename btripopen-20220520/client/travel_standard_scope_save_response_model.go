// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTravelStandardScopeSaveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TravelStandardScopeSaveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TravelStandardScopeSaveResponse
	GetStatusCode() *int32
	SetBody(v *TravelStandardScopeSaveResponseBody) *TravelStandardScopeSaveResponse
	GetBody() *TravelStandardScopeSaveResponseBody
}

type TravelStandardScopeSaveResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TravelStandardScopeSaveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TravelStandardScopeSaveResponse) String() string {
	return dara.Prettify(s)
}

func (s TravelStandardScopeSaveResponse) GoString() string {
	return s.String()
}

func (s *TravelStandardScopeSaveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TravelStandardScopeSaveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TravelStandardScopeSaveResponse) GetBody() *TravelStandardScopeSaveResponseBody {
	return s.Body
}

func (s *TravelStandardScopeSaveResponse) SetHeaders(v map[string]*string) *TravelStandardScopeSaveResponse {
	s.Headers = v
	return s
}

func (s *TravelStandardScopeSaveResponse) SetStatusCode(v int32) *TravelStandardScopeSaveResponse {
	s.StatusCode = &v
	return s
}

func (s *TravelStandardScopeSaveResponse) SetBody(v *TravelStandardScopeSaveResponseBody) *TravelStandardScopeSaveResponse {
	s.Body = v
	return s
}

func (s *TravelStandardScopeSaveResponse) Validate() error {
	return dara.Validate(s)
}
