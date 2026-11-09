// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCustomAuthPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentity(v string) *QueryCustomAuthPermissionRequest
	GetIdentity() *string
	SetIdentityType(v string) *QueryCustomAuthPermissionRequest
	GetIdentityType() *string
	SetInstanceId(v string) *QueryCustomAuthPermissionRequest
	GetInstanceId() *string
	SetNextToken(v string) *QueryCustomAuthPermissionRequest
	GetNextToken() *string
	SetSize(v int32) *QueryCustomAuthPermissionRequest
	GetSize() *int32
	SetTopic(v string) *QueryCustomAuthPermissionRequest
	GetTopic() *string
}

type QueryCustomAuthPermissionRequest struct {
	// The username or client ID.
	//
	// example:
	//
	// test
	Identity *string `json:"Identity,omitempty" xml:"Identity,omitempty"`
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
	// post-cn-0pp12gl****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The token that marks the end position of the previous returned page. To obtain the next batch of data, call the operation again by using the value of NextToken returned by the previous request. If you call this operation for the first time or want to query all results, set NextToken to an empty string.
	//
	// example:
	//
	// AAAAAThmKW2HkRgzo4G7IRRTK2fC6zZmAk6y0bwoNPFOOcSP
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of queries to be returned. Maximum value: 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The topic whose authorization information you want to query. Multi-level topics and wildcard characters are supported.
	//
	// example:
	//
	// test
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s QueryCustomAuthPermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCustomAuthPermissionRequest) GoString() string {
	return s.String()
}

func (s *QueryCustomAuthPermissionRequest) GetIdentity() *string {
	return s.Identity
}

func (s *QueryCustomAuthPermissionRequest) GetIdentityType() *string {
	return s.IdentityType
}

func (s *QueryCustomAuthPermissionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryCustomAuthPermissionRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryCustomAuthPermissionRequest) GetSize() *int32 {
	return s.Size
}

func (s *QueryCustomAuthPermissionRequest) GetTopic() *string {
	return s.Topic
}

func (s *QueryCustomAuthPermissionRequest) SetIdentity(v string) *QueryCustomAuthPermissionRequest {
	s.Identity = &v
	return s
}

func (s *QueryCustomAuthPermissionRequest) SetIdentityType(v string) *QueryCustomAuthPermissionRequest {
	s.IdentityType = &v
	return s
}

func (s *QueryCustomAuthPermissionRequest) SetInstanceId(v string) *QueryCustomAuthPermissionRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryCustomAuthPermissionRequest) SetNextToken(v string) *QueryCustomAuthPermissionRequest {
	s.NextToken = &v
	return s
}

func (s *QueryCustomAuthPermissionRequest) SetSize(v int32) *QueryCustomAuthPermissionRequest {
	s.Size = &v
	return s
}

func (s *QueryCustomAuthPermissionRequest) SetTopic(v string) *QueryCustomAuthPermissionRequest {
	s.Topic = &v
	return s
}

func (s *QueryCustomAuthPermissionRequest) Validate() error {
	return dara.Validate(s)
}
