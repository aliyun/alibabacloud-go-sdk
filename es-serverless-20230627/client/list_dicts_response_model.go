// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDictsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDictsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDictsResponse
	GetStatusCode() *int32
	SetBody(v *ListDictsResponseBody) *ListDictsResponse
	GetBody() *ListDictsResponseBody
}

type ListDictsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDictsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDictsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDictsResponse) GoString() string {
	return s.String()
}

func (s *ListDictsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDictsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDictsResponse) GetBody() *ListDictsResponseBody {
	return s.Body
}

func (s *ListDictsResponse) SetHeaders(v map[string]*string) *ListDictsResponse {
	s.Headers = v
	return s
}

func (s *ListDictsResponse) SetStatusCode(v int32) *ListDictsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDictsResponse) SetBody(v *ListDictsResponseBody) *ListDictsResponse {
	s.Body = v
	return s
}

func (s *ListDictsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
