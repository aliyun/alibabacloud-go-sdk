// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartPairDrillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartPairDrillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartPairDrillResponse
	GetStatusCode() *int32
	SetBody(v *StartPairDrillResponseBody) *StartPairDrillResponse
	GetBody() *StartPairDrillResponseBody
}

type StartPairDrillResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartPairDrillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartPairDrillResponse) String() string {
	return dara.Prettify(s)
}

func (s StartPairDrillResponse) GoString() string {
	return s.String()
}

func (s *StartPairDrillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartPairDrillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartPairDrillResponse) GetBody() *StartPairDrillResponseBody {
	return s.Body
}

func (s *StartPairDrillResponse) SetHeaders(v map[string]*string) *StartPairDrillResponse {
	s.Headers = v
	return s
}

func (s *StartPairDrillResponse) SetStatusCode(v int32) *StartPairDrillResponse {
	s.StatusCode = &v
	return s
}

func (s *StartPairDrillResponse) SetBody(v *StartPairDrillResponseBody) *StartPairDrillResponse {
	s.Body = v
	return s
}

func (s *StartPairDrillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
