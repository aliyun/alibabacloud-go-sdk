// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListABTestExperimentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListABTestExperimentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListABTestExperimentsResponse
	GetStatusCode() *int32
	SetBody(v *ListABTestExperimentsResponseBody) *ListABTestExperimentsResponse
	GetBody() *ListABTestExperimentsResponseBody
}

type ListABTestExperimentsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListABTestExperimentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListABTestExperimentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListABTestExperimentsResponse) GoString() string {
	return s.String()
}

func (s *ListABTestExperimentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListABTestExperimentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListABTestExperimentsResponse) GetBody() *ListABTestExperimentsResponseBody {
	return s.Body
}

func (s *ListABTestExperimentsResponse) SetHeaders(v map[string]*string) *ListABTestExperimentsResponse {
	s.Headers = v
	return s
}

func (s *ListABTestExperimentsResponse) SetStatusCode(v int32) *ListABTestExperimentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListABTestExperimentsResponse) SetBody(v *ListABTestExperimentsResponseBody) *ListABTestExperimentsResponse {
	s.Body = v
	return s
}

func (s *ListABTestExperimentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
