// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDingTalkStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDingTalkStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDingTalkStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDingTalkStatusResponseBody) *ModifyDingTalkStatusResponse
	GetBody() *ModifyDingTalkStatusResponseBody
}

type ModifyDingTalkStatusResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDingTalkStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDingTalkStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDingTalkStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyDingTalkStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDingTalkStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDingTalkStatusResponse) GetBody() *ModifyDingTalkStatusResponseBody {
	return s.Body
}

func (s *ModifyDingTalkStatusResponse) SetHeaders(v map[string]*string) *ModifyDingTalkStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyDingTalkStatusResponse) SetStatusCode(v int32) *ModifyDingTalkStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDingTalkStatusResponse) SetBody(v *ModifyDingTalkStatusResponseBody) *ModifyDingTalkStatusResponse {
	s.Body = v
	return s
}

func (s *ModifyDingTalkStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
