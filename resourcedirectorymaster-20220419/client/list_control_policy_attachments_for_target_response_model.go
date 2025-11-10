// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListControlPolicyAttachmentsForTargetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListControlPolicyAttachmentsForTargetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListControlPolicyAttachmentsForTargetResponse
	GetStatusCode() *int32
	SetBody(v *ListControlPolicyAttachmentsForTargetResponseBody) *ListControlPolicyAttachmentsForTargetResponse
	GetBody() *ListControlPolicyAttachmentsForTargetResponseBody
}

type ListControlPolicyAttachmentsForTargetResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListControlPolicyAttachmentsForTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListControlPolicyAttachmentsForTargetResponse) String() string {
	return dara.Prettify(s)
}

func (s ListControlPolicyAttachmentsForTargetResponse) GoString() string {
	return s.String()
}

func (s *ListControlPolicyAttachmentsForTargetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListControlPolicyAttachmentsForTargetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListControlPolicyAttachmentsForTargetResponse) GetBody() *ListControlPolicyAttachmentsForTargetResponseBody {
	return s.Body
}

func (s *ListControlPolicyAttachmentsForTargetResponse) SetHeaders(v map[string]*string) *ListControlPolicyAttachmentsForTargetResponse {
	s.Headers = v
	return s
}

func (s *ListControlPolicyAttachmentsForTargetResponse) SetStatusCode(v int32) *ListControlPolicyAttachmentsForTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *ListControlPolicyAttachmentsForTargetResponse) SetBody(v *ListControlPolicyAttachmentsForTargetResponseBody) *ListControlPolicyAttachmentsForTargetResponse {
	s.Body = v
	return s
}

func (s *ListControlPolicyAttachmentsForTargetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
