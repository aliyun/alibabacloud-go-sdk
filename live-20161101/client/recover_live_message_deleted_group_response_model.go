// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoverLiveMessageDeletedGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecoverLiveMessageDeletedGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecoverLiveMessageDeletedGroupResponse
	GetStatusCode() *int32
	SetBody(v *RecoverLiveMessageDeletedGroupResponseBody) *RecoverLiveMessageDeletedGroupResponse
	GetBody() *RecoverLiveMessageDeletedGroupResponseBody
}

type RecoverLiveMessageDeletedGroupResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecoverLiveMessageDeletedGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecoverLiveMessageDeletedGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s RecoverLiveMessageDeletedGroupResponse) GoString() string {
	return s.String()
}

func (s *RecoverLiveMessageDeletedGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecoverLiveMessageDeletedGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecoverLiveMessageDeletedGroupResponse) GetBody() *RecoverLiveMessageDeletedGroupResponseBody {
	return s.Body
}

func (s *RecoverLiveMessageDeletedGroupResponse) SetHeaders(v map[string]*string) *RecoverLiveMessageDeletedGroupResponse {
	s.Headers = v
	return s
}

func (s *RecoverLiveMessageDeletedGroupResponse) SetStatusCode(v int32) *RecoverLiveMessageDeletedGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RecoverLiveMessageDeletedGroupResponse) SetBody(v *RecoverLiveMessageDeletedGroupResponseBody) *RecoverLiveMessageDeletedGroupResponse {
	s.Body = v
	return s
}

func (s *RecoverLiveMessageDeletedGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
