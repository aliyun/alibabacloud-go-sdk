// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveEnvRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvId(v string) *RemoveEnvRequest
	GetEnvId() *string
}

type RemoveEnvRequest struct {
	// The ID of the environment that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10YPA8H
	EnvId *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
}

func (s RemoveEnvRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveEnvRequest) GoString() string {
	return s.String()
}

func (s *RemoveEnvRequest) GetEnvId() *string {
	return s.EnvId
}

func (s *RemoveEnvRequest) SetEnvId(v string) *RemoveEnvRequest {
	s.EnvId = &v
	return s
}

func (s *RemoveEnvRequest) Validate() error {
	return dara.Validate(s)
}
