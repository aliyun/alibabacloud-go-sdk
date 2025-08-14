// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTranscodeJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParentJobId(v string) *GetTranscodeJobRequest
	GetParentJobId() *string
}

type GetTranscodeJobRequest struct {
	// The job ID.
	//
	// example:
	//
	// ab4802364a2e49208c99efab82dfa8e8
	ParentJobId *string `json:"ParentJobId,omitempty" xml:"ParentJobId,omitempty"`
}

func (s GetTranscodeJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobRequest) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobRequest) GetParentJobId() *string {
	return s.ParentJobId
}

func (s *GetTranscodeJobRequest) SetParentJobId(v string) *GetTranscodeJobRequest {
	s.ParentJobId = &v
	return s
}

func (s *GetTranscodeJobRequest) Validate() error {
	return dara.Validate(s)
}
