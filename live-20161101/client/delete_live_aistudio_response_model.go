// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveAIStudioResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLiveAIStudioResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLiveAIStudioResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLiveAIStudioResponseBody) *DeleteLiveAIStudioResponse
	GetBody() *DeleteLiveAIStudioResponseBody
}

type DeleteLiveAIStudioResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveAIStudioResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLiveAIStudioResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveAIStudioResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveAIStudioResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLiveAIStudioResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLiveAIStudioResponse) GetBody() *DeleteLiveAIStudioResponseBody {
	return s.Body
}

func (s *DeleteLiveAIStudioResponse) SetHeaders(v map[string]*string) *DeleteLiveAIStudioResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveAIStudioResponse) SetStatusCode(v int32) *DeleteLiveAIStudioResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveAIStudioResponse) SetBody(v *DeleteLiveAIStudioResponseBody) *DeleteLiveAIStudioResponse {
	s.Body = v
	return s
}

func (s *DeleteLiveAIStudioResponse) Validate() error {
	return dara.Validate(s)
}
