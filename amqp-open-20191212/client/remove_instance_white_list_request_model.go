// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveInstanceWhiteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *RemoveInstanceWhiteListRequest
	GetInstanceId() *string
	SetWhiteListItemId(v int64) *RemoveInstanceWhiteListRequest
	GetWhiteListItemId() *int64
	SetWhiteListType(v int32) *RemoveInstanceWhiteListRequest
	GetWhiteListType() *int32
}

type RemoveInstanceWhiteListRequest struct {
	// The ID of the instance from which to remove a whitelist entry.
	//
	// This parameter is required.
	//
	// example:
	//
	// amqp-cn-5yd3aw******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Call the ListInstanceWhiteList operation to obtain this ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 420
	WhiteListItemId *int64 `json:"whiteListItemId,omitempty" xml:"whiteListItemId,omitempty"`
	// The type of the whitelist item. Specify 2 for an IP address or 1 for a VPC ID.
	//
	// A VPC whitelist only applies to instances with the anytunnel VPC endpoint type. Newer instances use PrivateLink for their VPC endpoints and do not support VPC whitelists.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	WhiteListType *int32 `json:"whiteListType,omitempty" xml:"whiteListType,omitempty"`
}

func (s RemoveInstanceWhiteListRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveInstanceWhiteListRequest) GoString() string {
	return s.String()
}

func (s *RemoveInstanceWhiteListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RemoveInstanceWhiteListRequest) GetWhiteListItemId() *int64 {
	return s.WhiteListItemId
}

func (s *RemoveInstanceWhiteListRequest) GetWhiteListType() *int32 {
	return s.WhiteListType
}

func (s *RemoveInstanceWhiteListRequest) SetInstanceId(v string) *RemoveInstanceWhiteListRequest {
	s.InstanceId = &v
	return s
}

func (s *RemoveInstanceWhiteListRequest) SetWhiteListItemId(v int64) *RemoveInstanceWhiteListRequest {
	s.WhiteListItemId = &v
	return s
}

func (s *RemoveInstanceWhiteListRequest) SetWhiteListType(v int32) *RemoveInstanceWhiteListRequest {
	s.WhiteListType = &v
	return s
}

func (s *RemoveInstanceWhiteListRequest) Validate() error {
	return dara.Validate(s)
}
