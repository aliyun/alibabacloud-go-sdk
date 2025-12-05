// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPtsSceneRunningStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPtsSceneRunningStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPtsSceneRunningStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetPtsSceneRunningStatusResponseBody) *GetPtsSceneRunningStatusResponse
	GetBody() *GetPtsSceneRunningStatusResponseBody
}

type GetPtsSceneRunningStatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPtsSceneRunningStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPtsSceneRunningStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPtsSceneRunningStatusResponse) GoString() string {
	return s.String()
}

func (s *GetPtsSceneRunningStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPtsSceneRunningStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPtsSceneRunningStatusResponse) GetBody() *GetPtsSceneRunningStatusResponseBody {
	return s.Body
}

func (s *GetPtsSceneRunningStatusResponse) SetHeaders(v map[string]*string) *GetPtsSceneRunningStatusResponse {
	s.Headers = v
	return s
}

func (s *GetPtsSceneRunningStatusResponse) SetStatusCode(v int32) *GetPtsSceneRunningStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPtsSceneRunningStatusResponse) SetBody(v *GetPtsSceneRunningStatusResponseBody) *GetPtsSceneRunningStatusResponse {
	s.Body = v
	return s
}

func (s *GetPtsSceneRunningStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
