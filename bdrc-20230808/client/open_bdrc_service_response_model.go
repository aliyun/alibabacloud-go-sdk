// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenBdrcServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenBdrcServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenBdrcServiceResponse
	GetStatusCode() *int32
	SetBody(v *OpenBdrcServiceResponseBody) *OpenBdrcServiceResponse
	GetBody() *OpenBdrcServiceResponseBody
}

type OpenBdrcServiceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenBdrcServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenBdrcServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenBdrcServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenBdrcServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenBdrcServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenBdrcServiceResponse) GetBody() *OpenBdrcServiceResponseBody {
	return s.Body
}

func (s *OpenBdrcServiceResponse) SetHeaders(v map[string]*string) *OpenBdrcServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenBdrcServiceResponse) SetStatusCode(v int32) *OpenBdrcServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenBdrcServiceResponse) SetBody(v *OpenBdrcServiceResponseBody) *OpenBdrcServiceResponse {
	s.Body = v
	return s
}

func (s *OpenBdrcServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
