// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveMinimapResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveMinimapResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveMinimapResponse
	GetStatusCode() *int32
	SetBody(v *SaveMinimapResponseBody) *SaveMinimapResponse
	GetBody() *SaveMinimapResponseBody
}

type SaveMinimapResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveMinimapResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveMinimapResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveMinimapResponse) GoString() string {
	return s.String()
}

func (s *SaveMinimapResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveMinimapResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveMinimapResponse) GetBody() *SaveMinimapResponseBody {
	return s.Body
}

func (s *SaveMinimapResponse) SetHeaders(v map[string]*string) *SaveMinimapResponse {
	s.Headers = v
	return s
}

func (s *SaveMinimapResponse) SetStatusCode(v int32) *SaveMinimapResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveMinimapResponse) SetBody(v *SaveMinimapResponseBody) *SaveMinimapResponse {
	s.Body = v
	return s
}

func (s *SaveMinimapResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
