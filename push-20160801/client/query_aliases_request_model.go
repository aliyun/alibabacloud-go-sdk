// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAliasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *QueryAliasesRequest
	GetAppKey() *int64
	SetDeviceId(v string) *QueryAliasesRequest
	GetDeviceId() *string
}

type QueryAliasesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 23267207
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e2ba19de97604f55b165576****
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
}

func (s QueryAliasesRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAliasesRequest) GoString() string {
	return s.String()
}

func (s *QueryAliasesRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *QueryAliasesRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *QueryAliasesRequest) SetAppKey(v int64) *QueryAliasesRequest {
	s.AppKey = &v
	return s
}

func (s *QueryAliasesRequest) SetDeviceId(v string) *QueryAliasesRequest {
	s.DeviceId = &v
	return s
}

func (s *QueryAliasesRequest) Validate() error {
	return dara.Validate(s)
}
