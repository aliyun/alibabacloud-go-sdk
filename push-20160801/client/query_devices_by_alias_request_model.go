// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDevicesByAliasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v string) *QueryDevicesByAliasRequest
	GetAlias() *string
	SetAppKey(v int64) *QueryDevicesByAliasRequest
	GetAppKey() *int64
}

type QueryDevicesByAliasRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// aliasName
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 23267207
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s QueryDevicesByAliasRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDevicesByAliasRequest) GoString() string {
	return s.String()
}

func (s *QueryDevicesByAliasRequest) GetAlias() *string {
	return s.Alias
}

func (s *QueryDevicesByAliasRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *QueryDevicesByAliasRequest) SetAlias(v string) *QueryDevicesByAliasRequest {
	s.Alias = &v
	return s
}

func (s *QueryDevicesByAliasRequest) SetAppKey(v int64) *QueryDevicesByAliasRequest {
	s.AppKey = &v
	return s
}

func (s *QueryDevicesByAliasRequest) Validate() error {
	return dara.Validate(s)
}
