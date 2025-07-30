// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetUserPropertyValueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPropertyId(v int64) *SetUserPropertyValueRequest
	GetPropertyId() *int64
	SetPropertyValueId(v int64) *SetUserPropertyValueRequest
	GetPropertyValueId() *int64
	SetUserId(v int64) *SetUserPropertyValueRequest
	GetUserId() *int64
	SetUserName(v string) *SetUserPropertyValueRequest
	GetUserName() *string
}

type SetUserPropertyValueRequest struct {
	// The property ID. You can call the [ListProperty](~~ListProperty~~) operation to query the property ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 390
	PropertyId *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
	// The ID of the property value. You can call the [ListProperty](~~ListProperty~~) operation to query the ID of the property value.
	//
	// This parameter is required.
	//
	// example:
	//
	// 978
	PropertyValueId *int64 `json:"PropertyValueId,omitempty" xml:"PropertyValueId,omitempty"`
	// The ID of the convenience user. You can call the [DescribeUsers](~~DescribeUsers~~) operation to query the user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username of the convenience user.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s SetUserPropertyValueRequest) String() string {
	return dara.Prettify(s)
}

func (s SetUserPropertyValueRequest) GoString() string {
	return s.String()
}

func (s *SetUserPropertyValueRequest) GetPropertyId() *int64 {
	return s.PropertyId
}

func (s *SetUserPropertyValueRequest) GetPropertyValueId() *int64 {
	return s.PropertyValueId
}

func (s *SetUserPropertyValueRequest) GetUserId() *int64 {
	return s.UserId
}

func (s *SetUserPropertyValueRequest) GetUserName() *string {
	return s.UserName
}

func (s *SetUserPropertyValueRequest) SetPropertyId(v int64) *SetUserPropertyValueRequest {
	s.PropertyId = &v
	return s
}

func (s *SetUserPropertyValueRequest) SetPropertyValueId(v int64) *SetUserPropertyValueRequest {
	s.PropertyValueId = &v
	return s
}

func (s *SetUserPropertyValueRequest) SetUserId(v int64) *SetUserPropertyValueRequest {
	s.UserId = &v
	return s
}

func (s *SetUserPropertyValueRequest) SetUserName(v string) *SetUserPropertyValueRequest {
	s.UserName = &v
	return s
}

func (s *SetUserPropertyValueRequest) Validate() error {
	return dara.Validate(s)
}
