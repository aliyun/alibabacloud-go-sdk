// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyObjectGroupOperationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *ModifyObjectGroupOperationRequest
	GetComment() *string
	SetDirection(v string) *ModifyObjectGroupOperationRequest
	GetDirection() *string
	SetLang(v string) *ModifyObjectGroupOperationRequest
	GetLang() *string
	SetObjectList(v []*string) *ModifyObjectGroupOperationRequest
	GetObjectList() []*string
	SetObjectOperation(v string) *ModifyObjectGroupOperationRequest
	GetObjectOperation() *string
	SetObjectType(v string) *ModifyObjectGroupOperationRequest
	GetObjectType() *string
	SetSourceIp(v string) *ModifyObjectGroupOperationRequest
	GetSourceIp() *string
}

type ModifyObjectGroupOperationRequest struct {
	// The remarks of the operation.
	//
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The direction of the traffic to which the access control policy applies.
	//
	// Valid values:
	//
	// 	- **in**: inbound.
	//
	// 	- **out**: outbound.
	//
	// This parameter is required.
	//
	// example:
	//
	// out
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// 	- **zh*	- (default)
	//
	// 	- **en**
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The operation objects.
	//
	// This parameter is required.
	ObjectList []*string `json:"ObjectList,omitempty" xml:"ObjectList,omitempty" type:"Repeated"`
	// The operation. Valid values:
	//
	// 	- **ignore**: adds the operation object to the whitelist.
	//
	// 	- **cancelIgnore**: removes the operation object from the whitelist.
	//
	// 	- **subscribe**: follows the operation object.
	//
	// 	- **unsubscribe**: unfollows the operation object.
	//
	// This parameter is required.
	//
	// example:
	//
	// ignore
	ObjectOperation *string `json:"ObjectOperation,omitempty" xml:"ObjectOperation,omitempty"`
	// The type of the operation object.
	//
	// Valid values:
	//
	// 	- **assetsIp**: the asset IP address.
	//
	// 	- **destinationIp**: the destination IP address.
	//
	// 	- **destinationPort**: the destination port.
	//
	// 	- **destinationDomain**: the destination domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// destinationDomain
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 123.xxx.251.60
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s ModifyObjectGroupOperationRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyObjectGroupOperationRequest) GoString() string {
	return s.String()
}

func (s *ModifyObjectGroupOperationRequest) GetComment() *string {
	return s.Comment
}

func (s *ModifyObjectGroupOperationRequest) GetDirection() *string {
	return s.Direction
}

func (s *ModifyObjectGroupOperationRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyObjectGroupOperationRequest) GetObjectList() []*string {
	return s.ObjectList
}

func (s *ModifyObjectGroupOperationRequest) GetObjectOperation() *string {
	return s.ObjectOperation
}

func (s *ModifyObjectGroupOperationRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *ModifyObjectGroupOperationRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ModifyObjectGroupOperationRequest) SetComment(v string) *ModifyObjectGroupOperationRequest {
	s.Comment = &v
	return s
}

func (s *ModifyObjectGroupOperationRequest) SetDirection(v string) *ModifyObjectGroupOperationRequest {
	s.Direction = &v
	return s
}

func (s *ModifyObjectGroupOperationRequest) SetLang(v string) *ModifyObjectGroupOperationRequest {
	s.Lang = &v
	return s
}

func (s *ModifyObjectGroupOperationRequest) SetObjectList(v []*string) *ModifyObjectGroupOperationRequest {
	s.ObjectList = v
	return s
}

func (s *ModifyObjectGroupOperationRequest) SetObjectOperation(v string) *ModifyObjectGroupOperationRequest {
	s.ObjectOperation = &v
	return s
}

func (s *ModifyObjectGroupOperationRequest) SetObjectType(v string) *ModifyObjectGroupOperationRequest {
	s.ObjectType = &v
	return s
}

func (s *ModifyObjectGroupOperationRequest) SetSourceIp(v string) *ModifyObjectGroupOperationRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyObjectGroupOperationRequest) Validate() error {
	return dara.Validate(s)
}
