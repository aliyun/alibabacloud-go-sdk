// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWorkNodeWorkforceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserIds(v []*int64) *AddWorkNodeWorkforceRequest
	GetUserIds() []*int64
}

type AddWorkNodeWorkforceRequest struct {
	// User List.
	UserIds []*int64 `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s AddWorkNodeWorkforceRequest) String() string {
	return dara.Prettify(s)
}

func (s AddWorkNodeWorkforceRequest) GoString() string {
	return s.String()
}

func (s *AddWorkNodeWorkforceRequest) GetUserIds() []*int64 {
	return s.UserIds
}

func (s *AddWorkNodeWorkforceRequest) SetUserIds(v []*int64) *AddWorkNodeWorkforceRequest {
	s.UserIds = v
	return s
}

func (s *AddWorkNodeWorkforceRequest) Validate() error {
	return dara.Validate(s)
}
