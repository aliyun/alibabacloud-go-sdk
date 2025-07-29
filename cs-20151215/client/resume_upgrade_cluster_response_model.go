// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeUpgradeClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResumeUpgradeClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResumeUpgradeClusterResponse
	GetStatusCode() *int32
}

type ResumeUpgradeClusterResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s ResumeUpgradeClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s ResumeUpgradeClusterResponse) GoString() string {
	return s.String()
}

func (s *ResumeUpgradeClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResumeUpgradeClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResumeUpgradeClusterResponse) SetHeaders(v map[string]*string) *ResumeUpgradeClusterResponse {
	s.Headers = v
	return s
}

func (s *ResumeUpgradeClusterResponse) SetStatusCode(v int32) *ResumeUpgradeClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeUpgradeClusterResponse) Validate() error {
	return dara.Validate(s)
}
