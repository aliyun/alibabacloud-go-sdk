// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCopySceneTaskStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCopySceneTaskStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCopySceneTaskStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetCopySceneTaskStatusResponseBody) *GetCopySceneTaskStatusResponse
	GetBody() *GetCopySceneTaskStatusResponseBody
}

type GetCopySceneTaskStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCopySceneTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCopySceneTaskStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCopySceneTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *GetCopySceneTaskStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCopySceneTaskStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCopySceneTaskStatusResponse) GetBody() *GetCopySceneTaskStatusResponseBody {
	return s.Body
}

func (s *GetCopySceneTaskStatusResponse) SetHeaders(v map[string]*string) *GetCopySceneTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *GetCopySceneTaskStatusResponse) SetStatusCode(v int32) *GetCopySceneTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCopySceneTaskStatusResponse) SetBody(v *GetCopySceneTaskStatusResponseBody) *GetCopySceneTaskStatusResponse {
	s.Body = v
	return s
}

func (s *GetCopySceneTaskStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
