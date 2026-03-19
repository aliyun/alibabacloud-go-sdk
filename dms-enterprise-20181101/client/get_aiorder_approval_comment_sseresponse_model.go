// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAIOrderApprovalCommentSSEResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAIOrderApprovalCommentSSEResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAIOrderApprovalCommentSSEResponse
	GetStatusCode() *int32
	SetId(v string) *GetAIOrderApprovalCommentSSEResponse
	GetId() *string
	SetEvent(v string) *GetAIOrderApprovalCommentSSEResponse
	GetEvent() *string
	SetBody(v *GetAIOrderApprovalCommentSSEResponseBody) *GetAIOrderApprovalCommentSSEResponse
	GetBody() *GetAIOrderApprovalCommentSSEResponseBody
}

type GetAIOrderApprovalCommentSSEResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                                   `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                                   `json:"event,omitempty" xml:"event,omitempty"`
	Body       *GetAIOrderApprovalCommentSSEResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAIOrderApprovalCommentSSEResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAIOrderApprovalCommentSSEResponse) GoString() string {
	return s.String()
}

func (s *GetAIOrderApprovalCommentSSEResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAIOrderApprovalCommentSSEResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAIOrderApprovalCommentSSEResponse) GetId() *string {
	return s.Id
}

func (s *GetAIOrderApprovalCommentSSEResponse) GetEvent() *string {
	return s.Event
}

func (s *GetAIOrderApprovalCommentSSEResponse) GetBody() *GetAIOrderApprovalCommentSSEResponseBody {
	return s.Body
}

func (s *GetAIOrderApprovalCommentSSEResponse) SetHeaders(v map[string]*string) *GetAIOrderApprovalCommentSSEResponse {
	s.Headers = v
	return s
}

func (s *GetAIOrderApprovalCommentSSEResponse) SetStatusCode(v int32) *GetAIOrderApprovalCommentSSEResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAIOrderApprovalCommentSSEResponse) SetId(v string) *GetAIOrderApprovalCommentSSEResponse {
	s.Id = &v
	return s
}

func (s *GetAIOrderApprovalCommentSSEResponse) SetEvent(v string) *GetAIOrderApprovalCommentSSEResponse {
	s.Event = &v
	return s
}

func (s *GetAIOrderApprovalCommentSSEResponse) SetBody(v *GetAIOrderApprovalCommentSSEResponseBody) *GetAIOrderApprovalCommentSSEResponse {
	s.Body = v
	return s
}

func (s *GetAIOrderApprovalCommentSSEResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
