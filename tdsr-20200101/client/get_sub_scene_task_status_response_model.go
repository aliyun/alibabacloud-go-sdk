// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubSceneTaskStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSubSceneTaskStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSubSceneTaskStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetSubSceneTaskStatusResponseBody) *GetSubSceneTaskStatusResponse
	GetBody() *GetSubSceneTaskStatusResponseBody
}

type GetSubSceneTaskStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSubSceneTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSubSceneTaskStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSubSceneTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *GetSubSceneTaskStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSubSceneTaskStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSubSceneTaskStatusResponse) GetBody() *GetSubSceneTaskStatusResponseBody {
	return s.Body
}

func (s *GetSubSceneTaskStatusResponse) SetHeaders(v map[string]*string) *GetSubSceneTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *GetSubSceneTaskStatusResponse) SetStatusCode(v int32) *GetSubSceneTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSubSceneTaskStatusResponse) SetBody(v *GetSubSceneTaskStatusResponseBody) *GetSubSceneTaskStatusResponse {
	s.Body = v
	return s
}

func (s *GetSubSceneTaskStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
