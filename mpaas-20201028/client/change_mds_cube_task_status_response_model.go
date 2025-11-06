// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeMdsCubeTaskStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeMdsCubeTaskStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeMdsCubeTaskStatusResponse
	GetStatusCode() *int32
	SetBody(v *ChangeMdsCubeTaskStatusResponseBody) *ChangeMdsCubeTaskStatusResponse
	GetBody() *ChangeMdsCubeTaskStatusResponseBody
}

type ChangeMdsCubeTaskStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeMdsCubeTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeMdsCubeTaskStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeMdsCubeTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *ChangeMdsCubeTaskStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeMdsCubeTaskStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeMdsCubeTaskStatusResponse) GetBody() *ChangeMdsCubeTaskStatusResponseBody {
	return s.Body
}

func (s *ChangeMdsCubeTaskStatusResponse) SetHeaders(v map[string]*string) *ChangeMdsCubeTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *ChangeMdsCubeTaskStatusResponse) SetStatusCode(v int32) *ChangeMdsCubeTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeMdsCubeTaskStatusResponse) SetBody(v *ChangeMdsCubeTaskStatusResponseBody) *ChangeMdsCubeTaskStatusResponse {
	s.Body = v
	return s
}

func (s *ChangeMdsCubeTaskStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
