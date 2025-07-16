// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlidingAssistantResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAlidingAssistantResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAlidingAssistantResponseBody
	GetSuccess() *bool
	SetVendorRequestId(v string) *DeleteAlidingAssistantResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *DeleteAlidingAssistantResponseBody
	GetVendorType() *string
}

type DeleteAlidingAssistantResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s DeleteAlidingAssistantResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlidingAssistantResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAlidingAssistantResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAlidingAssistantResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAlidingAssistantResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *DeleteAlidingAssistantResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *DeleteAlidingAssistantResponseBody) SetRequestId(v string) *DeleteAlidingAssistantResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAlidingAssistantResponseBody) SetSuccess(v bool) *DeleteAlidingAssistantResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAlidingAssistantResponseBody) SetVendorRequestId(v string) *DeleteAlidingAssistantResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *DeleteAlidingAssistantResponseBody) SetVendorType(v string) *DeleteAlidingAssistantResponseBody {
	s.VendorType = &v
	return s
}

func (s *DeleteAlidingAssistantResponseBody) Validate() error {
	return dara.Validate(s)
}
