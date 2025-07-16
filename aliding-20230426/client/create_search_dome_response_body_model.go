// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSearchDomeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArguments(v []interface{}) *CreateSearchDomeResponseBody
	GetArguments() []interface{}
	SetRequestId(v string) *CreateSearchDomeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateSearchDomeResponseBody
	GetSuccess() *bool
	SetVendorRequestId(v string) *CreateSearchDomeResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *CreateSearchDomeResponseBody
	GetVendorType() *string
}

type CreateSearchDomeResponseBody struct {
	// example:
	//
	// []
	Arguments []interface{} `json:"arguments,omitempty" xml:"arguments,omitempty" type:"Repeated"`
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

func (s CreateSearchDomeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSearchDomeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSearchDomeResponseBody) GetArguments() []interface{} {
	return s.Arguments
}

func (s *CreateSearchDomeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSearchDomeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSearchDomeResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *CreateSearchDomeResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *CreateSearchDomeResponseBody) SetArguments(v []interface{}) *CreateSearchDomeResponseBody {
	s.Arguments = v
	return s
}

func (s *CreateSearchDomeResponseBody) SetRequestId(v string) *CreateSearchDomeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSearchDomeResponseBody) SetSuccess(v bool) *CreateSearchDomeResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSearchDomeResponseBody) SetVendorRequestId(v string) *CreateSearchDomeResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *CreateSearchDomeResponseBody) SetVendorType(v string) *CreateSearchDomeResponseBody {
	s.VendorType = &v
	return s
}

func (s *CreateSearchDomeResponseBody) Validate() error {
	return dara.Validate(s)
}
