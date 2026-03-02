// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddOrQuitPdpLaneForServiceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *AddOrQuitLaneCmd) *AddOrQuitPdpLaneForServiceGroupRequest
	GetBody() *AddOrQuitLaneCmd
}

type AddOrQuitPdpLaneForServiceGroupRequest struct {
	Body *AddOrQuitLaneCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddOrQuitPdpLaneForServiceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s AddOrQuitPdpLaneForServiceGroupRequest) GoString() string {
	return s.String()
}

func (s *AddOrQuitPdpLaneForServiceGroupRequest) GetBody() *AddOrQuitLaneCmd {
	return s.Body
}

func (s *AddOrQuitPdpLaneForServiceGroupRequest) SetBody(v *AddOrQuitLaneCmd) *AddOrQuitPdpLaneForServiceGroupRequest {
	s.Body = v
	return s
}

func (s *AddOrQuitPdpLaneForServiceGroupRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
