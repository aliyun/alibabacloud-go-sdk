// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCustomAuthIdentityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *QueryCustomAuthIdentityRequest
	GetClientId() *string
	SetIdentityType(v string) *QueryCustomAuthIdentityRequest
	GetIdentityType() *string
	SetInstanceId(v string) *QueryCustomAuthIdentityRequest
	GetInstanceId() *string
	SetNextToken(v string) *QueryCustomAuthIdentityRequest
	GetNextToken() *string
	SetSize(v int32) *QueryCustomAuthIdentityRequest
	GetSize() *int32
	SetUsername(v string) *QueryCustomAuthIdentityRequest
	GetUsername() *string
}

type QueryCustomAuthIdentityRequest struct {
	// The client ID of the device. You can specify this parameter when the identity type is CLIENT.
	//
	// example:
	//
	// GID_test@@@test
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The identity type.
	//
	// example:
	//
	// USER
	IdentityType *string `json:"IdentityType,omitempty" xml:"IdentityType,omitempty"`
	// The ID of the ApsaraMQ for MQTT instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// post-111****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The token that marks the end position of the current query. You can pass this token as a parameter in the next call to continue paginating. If this is the first call or the last page is returned, the value is an empty string.
	//
	// example:
	//
	// eyJhY2NvdW50IjoiMTM4MTcxODk3NDQzMjQ1OSIsImV2ZW50SWQiOiJGMkUxOUE3QS1FM0Q0LTVCOEYtQkU4OS1CNkMyM0RBM0UyRjIiLCJsb2dJZCI6IjY2LTEzODE3MTg5NzQ0MzI0NTkiLCJydyI6IlciLCJ0aW1lIjoxNjc4MzI2MTI1MDAwfQ
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries to query. Maximum value: 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The username.
	//
	// example:
	//
	// test
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s QueryCustomAuthIdentityRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCustomAuthIdentityRequest) GoString() string {
	return s.String()
}

func (s *QueryCustomAuthIdentityRequest) GetClientId() *string {
	return s.ClientId
}

func (s *QueryCustomAuthIdentityRequest) GetIdentityType() *string {
	return s.IdentityType
}

func (s *QueryCustomAuthIdentityRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryCustomAuthIdentityRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryCustomAuthIdentityRequest) GetSize() *int32 {
	return s.Size
}

func (s *QueryCustomAuthIdentityRequest) GetUsername() *string {
	return s.Username
}

func (s *QueryCustomAuthIdentityRequest) SetClientId(v string) *QueryCustomAuthIdentityRequest {
	s.ClientId = &v
	return s
}

func (s *QueryCustomAuthIdentityRequest) SetIdentityType(v string) *QueryCustomAuthIdentityRequest {
	s.IdentityType = &v
	return s
}

func (s *QueryCustomAuthIdentityRequest) SetInstanceId(v string) *QueryCustomAuthIdentityRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryCustomAuthIdentityRequest) SetNextToken(v string) *QueryCustomAuthIdentityRequest {
	s.NextToken = &v
	return s
}

func (s *QueryCustomAuthIdentityRequest) SetSize(v int32) *QueryCustomAuthIdentityRequest {
	s.Size = &v
	return s
}

func (s *QueryCustomAuthIdentityRequest) SetUsername(v string) *QueryCustomAuthIdentityRequest {
	s.Username = &v
	return s
}

func (s *QueryCustomAuthIdentityRequest) Validate() error {
	return dara.Validate(s)
}
