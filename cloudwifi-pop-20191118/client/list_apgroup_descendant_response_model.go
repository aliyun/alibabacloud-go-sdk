// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApgroupDescendantResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApgroupDescendantResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApgroupDescendantResponse
	GetStatusCode() *int32
	SetBody(v *ListApgroupDescendantResponseBody) *ListApgroupDescendantResponse
	GetBody() *ListApgroupDescendantResponseBody
}

type ListApgroupDescendantResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApgroupDescendantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApgroupDescendantResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApgroupDescendantResponse) GoString() string {
	return s.String()
}

func (s *ListApgroupDescendantResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApgroupDescendantResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApgroupDescendantResponse) GetBody() *ListApgroupDescendantResponseBody {
	return s.Body
}

func (s *ListApgroupDescendantResponse) SetHeaders(v map[string]*string) *ListApgroupDescendantResponse {
	s.Headers = v
	return s
}

func (s *ListApgroupDescendantResponse) SetStatusCode(v int32) *ListApgroupDescendantResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApgroupDescendantResponse) SetBody(v *ListApgroupDescendantResponseBody) *ListApgroupDescendantResponse {
	s.Body = v
	return s
}

func (s *ListApgroupDescendantResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
