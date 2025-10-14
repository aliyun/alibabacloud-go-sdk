// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataSetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataSetsResponse
	GetStatusCode() *int32
	SetBody(v *ListDataSetsResponseBody) *ListDataSetsResponse
	GetBody() *ListDataSetsResponseBody
}

type ListDataSetsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataSetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataSetsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataSetsResponse) GoString() string {
	return s.String()
}

func (s *ListDataSetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataSetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataSetsResponse) GetBody() *ListDataSetsResponseBody {
	return s.Body
}

func (s *ListDataSetsResponse) SetHeaders(v map[string]*string) *ListDataSetsResponse {
	s.Headers = v
	return s
}

func (s *ListDataSetsResponse) SetStatusCode(v int32) *ListDataSetsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataSetsResponse) SetBody(v *ListDataSetsResponseBody) *ListDataSetsResponse {
	s.Body = v
	return s
}

func (s *ListDataSetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
