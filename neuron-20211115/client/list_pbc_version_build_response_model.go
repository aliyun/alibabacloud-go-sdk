// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPbcVersionBuildResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPbcVersionBuildResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPbcVersionBuildResponse
	GetStatusCode() *int32
	SetBody(v *PbcVersionListResult) *ListPbcVersionBuildResponse
	GetBody() *PbcVersionListResult
}

type ListPbcVersionBuildResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PbcVersionListResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPbcVersionBuildResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPbcVersionBuildResponse) GoString() string {
	return s.String()
}

func (s *ListPbcVersionBuildResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPbcVersionBuildResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPbcVersionBuildResponse) GetBody() *PbcVersionListResult {
	return s.Body
}

func (s *ListPbcVersionBuildResponse) SetHeaders(v map[string]*string) *ListPbcVersionBuildResponse {
	s.Headers = v
	return s
}

func (s *ListPbcVersionBuildResponse) SetStatusCode(v int32) *ListPbcVersionBuildResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPbcVersionBuildResponse) SetBody(v *PbcVersionListResult) *ListPbcVersionBuildResponse {
	s.Body = v
	return s
}

func (s *ListPbcVersionBuildResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
