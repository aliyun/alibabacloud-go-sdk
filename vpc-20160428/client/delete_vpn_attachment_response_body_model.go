// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpnAttachmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteVpnAttachmentResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteVpnAttachmentResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteVpnAttachmentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteVpnAttachmentResponseBody
	GetSuccess() *bool
}

type DeleteVpnAttachmentResponseBody struct {
	// The status code returned by the current operation. **200*	- indicates that the operation is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information returned by the current operation.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 29440C49-398F-3D06-BA8B-E3CD13F3246D
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
}

func (s DeleteVpnAttachmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpnAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVpnAttachmentResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteVpnAttachmentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteVpnAttachmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVpnAttachmentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteVpnAttachmentResponseBody) SetCode(v string) *DeleteVpnAttachmentResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteVpnAttachmentResponseBody) SetMessage(v string) *DeleteVpnAttachmentResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteVpnAttachmentResponseBody) SetRequestId(v string) *DeleteVpnAttachmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVpnAttachmentResponseBody) SetSuccess(v bool) *DeleteVpnAttachmentResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteVpnAttachmentResponseBody) Validate() error {
	return dara.Validate(s)
}
