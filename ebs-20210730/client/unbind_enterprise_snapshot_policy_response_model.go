// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindEnterpriseSnapshotPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindEnterpriseSnapshotPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindEnterpriseSnapshotPolicyResponse
	GetStatusCode() *int32
	SetBody(v *UnbindEnterpriseSnapshotPolicyResponseBody) *UnbindEnterpriseSnapshotPolicyResponse
	GetBody() *UnbindEnterpriseSnapshotPolicyResponseBody
}

type UnbindEnterpriseSnapshotPolicyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindEnterpriseSnapshotPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindEnterpriseSnapshotPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindEnterpriseSnapshotPolicyResponse) GoString() string {
	return s.String()
}

func (s *UnbindEnterpriseSnapshotPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindEnterpriseSnapshotPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindEnterpriseSnapshotPolicyResponse) GetBody() *UnbindEnterpriseSnapshotPolicyResponseBody {
	return s.Body
}

func (s *UnbindEnterpriseSnapshotPolicyResponse) SetHeaders(v map[string]*string) *UnbindEnterpriseSnapshotPolicyResponse {
	s.Headers = v
	return s
}

func (s *UnbindEnterpriseSnapshotPolicyResponse) SetStatusCode(v int32) *UnbindEnterpriseSnapshotPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindEnterpriseSnapshotPolicyResponse) SetBody(v *UnbindEnterpriseSnapshotPolicyResponseBody) *UnbindEnterpriseSnapshotPolicyResponse {
	s.Body = v
	return s
}

func (s *UnbindEnterpriseSnapshotPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
