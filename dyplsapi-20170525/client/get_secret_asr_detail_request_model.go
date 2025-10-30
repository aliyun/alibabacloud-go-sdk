// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecretAsrDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallId(v string) *GetSecretAsrDetailRequest
	GetCallId() *string
	SetCallTime(v string) *GetSecretAsrDetailRequest
	GetCallTime() *string
	SetPoolKey(v string) *GetSecretAsrDetailRequest
	GetPoolKey() *string
}

type GetSecretAsrDetailRequest struct {
	// The ID of the call record.
	//
	// You can log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) and view **Call Record ID*	- on the **Call Record Query*	- page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 225625****
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// The call initiation time in the call record.
	//
	// You can log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account). View **Call Initiated At*	- on the **Call Record Query*	- page, or view the call_time field in the Call Detail Record (CDR) receipt.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-03-05 12:00:00
	CallTime *string `json:"CallTime,omitempty" xml:"CallTime,omitempty"`
	// The key of the phone number pool.
	//
	// You can log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) and view the key of the phone number pool on the **Number Pool Management*	- page.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC2267****
	PoolKey *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
}

func (s GetSecretAsrDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSecretAsrDetailRequest) GoString() string {
	return s.String()
}

func (s *GetSecretAsrDetailRequest) GetCallId() *string {
	return s.CallId
}

func (s *GetSecretAsrDetailRequest) GetCallTime() *string {
	return s.CallTime
}

func (s *GetSecretAsrDetailRequest) GetPoolKey() *string {
	return s.PoolKey
}

func (s *GetSecretAsrDetailRequest) SetCallId(v string) *GetSecretAsrDetailRequest {
	s.CallId = &v
	return s
}

func (s *GetSecretAsrDetailRequest) SetCallTime(v string) *GetSecretAsrDetailRequest {
	s.CallTime = &v
	return s
}

func (s *GetSecretAsrDetailRequest) SetPoolKey(v string) *GetSecretAsrDetailRequest {
	s.PoolKey = &v
	return s
}

func (s *GetSecretAsrDetailRequest) Validate() error {
	return dara.Validate(s)
}
