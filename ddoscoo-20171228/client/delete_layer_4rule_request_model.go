// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLayer4RuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListeners(v string) *DeleteLayer4RuleRequest
	GetListeners() *string
}

type DeleteLayer4RuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// {"InstanceId":"0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc","Protocol":"tcp","FrontendPort":80}
	Listeners *string `json:"Listeners,omitempty" xml:"Listeners,omitempty"`
}

func (s DeleteLayer4RuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLayer4RuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteLayer4RuleRequest) GetListeners() *string {
	return s.Listeners
}

func (s *DeleteLayer4RuleRequest) SetListeners(v string) *DeleteLayer4RuleRequest {
	s.Listeners = &v
	return s
}

func (s *DeleteLayer4RuleRequest) Validate() error {
	return dara.Validate(s)
}
