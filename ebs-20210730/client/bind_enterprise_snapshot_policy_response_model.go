// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindEnterpriseSnapshotPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindEnterpriseSnapshotPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindEnterpriseSnapshotPolicyResponse
	GetStatusCode() *int32
	SetBody(v *BindEnterpriseSnapshotPolicyResponseBody) *BindEnterpriseSnapshotPolicyResponse
	GetBody() *BindEnterpriseSnapshotPolicyResponseBody
}

type BindEnterpriseSnapshotPolicyResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindEnterpriseSnapshotPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindEnterpriseSnapshotPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s BindEnterpriseSnapshotPolicyResponse) GoString() string {
	return s.String()
}

func (s *BindEnterpriseSnapshotPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindEnterpriseSnapshotPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindEnterpriseSnapshotPolicyResponse) GetBody() *BindEnterpriseSnapshotPolicyResponseBody {
	return s.Body
}

func (s *BindEnterpriseSnapshotPolicyResponse) SetHeaders(v map[string]*string) *BindEnterpriseSnapshotPolicyResponse {
	s.Headers = v
	return s
}

func (s *BindEnterpriseSnapshotPolicyResponse) SetStatusCode(v int32) *BindEnterpriseSnapshotPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *BindEnterpriseSnapshotPolicyResponse) SetBody(v *BindEnterpriseSnapshotPolicyResponseBody) *BindEnterpriseSnapshotPolicyResponse {
	s.Body = v
	return s
}

func (s *BindEnterpriseSnapshotPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
