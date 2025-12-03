// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppReleaseStageExecutionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAppReleaseStageExecutionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAppReleaseStageExecutionsResponse
	GetStatusCode() *int32
	SetBody(v *ListAppReleaseStageExecutionsResponseBody) *ListAppReleaseStageExecutionsResponse
	GetBody() *ListAppReleaseStageExecutionsResponseBody
}

type ListAppReleaseStageExecutionsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppReleaseStageExecutionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppReleaseStageExecutionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAppReleaseStageExecutionsResponse) GoString() string {
	return s.String()
}

func (s *ListAppReleaseStageExecutionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAppReleaseStageExecutionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAppReleaseStageExecutionsResponse) GetBody() *ListAppReleaseStageExecutionsResponseBody {
	return s.Body
}

func (s *ListAppReleaseStageExecutionsResponse) SetHeaders(v map[string]*string) *ListAppReleaseStageExecutionsResponse {
	s.Headers = v
	return s
}

func (s *ListAppReleaseStageExecutionsResponse) SetStatusCode(v int32) *ListAppReleaseStageExecutionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppReleaseStageExecutionsResponse) SetBody(v *ListAppReleaseStageExecutionsResponseBody) *ListAppReleaseStageExecutionsResponse {
	s.Body = v
	return s
}

func (s *ListAppReleaseStageExecutionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
