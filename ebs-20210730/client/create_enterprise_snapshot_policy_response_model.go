// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnterpriseSnapshotPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEnterpriseSnapshotPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEnterpriseSnapshotPolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateEnterpriseSnapshotPolicyResponseBody) *CreateEnterpriseSnapshotPolicyResponse
	GetBody() *CreateEnterpriseSnapshotPolicyResponseBody
}

type CreateEnterpriseSnapshotPolicyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEnterpriseSnapshotPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEnterpriseSnapshotPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEnterpriseSnapshotPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseSnapshotPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEnterpriseSnapshotPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEnterpriseSnapshotPolicyResponse) GetBody() *CreateEnterpriseSnapshotPolicyResponseBody {
	return s.Body
}

func (s *CreateEnterpriseSnapshotPolicyResponse) SetHeaders(v map[string]*string) *CreateEnterpriseSnapshotPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyResponse) SetStatusCode(v int32) *CreateEnterpriseSnapshotPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyResponse) SetBody(v *CreateEnterpriseSnapshotPolicyResponseBody) *CreateEnterpriseSnapshotPolicyResponse {
	s.Body = v
	return s
}

func (s *CreateEnterpriseSnapshotPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
