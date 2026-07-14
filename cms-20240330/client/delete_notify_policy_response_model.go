// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNotifyPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNotifyPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNotifyPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNotifyPolicyResponseBody) *DeleteNotifyPolicyResponse
	GetBody() *DeleteNotifyPolicyResponseBody
}

type DeleteNotifyPolicyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNotifyPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNotifyPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNotifyPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteNotifyPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNotifyPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNotifyPolicyResponse) GetBody() *DeleteNotifyPolicyResponseBody {
	return s.Body
}

func (s *DeleteNotifyPolicyResponse) SetHeaders(v map[string]*string) *DeleteNotifyPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteNotifyPolicyResponse) SetStatusCode(v int32) *DeleteNotifyPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNotifyPolicyResponse) SetBody(v *DeleteNotifyPolicyResponseBody) *DeleteNotifyPolicyResponse {
	s.Body = v
	return s
}

func (s *DeleteNotifyPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
