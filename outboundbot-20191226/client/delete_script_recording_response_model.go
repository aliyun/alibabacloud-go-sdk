// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScriptRecordingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteScriptRecordingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteScriptRecordingResponse
	GetStatusCode() *int32
	SetBody(v *DeleteScriptRecordingResponseBody) *DeleteScriptRecordingResponse
	GetBody() *DeleteScriptRecordingResponseBody
}

type DeleteScriptRecordingResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteScriptRecordingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteScriptRecordingResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteScriptRecordingResponse) GoString() string {
	return s.String()
}

func (s *DeleteScriptRecordingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteScriptRecordingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteScriptRecordingResponse) GetBody() *DeleteScriptRecordingResponseBody {
	return s.Body
}

func (s *DeleteScriptRecordingResponse) SetHeaders(v map[string]*string) *DeleteScriptRecordingResponse {
	s.Headers = v
	return s
}

func (s *DeleteScriptRecordingResponse) SetStatusCode(v int32) *DeleteScriptRecordingResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteScriptRecordingResponse) SetBody(v *DeleteScriptRecordingResponseBody) *DeleteScriptRecordingResponse {
	s.Body = v
	return s
}

func (s *DeleteScriptRecordingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
