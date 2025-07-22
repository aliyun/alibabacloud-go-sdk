// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEndpointSwitchTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEndpointSwitchTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEndpointSwitchTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetEndpointSwitchTaskResponseBody) *GetEndpointSwitchTaskResponse
	GetBody() *GetEndpointSwitchTaskResponseBody
}

type GetEndpointSwitchTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEndpointSwitchTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEndpointSwitchTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEndpointSwitchTaskResponse) GoString() string {
	return s.String()
}

func (s *GetEndpointSwitchTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEndpointSwitchTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEndpointSwitchTaskResponse) GetBody() *GetEndpointSwitchTaskResponseBody {
	return s.Body
}

func (s *GetEndpointSwitchTaskResponse) SetHeaders(v map[string]*string) *GetEndpointSwitchTaskResponse {
	s.Headers = v
	return s
}

func (s *GetEndpointSwitchTaskResponse) SetStatusCode(v int32) *GetEndpointSwitchTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEndpointSwitchTaskResponse) SetBody(v *GetEndpointSwitchTaskResponseBody) *GetEndpointSwitchTaskResponse {
	s.Body = v
	return s
}

func (s *GetEndpointSwitchTaskResponse) Validate() error {
	return dara.Validate(s)
}
