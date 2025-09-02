// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDataSourceShareResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SetDataSourceShareResponseBodyData) *SetDataSourceShareResponseBody
	GetData() *SetDataSourceShareResponseBodyData
	SetRequestId(v string) *SetDataSourceShareResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SetDataSourceShareResponseBody
	GetSuccess() *bool
}

type SetDataSourceShareResponseBody struct {
	// The information about the sharing operation.
	Data *SetDataSourceShareResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
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

func (s SetDataSourceShareResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDataSourceShareResponseBody) GoString() string {
	return s.String()
}

func (s *SetDataSourceShareResponseBody) GetData() *SetDataSourceShareResponseBodyData {
	return s.Data
}

func (s *SetDataSourceShareResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDataSourceShareResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetDataSourceShareResponseBody) SetData(v *SetDataSourceShareResponseBodyData) *SetDataSourceShareResponseBody {
	s.Data = v
	return s
}

func (s *SetDataSourceShareResponseBody) SetRequestId(v string) *SetDataSourceShareResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDataSourceShareResponseBody) SetSuccess(v bool) *SetDataSourceShareResponseBody {
	s.Success = &v
	return s
}

func (s *SetDataSourceShareResponseBody) Validate() error {
	return dara.Validate(s)
}

type SetDataSourceShareResponseBodyData struct {
	// The reason why the data source failed to be shared. If the data source is successfully shared, the value of this parameter is an empty string.
	//
	// example:
	//
	// datasource is wrong
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the data source was shared. Valid values:
	//
	// 	- success.
	//
	// 	- fail. You can view the value of the Message parameter to identify the cause of the failure.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SetDataSourceShareResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SetDataSourceShareResponseBodyData) GoString() string {
	return s.String()
}

func (s *SetDataSourceShareResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *SetDataSourceShareResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *SetDataSourceShareResponseBodyData) SetMessage(v string) *SetDataSourceShareResponseBodyData {
	s.Message = &v
	return s
}

func (s *SetDataSourceShareResponseBodyData) SetStatus(v string) *SetDataSourceShareResponseBodyData {
	s.Status = &v
	return s
}

func (s *SetDataSourceShareResponseBodyData) Validate() error {
	return dara.Validate(s)
}
