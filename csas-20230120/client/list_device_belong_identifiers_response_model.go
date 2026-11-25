// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceBelongIdentifiersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDeviceBelongIdentifiersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDeviceBelongIdentifiersResponse
	GetStatusCode() *int32
	SetBody(v *ListDeviceBelongIdentifiersResponseBody) *ListDeviceBelongIdentifiersResponse
	GetBody() *ListDeviceBelongIdentifiersResponseBody
}

type ListDeviceBelongIdentifiersResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeviceBelongIdentifiersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeviceBelongIdentifiersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceBelongIdentifiersResponse) GoString() string {
	return s.String()
}

func (s *ListDeviceBelongIdentifiersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDeviceBelongIdentifiersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDeviceBelongIdentifiersResponse) GetBody() *ListDeviceBelongIdentifiersResponseBody {
	return s.Body
}

func (s *ListDeviceBelongIdentifiersResponse) SetHeaders(v map[string]*string) *ListDeviceBelongIdentifiersResponse {
	s.Headers = v
	return s
}

func (s *ListDeviceBelongIdentifiersResponse) SetStatusCode(v int32) *ListDeviceBelongIdentifiersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeviceBelongIdentifiersResponse) SetBody(v *ListDeviceBelongIdentifiersResponseBody) *ListDeviceBelongIdentifiersResponse {
	s.Body = v
	return s
}

func (s *ListDeviceBelongIdentifiersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
