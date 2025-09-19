// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebLockCreateConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v string) *ModifyWebLockCreateConfigResponseBody
	GetConfigId() *string
	SetRequestId(v string) *ModifyWebLockCreateConfigResponseBody
	GetRequestId() *string
}

type ModifyWebLockCreateConfigResponseBody struct {
	// The configuration ID of the protected directory.
	//
	// example:
	//
	// 1404656
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// D9354C1A-D709-4873-9AAE-41513327B247
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebLockCreateConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebLockCreateConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebLockCreateConfigResponseBody) GetConfigId() *string {
	return s.ConfigId
}

func (s *ModifyWebLockCreateConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyWebLockCreateConfigResponseBody) SetConfigId(v string) *ModifyWebLockCreateConfigResponseBody {
	s.ConfigId = &v
	return s
}

func (s *ModifyWebLockCreateConfigResponseBody) SetRequestId(v string) *ModifyWebLockCreateConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyWebLockCreateConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
