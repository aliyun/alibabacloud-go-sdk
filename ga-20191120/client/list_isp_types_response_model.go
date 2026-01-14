// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIspTypesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIspTypesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIspTypesResponse
	GetStatusCode() *int32
	SetBody(v *ListIspTypesResponseBody) *ListIspTypesResponse
	GetBody() *ListIspTypesResponseBody
}

type ListIspTypesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIspTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIspTypesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIspTypesResponse) GoString() string {
	return s.String()
}

func (s *ListIspTypesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIspTypesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIspTypesResponse) GetBody() *ListIspTypesResponseBody {
	return s.Body
}

func (s *ListIspTypesResponse) SetHeaders(v map[string]*string) *ListIspTypesResponse {
	s.Headers = v
	return s
}

func (s *ListIspTypesResponse) SetStatusCode(v int32) *ListIspTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIspTypesResponse) SetBody(v *ListIspTypesResponseBody) *ListIspTypesResponse {
	s.Body = v
	return s
}

func (s *ListIspTypesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
