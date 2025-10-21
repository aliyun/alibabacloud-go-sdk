// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModelInstanceIdList(v []*int64) *DeleteModelInstanceRequest
	GetModelInstanceIdList() []*int64
}

type DeleteModelInstanceRequest struct {
	ModelInstanceIdList []*int64 `json:"ModelInstanceIdList,omitempty" xml:"ModelInstanceIdList,omitempty" type:"Repeated"`
}

func (s DeleteModelInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteModelInstanceRequest) GetModelInstanceIdList() []*int64 {
	return s.ModelInstanceIdList
}

func (s *DeleteModelInstanceRequest) SetModelInstanceIdList(v []*int64) *DeleteModelInstanceRequest {
	s.ModelInstanceIdList = v
	return s
}

func (s *DeleteModelInstanceRequest) Validate() error {
	return dara.Validate(s)
}
