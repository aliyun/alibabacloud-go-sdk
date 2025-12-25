// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPackSceneTaskStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPackSceneTaskStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPackSceneTaskStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetPackSceneTaskStatusResponseBody) *GetPackSceneTaskStatusResponse
	GetBody() *GetPackSceneTaskStatusResponseBody
}

type GetPackSceneTaskStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPackSceneTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPackSceneTaskStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPackSceneTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *GetPackSceneTaskStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPackSceneTaskStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPackSceneTaskStatusResponse) GetBody() *GetPackSceneTaskStatusResponseBody {
	return s.Body
}

func (s *GetPackSceneTaskStatusResponse) SetHeaders(v map[string]*string) *GetPackSceneTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *GetPackSceneTaskStatusResponse) SetStatusCode(v int32) *GetPackSceneTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPackSceneTaskStatusResponse) SetBody(v *GetPackSceneTaskStatusResponseBody) *GetPackSceneTaskStatusResponse {
	s.Body = v
	return s
}

func (s *GetPackSceneTaskStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
