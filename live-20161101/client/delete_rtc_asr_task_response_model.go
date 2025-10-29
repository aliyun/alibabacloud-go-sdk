// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRtcAsrTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRtcAsrTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRtcAsrTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRtcAsrTaskResponseBody) *DeleteRtcAsrTaskResponse
	GetBody() *DeleteRtcAsrTaskResponseBody
}

type DeleteRtcAsrTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRtcAsrTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRtcAsrTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRtcAsrTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteRtcAsrTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRtcAsrTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRtcAsrTaskResponse) GetBody() *DeleteRtcAsrTaskResponseBody {
	return s.Body
}

func (s *DeleteRtcAsrTaskResponse) SetHeaders(v map[string]*string) *DeleteRtcAsrTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteRtcAsrTaskResponse) SetStatusCode(v int32) *DeleteRtcAsrTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRtcAsrTaskResponse) SetBody(v *DeleteRtcAsrTaskResponseBody) *DeleteRtcAsrTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteRtcAsrTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
