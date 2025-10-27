// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMachineAppsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMachineAppsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMachineAppsResponse
	GetStatusCode() *int32
	SetBody(v *ListMachineAppsResponseBody) *ListMachineAppsResponse
	GetBody() *ListMachineAppsResponseBody
}

type ListMachineAppsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMachineAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMachineAppsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMachineAppsResponse) GoString() string {
	return s.String()
}

func (s *ListMachineAppsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMachineAppsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMachineAppsResponse) GetBody() *ListMachineAppsResponseBody {
	return s.Body
}

func (s *ListMachineAppsResponse) SetHeaders(v map[string]*string) *ListMachineAppsResponse {
	s.Headers = v
	return s
}

func (s *ListMachineAppsResponse) SetStatusCode(v int32) *ListMachineAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMachineAppsResponse) SetBody(v *ListMachineAppsResponseBody) *ListMachineAppsResponse {
	s.Body = v
	return s
}

func (s *ListMachineAppsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
