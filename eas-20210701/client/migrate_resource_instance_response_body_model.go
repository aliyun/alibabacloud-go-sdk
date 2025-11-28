// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateResourceInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *MigrateResourceInstanceResponseBody
	GetInstanceIds() []*string
	SetMessage(v string) *MigrateResourceInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *MigrateResourceInstanceResponseBody
	GetRequestId() *string
}

type MigrateResourceInstanceResponseBody struct {
	// The instance ID.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MigrateResourceInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MigrateResourceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *MigrateResourceInstanceResponseBody) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *MigrateResourceInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MigrateResourceInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MigrateResourceInstanceResponseBody) SetInstanceIds(v []*string) *MigrateResourceInstanceResponseBody {
	s.InstanceIds = v
	return s
}

func (s *MigrateResourceInstanceResponseBody) SetMessage(v string) *MigrateResourceInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *MigrateResourceInstanceResponseBody) SetRequestId(v string) *MigrateResourceInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *MigrateResourceInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
