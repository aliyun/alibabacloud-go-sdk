// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubAccountInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSubAccountInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSubAccountInfoResponse
	GetStatusCode() *int32
	SetBody(v *ListSubAccountInfoResult) *ListSubAccountInfoResponse
	GetBody() *ListSubAccountInfoResult
}

type ListSubAccountInfoResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSubAccountInfoResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSubAccountInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSubAccountInfoResponse) GoString() string {
	return s.String()
}

func (s *ListSubAccountInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSubAccountInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSubAccountInfoResponse) GetBody() *ListSubAccountInfoResult {
	return s.Body
}

func (s *ListSubAccountInfoResponse) SetHeaders(v map[string]*string) *ListSubAccountInfoResponse {
	s.Headers = v
	return s
}

func (s *ListSubAccountInfoResponse) SetStatusCode(v int32) *ListSubAccountInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSubAccountInfoResponse) SetBody(v *ListSubAccountInfoResult) *ListSubAccountInfoResponse {
	s.Body = v
	return s
}

func (s *ListSubAccountInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
