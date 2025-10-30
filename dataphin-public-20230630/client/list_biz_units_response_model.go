// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBizUnitsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBizUnitsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBizUnitsResponse
	GetStatusCode() *int32
	SetBody(v *ListBizUnitsResponseBody) *ListBizUnitsResponse
	GetBody() *ListBizUnitsResponseBody
}

type ListBizUnitsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBizUnitsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBizUnitsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBizUnitsResponse) GoString() string {
	return s.String()
}

func (s *ListBizUnitsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBizUnitsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBizUnitsResponse) GetBody() *ListBizUnitsResponseBody {
	return s.Body
}

func (s *ListBizUnitsResponse) SetHeaders(v map[string]*string) *ListBizUnitsResponse {
	s.Headers = v
	return s
}

func (s *ListBizUnitsResponse) SetStatusCode(v int32) *ListBizUnitsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBizUnitsResponse) SetBody(v *ListBizUnitsResponseBody) *ListBizUnitsResponse {
	s.Body = v
	return s
}

func (s *ListBizUnitsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
