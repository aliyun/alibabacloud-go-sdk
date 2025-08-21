// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDevicesByAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccount(v string) *QueryDevicesByAccountRequest
	GetAccount() *string
	SetAppKey(v int64) *QueryDevicesByAccountRequest
	GetAppKey() *int64
}

type QueryDevicesByAccountRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// accountName
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 23267207
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s QueryDevicesByAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDevicesByAccountRequest) GoString() string {
	return s.String()
}

func (s *QueryDevicesByAccountRequest) GetAccount() *string {
	return s.Account
}

func (s *QueryDevicesByAccountRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *QueryDevicesByAccountRequest) SetAccount(v string) *QueryDevicesByAccountRequest {
	s.Account = &v
	return s
}

func (s *QueryDevicesByAccountRequest) SetAppKey(v int64) *QueryDevicesByAccountRequest {
	s.AppKey = &v
	return s
}

func (s *QueryDevicesByAccountRequest) Validate() error {
	return dara.Validate(s)
}
