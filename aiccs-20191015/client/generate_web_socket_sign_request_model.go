// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateWebSocketSignRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *GenerateWebSocketSignRequest
	GetAccountName() *string
	SetClientToken(v string) *GenerateWebSocketSignRequest
	GetClientToken() *string
	SetInstanceId(v string) *GenerateWebSocketSignRequest
	GetInstanceId() *string
}

type GenerateWebSocketSignRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123@123.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// 46c1341e-2648-447a-9b11-70b6a298d94d
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GenerateWebSocketSignRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateWebSocketSignRequest) GoString() string {
	return s.String()
}

func (s *GenerateWebSocketSignRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *GenerateWebSocketSignRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GenerateWebSocketSignRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GenerateWebSocketSignRequest) SetAccountName(v string) *GenerateWebSocketSignRequest {
	s.AccountName = &v
	return s
}

func (s *GenerateWebSocketSignRequest) SetClientToken(v string) *GenerateWebSocketSignRequest {
	s.ClientToken = &v
	return s
}

func (s *GenerateWebSocketSignRequest) SetInstanceId(v string) *GenerateWebSocketSignRequest {
	s.InstanceId = &v
	return s
}

func (s *GenerateWebSocketSignRequest) Validate() error {
	return dara.Validate(s)
}
