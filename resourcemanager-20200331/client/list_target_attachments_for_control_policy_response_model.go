// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTargetAttachmentsForControlPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTargetAttachmentsForControlPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTargetAttachmentsForControlPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ListTargetAttachmentsForControlPolicyResponseBody) *ListTargetAttachmentsForControlPolicyResponse
	GetBody() *ListTargetAttachmentsForControlPolicyResponseBody
}

type ListTargetAttachmentsForControlPolicyResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTargetAttachmentsForControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTargetAttachmentsForControlPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTargetAttachmentsForControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *ListTargetAttachmentsForControlPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTargetAttachmentsForControlPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTargetAttachmentsForControlPolicyResponse) GetBody() *ListTargetAttachmentsForControlPolicyResponseBody {
	return s.Body
}

func (s *ListTargetAttachmentsForControlPolicyResponse) SetHeaders(v map[string]*string) *ListTargetAttachmentsForControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyResponse) SetStatusCode(v int32) *ListTargetAttachmentsForControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyResponse) SetBody(v *ListTargetAttachmentsForControlPolicyResponseBody) *ListTargetAttachmentsForControlPolicyResponse {
	s.Body = v
	return s
}

func (s *ListTargetAttachmentsForControlPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
