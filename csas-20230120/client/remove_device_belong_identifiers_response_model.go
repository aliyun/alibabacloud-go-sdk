// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDeviceBelongIdentifiersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveDeviceBelongIdentifiersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveDeviceBelongIdentifiersResponse
	GetStatusCode() *int32
	SetBody(v *RemoveDeviceBelongIdentifiersResponseBody) *RemoveDeviceBelongIdentifiersResponse
	GetBody() *RemoveDeviceBelongIdentifiersResponseBody
}

type RemoveDeviceBelongIdentifiersResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveDeviceBelongIdentifiersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveDeviceBelongIdentifiersResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveDeviceBelongIdentifiersResponse) GoString() string {
	return s.String()
}

func (s *RemoveDeviceBelongIdentifiersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveDeviceBelongIdentifiersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveDeviceBelongIdentifiersResponse) GetBody() *RemoveDeviceBelongIdentifiersResponseBody {
	return s.Body
}

func (s *RemoveDeviceBelongIdentifiersResponse) SetHeaders(v map[string]*string) *RemoveDeviceBelongIdentifiersResponse {
	s.Headers = v
	return s
}

func (s *RemoveDeviceBelongIdentifiersResponse) SetStatusCode(v int32) *RemoveDeviceBelongIdentifiersResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveDeviceBelongIdentifiersResponse) SetBody(v *RemoveDeviceBelongIdentifiersResponseBody) *RemoveDeviceBelongIdentifiersResponse {
	s.Body = v
	return s
}

func (s *RemoveDeviceBelongIdentifiersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
