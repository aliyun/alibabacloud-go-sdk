// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransformClusterMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TransformClusterMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TransformClusterMemberResponse
	GetStatusCode() *int32
	SetBody(v *TransformClusterMemberResponseBody) *TransformClusterMemberResponse
	GetBody() *TransformClusterMemberResponseBody
}

type TransformClusterMemberResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TransformClusterMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TransformClusterMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s TransformClusterMemberResponse) GoString() string {
	return s.String()
}

func (s *TransformClusterMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TransformClusterMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TransformClusterMemberResponse) GetBody() *TransformClusterMemberResponseBody {
	return s.Body
}

func (s *TransformClusterMemberResponse) SetHeaders(v map[string]*string) *TransformClusterMemberResponse {
	s.Headers = v
	return s
}

func (s *TransformClusterMemberResponse) SetStatusCode(v int32) *TransformClusterMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *TransformClusterMemberResponse) SetBody(v *TransformClusterMemberResponseBody) *TransformClusterMemberResponse {
	s.Body = v
	return s
}

func (s *TransformClusterMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
