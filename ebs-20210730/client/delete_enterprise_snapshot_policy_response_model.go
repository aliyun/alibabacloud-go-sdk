// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnterpriseSnapshotPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEnterpriseSnapshotPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEnterpriseSnapshotPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEnterpriseSnapshotPolicyResponseBody) *DeleteEnterpriseSnapshotPolicyResponse
	GetBody() *DeleteEnterpriseSnapshotPolicyResponseBody
}

type DeleteEnterpriseSnapshotPolicyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEnterpriseSnapshotPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEnterpriseSnapshotPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnterpriseSnapshotPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteEnterpriseSnapshotPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEnterpriseSnapshotPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEnterpriseSnapshotPolicyResponse) GetBody() *DeleteEnterpriseSnapshotPolicyResponseBody {
	return s.Body
}

func (s *DeleteEnterpriseSnapshotPolicyResponse) SetHeaders(v map[string]*string) *DeleteEnterpriseSnapshotPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteEnterpriseSnapshotPolicyResponse) SetStatusCode(v int32) *DeleteEnterpriseSnapshotPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEnterpriseSnapshotPolicyResponse) SetBody(v *DeleteEnterpriseSnapshotPolicyResponseBody) *DeleteEnterpriseSnapshotPolicyResponse {
	s.Body = v
	return s
}

func (s *DeleteEnterpriseSnapshotPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
