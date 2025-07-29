// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeComponentUpgradeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResumeComponentUpgradeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResumeComponentUpgradeResponse
	GetStatusCode() *int32
}

type ResumeComponentUpgradeResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s ResumeComponentUpgradeResponse) String() string {
	return dara.Prettify(s)
}

func (s ResumeComponentUpgradeResponse) GoString() string {
	return s.String()
}

func (s *ResumeComponentUpgradeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResumeComponentUpgradeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResumeComponentUpgradeResponse) SetHeaders(v map[string]*string) *ResumeComponentUpgradeResponse {
	s.Headers = v
	return s
}

func (s *ResumeComponentUpgradeResponse) SetStatusCode(v int32) *ResumeComponentUpgradeResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeComponentUpgradeResponse) Validate() error {
	return dara.Validate(s)
}
