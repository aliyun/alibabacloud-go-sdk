// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserPropertyValueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPropertyId(v int64) *DeleteUserPropertyValueRequest
	GetPropertyId() *int64
	SetPropertyValueId(v int64) *DeleteUserPropertyValueRequest
	GetPropertyValueId() *int64
	SetUserId(v int64) *DeleteUserPropertyValueRequest
	GetUserId() *int64
}

type DeleteUserPropertyValueRequest struct {
	// The property ID. You can call the [ListProperty](~~ListProperty~~) operation to query the property ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 390
	PropertyId *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
	// The property value ID. You can call the [ListProperty](~~ListProperty~~) operation to query the property value ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 978
	PropertyValueId *int64 `json:"PropertyValueId,omitempty" xml:"PropertyValueId,omitempty"`
	// The user ID. You can call the [DescribeUsers](~~DescribeUsers~~) operation to query the user ID, which is the return value of the `Id` parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteUserPropertyValueRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserPropertyValueRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserPropertyValueRequest) GetPropertyId() *int64 {
	return s.PropertyId
}

func (s *DeleteUserPropertyValueRequest) GetPropertyValueId() *int64 {
	return s.PropertyValueId
}

func (s *DeleteUserPropertyValueRequest) GetUserId() *int64 {
	return s.UserId
}

func (s *DeleteUserPropertyValueRequest) SetPropertyId(v int64) *DeleteUserPropertyValueRequest {
	s.PropertyId = &v
	return s
}

func (s *DeleteUserPropertyValueRequest) SetPropertyValueId(v int64) *DeleteUserPropertyValueRequest {
	s.PropertyValueId = &v
	return s
}

func (s *DeleteUserPropertyValueRequest) SetUserId(v int64) *DeleteUserPropertyValueRequest {
	s.UserId = &v
	return s
}

func (s *DeleteUserPropertyValueRequest) Validate() error {
	return dara.Validate(s)
}
