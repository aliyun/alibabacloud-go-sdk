// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScriptRecordingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListScriptRecordingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListScriptRecordingResponse
	GetStatusCode() *int32
	SetBody(v *ListScriptRecordingResponseBody) *ListScriptRecordingResponse
	GetBody() *ListScriptRecordingResponseBody
}

type ListScriptRecordingResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListScriptRecordingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListScriptRecordingResponse) String() string {
	return dara.Prettify(s)
}

func (s ListScriptRecordingResponse) GoString() string {
	return s.String()
}

func (s *ListScriptRecordingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListScriptRecordingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListScriptRecordingResponse) GetBody() *ListScriptRecordingResponseBody {
	return s.Body
}

func (s *ListScriptRecordingResponse) SetHeaders(v map[string]*string) *ListScriptRecordingResponse {
	s.Headers = v
	return s
}

func (s *ListScriptRecordingResponse) SetStatusCode(v int32) *ListScriptRecordingResponse {
	s.StatusCode = &v
	return s
}

func (s *ListScriptRecordingResponse) SetBody(v *ListScriptRecordingResponseBody) *ListScriptRecordingResponse {
	s.Body = v
	return s
}

func (s *ListScriptRecordingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
