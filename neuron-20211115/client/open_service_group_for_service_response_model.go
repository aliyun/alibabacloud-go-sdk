// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenServiceGroupForServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenServiceGroupForServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenServiceGroupForServiceResponse
	GetStatusCode() *int32
	SetBody(v *OpenServiceGroupForServiceResponseBody) *OpenServiceGroupForServiceResponse
	GetBody() *OpenServiceGroupForServiceResponseBody
}

type OpenServiceGroupForServiceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenServiceGroupForServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenServiceGroupForServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenServiceGroupForServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenServiceGroupForServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenServiceGroupForServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenServiceGroupForServiceResponse) GetBody() *OpenServiceGroupForServiceResponseBody {
	return s.Body
}

func (s *OpenServiceGroupForServiceResponse) SetHeaders(v map[string]*string) *OpenServiceGroupForServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenServiceGroupForServiceResponse) SetStatusCode(v int32) *OpenServiceGroupForServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenServiceGroupForServiceResponse) SetBody(v *OpenServiceGroupForServiceResponseBody) *OpenServiceGroupForServiceResponse {
	s.Body = v
	return s
}

func (s *OpenServiceGroupForServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
