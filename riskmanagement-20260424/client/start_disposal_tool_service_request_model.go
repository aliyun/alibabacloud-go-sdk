// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDisposalToolServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthType(v string) *StartDisposalToolServiceRequest
	GetAuthType() *string
}

type StartDisposalToolServiceRequest struct {
	// example:
	//
	// DisposalTool
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
}

func (s StartDisposalToolServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s StartDisposalToolServiceRequest) GoString() string {
	return s.String()
}

func (s *StartDisposalToolServiceRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *StartDisposalToolServiceRequest) SetAuthType(v string) *StartDisposalToolServiceRequest {
	s.AuthType = &v
	return s
}

func (s *StartDisposalToolServiceRequest) Validate() error {
	return dara.Validate(s)
}
