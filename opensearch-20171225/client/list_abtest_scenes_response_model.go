// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListABTestScenesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListABTestScenesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListABTestScenesResponse
	GetStatusCode() *int32
	SetBody(v *ListABTestScenesResponseBody) *ListABTestScenesResponse
	GetBody() *ListABTestScenesResponseBody
}

type ListABTestScenesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListABTestScenesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListABTestScenesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListABTestScenesResponse) GoString() string {
	return s.String()
}

func (s *ListABTestScenesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListABTestScenesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListABTestScenesResponse) GetBody() *ListABTestScenesResponseBody {
	return s.Body
}

func (s *ListABTestScenesResponse) SetHeaders(v map[string]*string) *ListABTestScenesResponse {
	s.Headers = v
	return s
}

func (s *ListABTestScenesResponse) SetStatusCode(v int32) *ListABTestScenesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListABTestScenesResponse) SetBody(v *ListABTestScenesResponseBody) *ListABTestScenesResponse {
	s.Body = v
	return s
}

func (s *ListABTestScenesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
