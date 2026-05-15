// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferCallToSkillGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TransferCallToSkillGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TransferCallToSkillGroupResponse
	GetStatusCode() *int32
	SetBody(v *TransferCallToSkillGroupResponseBody) *TransferCallToSkillGroupResponse
	GetBody() *TransferCallToSkillGroupResponseBody
}

type TransferCallToSkillGroupResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TransferCallToSkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TransferCallToSkillGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s TransferCallToSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *TransferCallToSkillGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TransferCallToSkillGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TransferCallToSkillGroupResponse) GetBody() *TransferCallToSkillGroupResponseBody {
	return s.Body
}

func (s *TransferCallToSkillGroupResponse) SetHeaders(v map[string]*string) *TransferCallToSkillGroupResponse {
	s.Headers = v
	return s
}

func (s *TransferCallToSkillGroupResponse) SetStatusCode(v int32) *TransferCallToSkillGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *TransferCallToSkillGroupResponse) SetBody(v *TransferCallToSkillGroupResponseBody) *TransferCallToSkillGroupResponse {
	s.Body = v
	return s
}

func (s *TransferCallToSkillGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
