// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootARMServerInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetServerId(v string) *RebootARMServerInstanceRequest
	GetServerId() *string
}

type RebootARMServerInstanceRequest struct {
	// The ID of the server.
	//
	// example:
	//
	// yourInstance ID
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
}

func (s RebootARMServerInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RebootARMServerInstanceRequest) GoString() string {
	return s.String()
}

func (s *RebootARMServerInstanceRequest) GetServerId() *string {
	return s.ServerId
}

func (s *RebootARMServerInstanceRequest) SetServerId(v string) *RebootARMServerInstanceRequest {
	s.ServerId = &v
	return s
}

func (s *RebootARMServerInstanceRequest) Validate() error {
	return dara.Validate(s)
}
