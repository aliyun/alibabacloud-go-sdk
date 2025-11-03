// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveSourcesFromTrafficMirrorSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveSourcesFromTrafficMirrorSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveSourcesFromTrafficMirrorSessionResponse
	GetStatusCode() *int32
	SetBody(v *RemoveSourcesFromTrafficMirrorSessionResponseBody) *RemoveSourcesFromTrafficMirrorSessionResponse
	GetBody() *RemoveSourcesFromTrafficMirrorSessionResponseBody
}

type RemoveSourcesFromTrafficMirrorSessionResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveSourcesFromTrafficMirrorSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveSourcesFromTrafficMirrorSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveSourcesFromTrafficMirrorSessionResponse) GoString() string {
	return s.String()
}

func (s *RemoveSourcesFromTrafficMirrorSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveSourcesFromTrafficMirrorSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveSourcesFromTrafficMirrorSessionResponse) GetBody() *RemoveSourcesFromTrafficMirrorSessionResponseBody {
	return s.Body
}

func (s *RemoveSourcesFromTrafficMirrorSessionResponse) SetHeaders(v map[string]*string) *RemoveSourcesFromTrafficMirrorSessionResponse {
	s.Headers = v
	return s
}

func (s *RemoveSourcesFromTrafficMirrorSessionResponse) SetStatusCode(v int32) *RemoveSourcesFromTrafficMirrorSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveSourcesFromTrafficMirrorSessionResponse) SetBody(v *RemoveSourcesFromTrafficMirrorSessionResponseBody) *RemoveSourcesFromTrafficMirrorSessionResponse {
	s.Body = v
	return s
}

func (s *RemoveSourcesFromTrafficMirrorSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
