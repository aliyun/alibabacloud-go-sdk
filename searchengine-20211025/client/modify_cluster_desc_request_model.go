// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterDescRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v map[string]interface{}) *ModifyClusterDescRequest
	GetBody() map[string]interface{}
}

type ModifyClusterDescRequest struct {
	// The request body.
	//
	// example:
	//
	// {}
	Body map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyClusterDescRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterDescRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterDescRequest) GetBody() map[string]interface{} {
	return s.Body
}

func (s *ModifyClusterDescRequest) SetBody(v map[string]interface{}) *ModifyClusterDescRequest {
	s.Body = v
	return s
}

func (s *ModifyClusterDescRequest) Validate() error {
	return dara.Validate(s)
}
