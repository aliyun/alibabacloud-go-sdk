// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDisposalToolStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthType(v string) *GetDisposalToolStatusRequest
	GetAuthType() *string
}

type GetDisposalToolStatusRequest struct {
	// The authorization type.
	//
	// - **DisposalTool**: one-click disposal authorization type
	//
	// example:
	//
	// DisposalTool
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
}

func (s GetDisposalToolStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDisposalToolStatusRequest) GoString() string {
	return s.String()
}

func (s *GetDisposalToolStatusRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *GetDisposalToolStatusRequest) SetAuthType(v string) *GetDisposalToolStatusRequest {
	s.AuthType = &v
	return s
}

func (s *GetDisposalToolStatusRequest) Validate() error {
	return dara.Validate(s)
}
