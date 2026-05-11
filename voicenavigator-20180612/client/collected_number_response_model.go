// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCollectedNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CollectedNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CollectedNumberResponse
	GetStatusCode() *int32
	SetBody(v *CollectedNumberResponseBody) *CollectedNumberResponse
	GetBody() *CollectedNumberResponseBody
}

type CollectedNumberResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CollectedNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CollectedNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s CollectedNumberResponse) GoString() string {
	return s.String()
}

func (s *CollectedNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CollectedNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CollectedNumberResponse) GetBody() *CollectedNumberResponseBody {
	return s.Body
}

func (s *CollectedNumberResponse) SetHeaders(v map[string]*string) *CollectedNumberResponse {
	s.Headers = v
	return s
}

func (s *CollectedNumberResponse) SetStatusCode(v int32) *CollectedNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *CollectedNumberResponse) SetBody(v *CollectedNumberResponseBody) *CollectedNumberResponse {
	s.Body = v
	return s
}

func (s *CollectedNumberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
