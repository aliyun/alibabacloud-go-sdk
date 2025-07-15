// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpnAttachmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateVpnAttachmentResponseBody
	GetCode() *string
	SetCreateTime(v int64) *CreateVpnAttachmentResponseBody
	GetCreateTime() *int64
	SetMessage(v string) *CreateVpnAttachmentResponseBody
	GetMessage() *string
	SetName(v string) *CreateVpnAttachmentResponseBody
	GetName() *string
	SetRequestId(v string) *CreateVpnAttachmentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateVpnAttachmentResponseBody
	GetSuccess() *bool
	SetVpnConnectionId(v string) *CreateVpnAttachmentResponseBody
	GetVpnConnectionId() *string
}

type CreateVpnAttachmentResponseBody struct {
	// The status code returned by the current operation. **200*	- indicates that the operation is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The timestamp generated when the IPsec-VPN connection was established. Unit: milliseconds.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1658201810000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The information returned by the current operation.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the IPsec-VPN connection.
	//
	// example:
	//
	// nametest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 88187252-0E26-3C4D-9D1D-32A04454EBBA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the current operation is successful.
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the IPsec-VPN connection.
	//
	// example:
	//
	// vco-p0wb09rama8qwwgfn****
	VpnConnectionId *string `json:"VpnConnectionId,omitempty" xml:"VpnConnectionId,omitempty"`
}

func (s CreateVpnAttachmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVpnAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVpnAttachmentResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateVpnAttachmentResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *CreateVpnAttachmentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateVpnAttachmentResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateVpnAttachmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVpnAttachmentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateVpnAttachmentResponseBody) GetVpnConnectionId() *string {
	return s.VpnConnectionId
}

func (s *CreateVpnAttachmentResponseBody) SetCode(v string) *CreateVpnAttachmentResponseBody {
	s.Code = &v
	return s
}

func (s *CreateVpnAttachmentResponseBody) SetCreateTime(v int64) *CreateVpnAttachmentResponseBody {
	s.CreateTime = &v
	return s
}

func (s *CreateVpnAttachmentResponseBody) SetMessage(v string) *CreateVpnAttachmentResponseBody {
	s.Message = &v
	return s
}

func (s *CreateVpnAttachmentResponseBody) SetName(v string) *CreateVpnAttachmentResponseBody {
	s.Name = &v
	return s
}

func (s *CreateVpnAttachmentResponseBody) SetRequestId(v string) *CreateVpnAttachmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVpnAttachmentResponseBody) SetSuccess(v bool) *CreateVpnAttachmentResponseBody {
	s.Success = &v
	return s
}

func (s *CreateVpnAttachmentResponseBody) SetVpnConnectionId(v string) *CreateVpnAttachmentResponseBody {
	s.VpnConnectionId = &v
	return s
}

func (s *CreateVpnAttachmentResponseBody) Validate() error {
	return dara.Validate(s)
}
