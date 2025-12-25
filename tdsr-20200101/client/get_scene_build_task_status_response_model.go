// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSceneBuildTaskStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSceneBuildTaskStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSceneBuildTaskStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetSceneBuildTaskStatusResponseBody) *GetSceneBuildTaskStatusResponse
	GetBody() *GetSceneBuildTaskStatusResponseBody
}

type GetSceneBuildTaskStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSceneBuildTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSceneBuildTaskStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSceneBuildTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *GetSceneBuildTaskStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSceneBuildTaskStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSceneBuildTaskStatusResponse) GetBody() *GetSceneBuildTaskStatusResponseBody {
	return s.Body
}

func (s *GetSceneBuildTaskStatusResponse) SetHeaders(v map[string]*string) *GetSceneBuildTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *GetSceneBuildTaskStatusResponse) SetStatusCode(v int32) *GetSceneBuildTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSceneBuildTaskStatusResponse) SetBody(v *GetSceneBuildTaskStatusResponseBody) *GetSceneBuildTaskStatusResponse {
	s.Body = v
	return s
}

func (s *GetSceneBuildTaskStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
