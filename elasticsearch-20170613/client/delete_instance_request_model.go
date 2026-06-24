// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteInstanceRequest
	GetClientToken() *string
	SetDeleteType(v string) *DeleteInstanceRequest
	GetDeleteType() *string
}

type DeleteInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The release type. Valid values:
	//
	// - immediate: The instance is immediately deleted. After deletion, the system permanently clears all data, and the instance no longer appears in the instance list.
	//
	// - protective: The instance is frozen for 24 hours before the data is permanently cleared. During this period, the instance still appears in the instance list, and you can choose to [restore the instance](https://help.aliyun.com/document_detail/202195.html) or [immediately release it](https://help.aliyun.com/document_detail/202195.html).
	//
	// example:
	//
	// protective
	DeleteType *string `json:"deleteType,omitempty" xml:"deleteType,omitempty"`
}

func (s DeleteInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteInstanceRequest) GetDeleteType() *string {
	return s.DeleteType
}

func (s *DeleteInstanceRequest) SetClientToken(v string) *DeleteInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteInstanceRequest) SetDeleteType(v string) *DeleteInstanceRequest {
	s.DeleteType = &v
	return s
}

func (s *DeleteInstanceRequest) Validate() error {
	return dara.Validate(s)
}
