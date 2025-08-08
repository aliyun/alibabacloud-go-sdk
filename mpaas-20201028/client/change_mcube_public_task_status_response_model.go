// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeMcubePublicTaskStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeMcubePublicTaskStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeMcubePublicTaskStatusResponse
	GetStatusCode() *int32
	SetBody(v *ChangeMcubePublicTaskStatusResponseBody) *ChangeMcubePublicTaskStatusResponse
	GetBody() *ChangeMcubePublicTaskStatusResponseBody
}

type ChangeMcubePublicTaskStatusResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeMcubePublicTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeMcubePublicTaskStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeMcubePublicTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *ChangeMcubePublicTaskStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeMcubePublicTaskStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeMcubePublicTaskStatusResponse) GetBody() *ChangeMcubePublicTaskStatusResponseBody {
	return s.Body
}

func (s *ChangeMcubePublicTaskStatusResponse) SetHeaders(v map[string]*string) *ChangeMcubePublicTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *ChangeMcubePublicTaskStatusResponse) SetStatusCode(v int32) *ChangeMcubePublicTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeMcubePublicTaskStatusResponse) SetBody(v *ChangeMcubePublicTaskStatusResponseBody) *ChangeMcubePublicTaskStatusResponse {
	s.Body = v
	return s
}

func (s *ChangeMcubePublicTaskStatusResponse) Validate() error {
	return dara.Validate(s)
}
