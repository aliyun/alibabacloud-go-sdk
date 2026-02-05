// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceApplyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListServiceApplyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListServiceApplyResponse
	GetStatusCode() *int32
	SetBody(v *ListServiceApplyResponseBody) *ListServiceApplyResponse
	GetBody() *ListServiceApplyResponseBody
}

type ListServiceApplyResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceApplyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceApplyResponse) String() string {
	return dara.Prettify(s)
}

func (s ListServiceApplyResponse) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListServiceApplyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListServiceApplyResponse) GetBody() *ListServiceApplyResponseBody {
	return s.Body
}

func (s *ListServiceApplyResponse) SetHeaders(v map[string]*string) *ListServiceApplyResponse {
	s.Headers = v
	return s
}

func (s *ListServiceApplyResponse) SetStatusCode(v int32) *ListServiceApplyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceApplyResponse) SetBody(v *ListServiceApplyResponseBody) *ListServiceApplyResponse {
	s.Body = v
	return s
}

func (s *ListServiceApplyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
