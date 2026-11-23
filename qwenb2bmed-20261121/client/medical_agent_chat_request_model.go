// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMedicalAgentChatRequest interface {
	dara.Model
	String() string
	GoString() string
}

type MedicalAgentChatRequest struct {
}

func (s MedicalAgentChatRequest) String() string {
	return dara.Prettify(s)
}

func (s MedicalAgentChatRequest) GoString() string {
	return s.String()
}

func (s *MedicalAgentChatRequest) Validate() error {
	return dara.Validate(s)
}
