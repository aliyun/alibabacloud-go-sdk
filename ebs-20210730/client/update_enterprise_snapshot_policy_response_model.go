// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEnterpriseSnapshotPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEnterpriseSnapshotPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEnterpriseSnapshotPolicyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEnterpriseSnapshotPolicyResponseBody) *UpdateEnterpriseSnapshotPolicyResponse
	GetBody() *UpdateEnterpriseSnapshotPolicyResponseBody
}

type UpdateEnterpriseSnapshotPolicyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEnterpriseSnapshotPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEnterpriseSnapshotPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnterpriseSnapshotPolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateEnterpriseSnapshotPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEnterpriseSnapshotPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEnterpriseSnapshotPolicyResponse) GetBody() *UpdateEnterpriseSnapshotPolicyResponseBody {
	return s.Body
}

func (s *UpdateEnterpriseSnapshotPolicyResponse) SetHeaders(v map[string]*string) *UpdateEnterpriseSnapshotPolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyResponse) SetStatusCode(v int32) *UpdateEnterpriseSnapshotPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyResponse) SetBody(v *UpdateEnterpriseSnapshotPolicyResponseBody) *UpdateEnterpriseSnapshotPolicyResponse {
	s.Body = v
	return s
}

func (s *UpdateEnterpriseSnapshotPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
