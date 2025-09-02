// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDIProjectConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpdateDIProjectConfigResponseBodyData) *UpdateDIProjectConfigResponseBody
	GetData() *UpdateDIProjectConfigResponseBodyData
	SetRequestId(v string) *UpdateDIProjectConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDIProjectConfigResponseBody
	GetSuccess() *bool
}

type UpdateDIProjectConfigResponseBody struct {
	// The information about the modification.
	Data *UpdateDIProjectConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID. You can locate logs and troubleshoot issues based on the ID.
	//
	// example:
	//
	// 0000-ABCD-EFG
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDIProjectConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDIProjectConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDIProjectConfigResponseBody) GetData() *UpdateDIProjectConfigResponseBodyData {
	return s.Data
}

func (s *UpdateDIProjectConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDIProjectConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDIProjectConfigResponseBody) SetData(v *UpdateDIProjectConfigResponseBodyData) *UpdateDIProjectConfigResponseBody {
	s.Data = v
	return s
}

func (s *UpdateDIProjectConfigResponseBody) SetRequestId(v string) *UpdateDIProjectConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDIProjectConfigResponseBody) SetSuccess(v bool) *UpdateDIProjectConfigResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDIProjectConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateDIProjectConfigResponseBodyData struct {
	// Indicates whether the default global configuration of synchronization solutions is modified. Valid values:
	//
	// 	- success
	//
	// 	- fail
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateDIProjectConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateDIProjectConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateDIProjectConfigResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *UpdateDIProjectConfigResponseBodyData) SetStatus(v string) *UpdateDIProjectConfigResponseBodyData {
	s.Status = &v
	return s
}

func (s *UpdateDIProjectConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
