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
	// The remark for the operation.
	//
	// example:
	//
	// Remark for batch whitelisting
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The traffic direction controlled by the access control policy. Valid values:
	//
	// - **in**: Inbound traffic.
	//
	// - **out**: Outbound traffic.
	//
	// This parameter is required.
	//
	// example:
	//
	// out
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The language of the response message. Valid values:
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The list of objects to operate on.
	//
	// This parameter is required.
	ObjectList []*string `json:"ObjectList,omitempty" xml:"ObjectList,omitempty" type:"Repeated"`
	// Specifies the operation to perform on the object, such as adding to the whitelist, removing from the whitelist, following, or unfollowing. Valid values:
	//
	// - **subscribe**: Follow.
	//
	// - **unsubscribe**: Unfollow.
	//
	// - **ignore**: Add to the whitelist.
	//
	// - **cancelIgnore**: Remove from the whitelist.
	//
	// This parameter is required.
	//
	// example:
	//
	// ignore
	ObjectOperation *string `json:"ObjectOperation,omitempty" xml:"ObjectOperation,omitempty"`
	// The type of the object to add to the outbound connections whitelist or follow. Valid values:
	//
	// - **assetsIp**: Asset IP address.
	//
	// - **destinationIp**: Destination IP address.
	//
	// - **destinationPort**: Destination port.
	//
	// - **destinationDomain**: Destination domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// destinationDomain
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The source IP address of the requester.
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
