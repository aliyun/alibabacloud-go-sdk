// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOnlineConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v map[string]*string) *ModifyOnlineConfigRequest
	GetBody() map[string]*string
}

type ModifyOnlineConfigRequest struct {
	// The request body.
	Body map[string]*string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyOnlineConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyOnlineConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyOnlineConfigRequest) GetBody() map[string]*string {
	return s.Body
}

func (s *ModifyOnlineConfigRequest) SetBody(v map[string]*string) *ModifyOnlineConfigRequest {
	s.Body = v
	return s
}

func (s *ModifyOnlineConfigRequest) Validate() error {
	return dara.Validate(s)
}
