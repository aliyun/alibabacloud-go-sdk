// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWhiteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegistryId(v int64) *UpdateWhiteListRequest
	GetRegistryId() *int64
	SetWhiteList(v string) *UpdateWhiteListRequest
	GetWhiteList() *string
}

type UpdateWhiteListRequest struct {
	// The ID of the image repository.
	//
	// >  You can call the [PageImageRegistry](~~PageImageRegistry~~) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 19882
	RegistryId *int64 `json:"RegistryId,omitempty" xml:"RegistryId,omitempty"`
	// The IP address whitelist. Separate multiple IP addresses with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.XXX.XXX,192.180.XXX.XXX
	WhiteList *string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty"`
}

func (s UpdateWhiteListRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWhiteListRequest) GoString() string {
	return s.String()
}

func (s *UpdateWhiteListRequest) GetRegistryId() *int64 {
	return s.RegistryId
}

func (s *UpdateWhiteListRequest) GetWhiteList() *string {
	return s.WhiteList
}

func (s *UpdateWhiteListRequest) SetRegistryId(v int64) *UpdateWhiteListRequest {
	s.RegistryId = &v
	return s
}

func (s *UpdateWhiteListRequest) SetWhiteList(v string) *UpdateWhiteListRequest {
	s.WhiteList = &v
	return s
}

func (s *UpdateWhiteListRequest) Validate() error {
	return dara.Validate(s)
}
