// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlterShareReceiversRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddedReceivers(v []*string) *AlterShareReceiversRequest
	GetAddedReceivers() []*string
	SetRemovedReceivers(v []*string) *AlterShareReceiversRequest
	GetRemovedReceivers() []*string
}

type AlterShareReceiversRequest struct {
	AddedReceivers   []*string `json:"addedReceivers,omitempty" xml:"addedReceivers,omitempty" type:"Repeated"`
	RemovedReceivers []*string `json:"removedReceivers,omitempty" xml:"removedReceivers,omitempty" type:"Repeated"`
}

func (s AlterShareReceiversRequest) String() string {
	return dara.Prettify(s)
}

func (s AlterShareReceiversRequest) GoString() string {
	return s.String()
}

func (s *AlterShareReceiversRequest) GetAddedReceivers() []*string {
	return s.AddedReceivers
}

func (s *AlterShareReceiversRequest) GetRemovedReceivers() []*string {
	return s.RemovedReceivers
}

func (s *AlterShareReceiversRequest) SetAddedReceivers(v []*string) *AlterShareReceiversRequest {
	s.AddedReceivers = v
	return s
}

func (s *AlterShareReceiversRequest) SetRemovedReceivers(v []*string) *AlterShareReceiversRequest {
	s.RemovedReceivers = v
	return s
}

func (s *AlterShareReceiversRequest) Validate() error {
	return dara.Validate(s)
}
