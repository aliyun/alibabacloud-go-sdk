// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserPropertyValueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessChannel(v string) *DeleteUserPropertyValueRequest
	GetBusinessChannel() *string
	SetPropertyId(v int64) *DeleteUserPropertyValueRequest
	GetPropertyId() *int64
	SetPropertyValueId(v int64) *DeleteUserPropertyValueRequest
	GetPropertyValueId() *int64
	SetUserId(v int64) *DeleteUserPropertyValueRequest
	GetUserId() *int64
}

type DeleteUserPropertyValueRequest struct {
	// The channel.
	//
	// example:
	//
	// ENTERPRISE
	BusinessChannel *string `json:"BusinessChannel,omitempty" xml:"BusinessChannel,omitempty"`
	// The property ID. You can call [ListProperty](~~ListProperty~~) to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 390
	PropertyId *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
	// The property value ID. You can call [ListProperty](~~ListProperty~~) to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 978
	PropertyValueId *int64 `json:"PropertyValueId,omitempty" xml:"PropertyValueId,omitempty"`
	// The user ID. You can call [DescribeUsers](~~DescribeUsers~~) to obtain this parameter, which corresponds to the value of the `Id` response parameter.
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

func (s *DeleteUserPropertyValueRequest) GetBusinessChannel() *string {
	return s.BusinessChannel
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

func (s *DeleteUserPropertyValueRequest) SetBusinessChannel(v string) *DeleteUserPropertyValueRequest {
	s.BusinessChannel = &v
	return s
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
