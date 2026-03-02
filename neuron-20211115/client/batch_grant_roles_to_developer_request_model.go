// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGrantRolesToDeveloperRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *BatchGrantRolesCmd) *BatchGrantRolesToDeveloperRequest
	GetBody() *BatchGrantRolesCmd
}

type BatchGrantRolesToDeveloperRequest struct {
	// This parameter is required.
	Body *BatchGrantRolesCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchGrantRolesToDeveloperRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchGrantRolesToDeveloperRequest) GoString() string {
	return s.String()
}

func (s *BatchGrantRolesToDeveloperRequest) GetBody() *BatchGrantRolesCmd {
	return s.Body
}

func (s *BatchGrantRolesToDeveloperRequest) SetBody(v *BatchGrantRolesCmd) *BatchGrantRolesToDeveloperRequest {
	s.Body = v
	return s
}

func (s *BatchGrantRolesToDeveloperRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
