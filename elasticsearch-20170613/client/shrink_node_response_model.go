// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iShrinkNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ShrinkNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ShrinkNodeResponse
	GetStatusCode() *int32
	SetBody(v *ShrinkNodeResponseBody) *ShrinkNodeResponse
	GetBody() *ShrinkNodeResponseBody
}

type ShrinkNodeResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ShrinkNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ShrinkNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s ShrinkNodeResponse) GoString() string {
	return s.String()
}

func (s *ShrinkNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ShrinkNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ShrinkNodeResponse) GetBody() *ShrinkNodeResponseBody {
	return s.Body
}

func (s *ShrinkNodeResponse) SetHeaders(v map[string]*string) *ShrinkNodeResponse {
	s.Headers = v
	return s
}

func (s *ShrinkNodeResponse) SetStatusCode(v int32) *ShrinkNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *ShrinkNodeResponse) SetBody(v *ShrinkNodeResponseBody) *ShrinkNodeResponse {
	s.Body = v
	return s
}

func (s *ShrinkNodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
