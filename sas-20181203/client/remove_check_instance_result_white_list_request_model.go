// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveCheckInstanceResultWhiteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckId(v int64) *RemoveCheckInstanceResultWhiteListRequest
	GetCheckId() *int64
	SetInstanceIds(v []*string) *RemoveCheckInstanceResultWhiteListRequest
	GetInstanceIds() []*string
}

type RemoveCheckInstanceResultWhiteListRequest struct {
	// The ID of the check item.
	//
	// example:
	//
	// 11
	CheckId *int64 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The IDs of instances. Separate multiple IDs with commas (,).
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s RemoveCheckInstanceResultWhiteListRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveCheckInstanceResultWhiteListRequest) GoString() string {
	return s.String()
}

func (s *RemoveCheckInstanceResultWhiteListRequest) GetCheckId() *int64 {
	return s.CheckId
}

func (s *RemoveCheckInstanceResultWhiteListRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *RemoveCheckInstanceResultWhiteListRequest) SetCheckId(v int64) *RemoveCheckInstanceResultWhiteListRequest {
	s.CheckId = &v
	return s
}

func (s *RemoveCheckInstanceResultWhiteListRequest) SetInstanceIds(v []*string) *RemoveCheckInstanceResultWhiteListRequest {
	s.InstanceIds = v
	return s
}

func (s *RemoveCheckInstanceResultWhiteListRequest) Validate() error {
	return dara.Validate(s)
}
