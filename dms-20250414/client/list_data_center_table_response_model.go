// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataCenterTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataCenterTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataCenterTableResponse
	GetStatusCode() *int32
	SetBody(v *ListDataCenterTableResponseBody) *ListDataCenterTableResponse
	GetBody() *ListDataCenterTableResponseBody
}

type ListDataCenterTableResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataCenterTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataCenterTableResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataCenterTableResponse) GoString() string {
	return s.String()
}

func (s *ListDataCenterTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataCenterTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataCenterTableResponse) GetBody() *ListDataCenterTableResponseBody {
	return s.Body
}

func (s *ListDataCenterTableResponse) SetHeaders(v map[string]*string) *ListDataCenterTableResponse {
	s.Headers = v
	return s
}

func (s *ListDataCenterTableResponse) SetStatusCode(v int32) *ListDataCenterTableResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataCenterTableResponse) SetBody(v *ListDataCenterTableResponseBody) *ListDataCenterTableResponse {
	s.Body = v
	return s
}

func (s *ListDataCenterTableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
