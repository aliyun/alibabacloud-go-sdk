// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseARMServerInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ReleaseARMServerInstanceRequest
	GetInstanceId() *string
}

type ReleaseARMServerInstanceRequest struct {
	// The ID of the server.
	//
	// This parameter is required.
	//
	// example:
	//
	// yourInstance ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ReleaseARMServerInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseARMServerInstanceRequest) GoString() string {
	return s.String()
}

func (s *ReleaseARMServerInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReleaseARMServerInstanceRequest) SetInstanceId(v string) *ReleaseARMServerInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ReleaseARMServerInstanceRequest) Validate() error {
	return dara.Validate(s)
}
