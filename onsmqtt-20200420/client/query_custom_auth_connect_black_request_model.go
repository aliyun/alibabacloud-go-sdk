// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCustomAuthConnectBlackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *QueryCustomAuthConnectBlackRequest
	GetClientId() *string
	SetInstanceId(v string) *QueryCustomAuthConnectBlackRequest
	GetInstanceId() *string
	SetNextToken(v string) *QueryCustomAuthConnectBlackRequest
	GetNextToken() *string
	SetSize(v int32) *QueryCustomAuthConnectBlackRequest
	GetSize() *int32
}

type QueryCustomAuthConnectBlackRequest struct {
	// The ID of the client to be queried.
	//
	// example:
	//
	// GID_test@@@test
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The ID of the ApsaraMQ for MQTT instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// post-111****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The token that marks the end position of the previous returned page. To obtain the next batch of data, call the operation again by using the value of NextToken returned by the previous request. If you call this operation for the first time or want to query all results, set NextToken to an empty string.
	//
	// example:
	//
	// xOfRU60sGEwN1OlFBIL8Ew==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of clients to be queried. Maximum value: 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s QueryCustomAuthConnectBlackRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCustomAuthConnectBlackRequest) GoString() string {
	return s.String()
}

func (s *QueryCustomAuthConnectBlackRequest) GetClientId() *string {
	return s.ClientId
}

func (s *QueryCustomAuthConnectBlackRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryCustomAuthConnectBlackRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryCustomAuthConnectBlackRequest) GetSize() *int32 {
	return s.Size
}

func (s *QueryCustomAuthConnectBlackRequest) SetClientId(v string) *QueryCustomAuthConnectBlackRequest {
	s.ClientId = &v
	return s
}

func (s *QueryCustomAuthConnectBlackRequest) SetInstanceId(v string) *QueryCustomAuthConnectBlackRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryCustomAuthConnectBlackRequest) SetNextToken(v string) *QueryCustomAuthConnectBlackRequest {
	s.NextToken = &v
	return s
}

func (s *QueryCustomAuthConnectBlackRequest) SetSize(v int32) *QueryCustomAuthConnectBlackRequest {
	s.Size = &v
	return s
}

func (s *QueryCustomAuthConnectBlackRequest) Validate() error {
	return dara.Validate(s)
}
