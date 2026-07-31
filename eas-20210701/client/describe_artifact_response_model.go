// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeArtifactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeArtifactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeArtifactResponse
	GetStatusCode() *int32
	SetBody(v *Artifact) *DescribeArtifactResponse
	GetBody() *Artifact
}

type DescribeArtifactResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Artifact          `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeArtifactResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeArtifactResponse) GoString() string {
	return s.String()
}

func (s *DescribeArtifactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeArtifactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeArtifactResponse) GetBody() *Artifact {
	return s.Body
}

func (s *DescribeArtifactResponse) SetHeaders(v map[string]*string) *DescribeArtifactResponse {
	s.Headers = v
	return s
}

func (s *DescribeArtifactResponse) SetStatusCode(v int32) *DescribeArtifactResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeArtifactResponse) SetBody(v *Artifact) *DescribeArtifactResponse {
	s.Body = v
	return s
}

func (s *DescribeArtifactResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
