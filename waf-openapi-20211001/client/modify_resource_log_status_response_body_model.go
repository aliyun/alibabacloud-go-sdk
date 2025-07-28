// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyResourceLogStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyResourceLogStatusResponseBody
	GetRequestId() *string
	SetStatus(v bool) *ModifyResourceLogStatusResponseBody
	GetStatus() *bool
}

type ModifyResourceLogStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 7C55A3E5-638A-5D6E-9A2F-C3CE5A677EC5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the log collection feature is enabled for the protected object. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyResourceLogStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyResourceLogStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyResourceLogStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyResourceLogStatusResponseBody) GetStatus() *bool {
	return s.Status
}

func (s *ModifyResourceLogStatusResponseBody) SetRequestId(v string) *ModifyResourceLogStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyResourceLogStatusResponseBody) SetStatus(v bool) *ModifyResourceLogStatusResponseBody {
	s.Status = &v
	return s
}

func (s *ModifyResourceLogStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
