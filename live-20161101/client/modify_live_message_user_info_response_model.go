// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLiveMessageUserInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyLiveMessageUserInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyLiveMessageUserInfoResponse
	GetStatusCode() *int32
	SetBody(v *ModifyLiveMessageUserInfoResponseBody) *ModifyLiveMessageUserInfoResponse
	GetBody() *ModifyLiveMessageUserInfoResponseBody
}

type ModifyLiveMessageUserInfoResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyLiveMessageUserInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyLiveMessageUserInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyLiveMessageUserInfoResponse) GoString() string {
	return s.String()
}

func (s *ModifyLiveMessageUserInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyLiveMessageUserInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyLiveMessageUserInfoResponse) GetBody() *ModifyLiveMessageUserInfoResponseBody {
	return s.Body
}

func (s *ModifyLiveMessageUserInfoResponse) SetHeaders(v map[string]*string) *ModifyLiveMessageUserInfoResponse {
	s.Headers = v
	return s
}

func (s *ModifyLiveMessageUserInfoResponse) SetStatusCode(v int32) *ModifyLiveMessageUserInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLiveMessageUserInfoResponse) SetBody(v *ModifyLiveMessageUserInfoResponseBody) *ModifyLiveMessageUserInfoResponse {
	s.Body = v
	return s
}

func (s *ModifyLiveMessageUserInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
