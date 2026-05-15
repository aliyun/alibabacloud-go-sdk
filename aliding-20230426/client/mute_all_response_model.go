// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMuteAllResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MuteAllResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MuteAllResponse
	GetStatusCode() *int32
	SetBody(v *MuteAllResponseBody) *MuteAllResponse
	GetBody() *MuteAllResponseBody
}

type MuteAllResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MuteAllResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MuteAllResponse) String() string {
	return dara.Prettify(s)
}

func (s MuteAllResponse) GoString() string {
	return s.String()
}

func (s *MuteAllResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MuteAllResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MuteAllResponse) GetBody() *MuteAllResponseBody {
	return s.Body
}

func (s *MuteAllResponse) SetHeaders(v map[string]*string) *MuteAllResponse {
	s.Headers = v
	return s
}

func (s *MuteAllResponse) SetStatusCode(v int32) *MuteAllResponse {
	s.StatusCode = &v
	return s
}

func (s *MuteAllResponse) SetBody(v *MuteAllResponseBody) *MuteAllResponse {
	s.Body = v
	return s
}

func (s *MuteAllResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
