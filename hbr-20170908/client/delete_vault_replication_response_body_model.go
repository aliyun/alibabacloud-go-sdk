// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVaultReplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteVaultReplicationResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteVaultReplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteVaultReplicationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteVaultReplicationResponseBody
	GetSuccess() *bool
}

type DeleteVaultReplicationResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// C438054F-9088-5D1B-AED0-0EA86D9C65F4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteVaultReplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVaultReplicationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVaultReplicationResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteVaultReplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteVaultReplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVaultReplicationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteVaultReplicationResponseBody) SetCode(v string) *DeleteVaultReplicationResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteVaultReplicationResponseBody) SetMessage(v string) *DeleteVaultReplicationResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteVaultReplicationResponseBody) SetRequestId(v string) *DeleteVaultReplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVaultReplicationResponseBody) SetSuccess(v bool) *DeleteVaultReplicationResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteVaultReplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
