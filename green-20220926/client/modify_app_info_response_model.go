// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAppInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAppInfoResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAppInfoResponseBody) *ModifyAppInfoResponse
	GetBody() *ModifyAppInfoResponseBody
}

type ModifyAppInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAppInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAppInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppInfoResponse) GoString() string {
	return s.String()
}

func (s *ModifyAppInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAppInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAppInfoResponse) GetBody() *ModifyAppInfoResponseBody {
	return s.Body
}

func (s *ModifyAppInfoResponse) SetHeaders(v map[string]*string) *ModifyAppInfoResponse {
	s.Headers = v
	return s
}

func (s *ModifyAppInfoResponse) SetStatusCode(v int32) *ModifyAppInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAppInfoResponse) SetBody(v *ModifyAppInfoResponseBody) *ModifyAppInfoResponse {
	s.Body = v
	return s
}

func (s *ModifyAppInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
