// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDebugCollectedNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DebugCollectedNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DebugCollectedNumberResponse
	GetStatusCode() *int32
	SetBody(v *DebugCollectedNumberResponseBody) *DebugCollectedNumberResponse
	GetBody() *DebugCollectedNumberResponseBody
}

type DebugCollectedNumberResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DebugCollectedNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DebugCollectedNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s DebugCollectedNumberResponse) GoString() string {
	return s.String()
}

func (s *DebugCollectedNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DebugCollectedNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DebugCollectedNumberResponse) GetBody() *DebugCollectedNumberResponseBody {
	return s.Body
}

func (s *DebugCollectedNumberResponse) SetHeaders(v map[string]*string) *DebugCollectedNumberResponse {
	s.Headers = v
	return s
}

func (s *DebugCollectedNumberResponse) SetStatusCode(v int32) *DebugCollectedNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *DebugCollectedNumberResponse) SetBody(v *DebugCollectedNumberResponseBody) *DebugCollectedNumberResponse {
	s.Body = v
	return s
}

func (s *DebugCollectedNumberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
