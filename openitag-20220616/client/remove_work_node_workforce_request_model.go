// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveWorkNodeWorkforceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserIds(v []*int64) *RemoveWorkNodeWorkforceRequest
	GetUserIds() []*int64
}

type RemoveWorkNodeWorkforceRequest struct {
	// User IDs.
	UserIds []*int64 `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s RemoveWorkNodeWorkforceRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveWorkNodeWorkforceRequest) GoString() string {
	return s.String()
}

func (s *RemoveWorkNodeWorkforceRequest) GetUserIds() []*int64 {
	return s.UserIds
}

func (s *RemoveWorkNodeWorkforceRequest) SetUserIds(v []*int64) *RemoveWorkNodeWorkforceRequest {
	s.UserIds = v
	return s
}

func (s *RemoveWorkNodeWorkforceRequest) Validate() error {
	return dara.Validate(s)
}
