// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstalledSoftwaresResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstalledSoftwaresResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstalledSoftwaresResponse
	GetStatusCode() *int32
	SetBody(v *ListInstalledSoftwaresResponseBody) *ListInstalledSoftwaresResponse
	GetBody() *ListInstalledSoftwaresResponseBody
}

type ListInstalledSoftwaresResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstalledSoftwaresResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstalledSoftwaresResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstalledSoftwaresResponse) GoString() string {
	return s.String()
}

func (s *ListInstalledSoftwaresResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstalledSoftwaresResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstalledSoftwaresResponse) GetBody() *ListInstalledSoftwaresResponseBody {
	return s.Body
}

func (s *ListInstalledSoftwaresResponse) SetHeaders(v map[string]*string) *ListInstalledSoftwaresResponse {
	s.Headers = v
	return s
}

func (s *ListInstalledSoftwaresResponse) SetStatusCode(v int32) *ListInstalledSoftwaresResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstalledSoftwaresResponse) SetBody(v *ListInstalledSoftwaresResponseBody) *ListInstalledSoftwaresResponse {
	s.Body = v
	return s
}

func (s *ListInstalledSoftwaresResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
