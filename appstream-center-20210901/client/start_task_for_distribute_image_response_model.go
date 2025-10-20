// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartTaskForDistributeImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartTaskForDistributeImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartTaskForDistributeImageResponse
	GetStatusCode() *int32
	SetBody(v *StartTaskForDistributeImageResponseBody) *StartTaskForDistributeImageResponse
	GetBody() *StartTaskForDistributeImageResponseBody
}

type StartTaskForDistributeImageResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartTaskForDistributeImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartTaskForDistributeImageResponse) String() string {
	return dara.Prettify(s)
}

func (s StartTaskForDistributeImageResponse) GoString() string {
	return s.String()
}

func (s *StartTaskForDistributeImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartTaskForDistributeImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartTaskForDistributeImageResponse) GetBody() *StartTaskForDistributeImageResponseBody {
	return s.Body
}

func (s *StartTaskForDistributeImageResponse) SetHeaders(v map[string]*string) *StartTaskForDistributeImageResponse {
	s.Headers = v
	return s
}

func (s *StartTaskForDistributeImageResponse) SetStatusCode(v int32) *StartTaskForDistributeImageResponse {
	s.StatusCode = &v
	return s
}

func (s *StartTaskForDistributeImageResponse) SetBody(v *StartTaskForDistributeImageResponseBody) *StartTaskForDistributeImageResponse {
	s.Body = v
	return s
}

func (s *StartTaskForDistributeImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
