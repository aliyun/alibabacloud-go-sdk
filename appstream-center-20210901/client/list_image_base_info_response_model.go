// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageBaseInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListImageBaseInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListImageBaseInfoResponse
	GetStatusCode() *int32
	SetBody(v *ListImageBaseInfoResponseBody) *ListImageBaseInfoResponse
	GetBody() *ListImageBaseInfoResponseBody
}

type ListImageBaseInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListImageBaseInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListImageBaseInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ListImageBaseInfoResponse) GoString() string {
	return s.String()
}

func (s *ListImageBaseInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListImageBaseInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListImageBaseInfoResponse) GetBody() *ListImageBaseInfoResponseBody {
	return s.Body
}

func (s *ListImageBaseInfoResponse) SetHeaders(v map[string]*string) *ListImageBaseInfoResponse {
	s.Headers = v
	return s
}

func (s *ListImageBaseInfoResponse) SetStatusCode(v int32) *ListImageBaseInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListImageBaseInfoResponse) SetBody(v *ListImageBaseInfoResponseBody) *ListImageBaseInfoResponse {
	s.Body = v
	return s
}

func (s *ListImageBaseInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
