// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenDcdnServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenDcdnServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenDcdnServiceResponse
	GetStatusCode() *int32
	SetBody(v *OpenDcdnServiceResponseBody) *OpenDcdnServiceResponse
	GetBody() *OpenDcdnServiceResponseBody
}

type OpenDcdnServiceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenDcdnServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenDcdnServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenDcdnServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenDcdnServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenDcdnServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenDcdnServiceResponse) GetBody() *OpenDcdnServiceResponseBody {
	return s.Body
}

func (s *OpenDcdnServiceResponse) SetHeaders(v map[string]*string) *OpenDcdnServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenDcdnServiceResponse) SetStatusCode(v int32) *OpenDcdnServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenDcdnServiceResponse) SetBody(v *OpenDcdnServiceResponseBody) *OpenDcdnServiceResponse {
	s.Body = v
	return s
}

func (s *OpenDcdnServiceResponse) Validate() error {
	return dara.Validate(s)
}
