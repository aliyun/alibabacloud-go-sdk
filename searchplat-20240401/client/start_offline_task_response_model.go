// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartOfflineTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartOfflineTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartOfflineTaskResponse
	GetStatusCode() *int32
	SetBody(v *StartOfflineTaskResponseBody) *StartOfflineTaskResponse
	GetBody() *StartOfflineTaskResponseBody
}

type StartOfflineTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartOfflineTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartOfflineTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StartOfflineTaskResponse) GoString() string {
	return s.String()
}

func (s *StartOfflineTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartOfflineTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartOfflineTaskResponse) GetBody() *StartOfflineTaskResponseBody {
	return s.Body
}

func (s *StartOfflineTaskResponse) SetHeaders(v map[string]*string) *StartOfflineTaskResponse {
	s.Headers = v
	return s
}

func (s *StartOfflineTaskResponse) SetStatusCode(v int32) *StartOfflineTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StartOfflineTaskResponse) SetBody(v *StartOfflineTaskResponseBody) *StartOfflineTaskResponse {
	s.Body = v
	return s
}

func (s *StartOfflineTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
