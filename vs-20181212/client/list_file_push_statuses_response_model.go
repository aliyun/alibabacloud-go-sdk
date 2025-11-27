// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFilePushStatusesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFilePushStatusesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFilePushStatusesResponse
	GetStatusCode() *int32
	SetBody(v *ListFilePushStatusesResponseBody) *ListFilePushStatusesResponse
	GetBody() *ListFilePushStatusesResponseBody
}

type ListFilePushStatusesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFilePushStatusesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFilePushStatusesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFilePushStatusesResponse) GoString() string {
	return s.String()
}

func (s *ListFilePushStatusesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFilePushStatusesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFilePushStatusesResponse) GetBody() *ListFilePushStatusesResponseBody {
	return s.Body
}

func (s *ListFilePushStatusesResponse) SetHeaders(v map[string]*string) *ListFilePushStatusesResponse {
	s.Headers = v
	return s
}

func (s *ListFilePushStatusesResponse) SetStatusCode(v int32) *ListFilePushStatusesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFilePushStatusesResponse) SetBody(v *ListFilePushStatusesResponseBody) *ListFilePushStatusesResponse {
	s.Body = v
	return s
}

func (s *ListFilePushStatusesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
