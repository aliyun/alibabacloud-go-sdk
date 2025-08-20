// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCompactionServiceSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeCompactionServiceSwitchResponseBodyData) *DescribeCompactionServiceSwitchResponseBody
	GetData() *DescribeCompactionServiceSwitchResponseBodyData
	SetRequestId(v string) *DescribeCompactionServiceSwitchResponseBody
	GetRequestId() *string
}

type DescribeCompactionServiceSwitchResponseBody struct {
	// The returned data.
	Data *DescribeCompactionServiceSwitchResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D761DA51-12F8-5457-AAA9-F52B9F436D2D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCompactionServiceSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCompactionServiceSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCompactionServiceSwitchResponseBody) GetData() *DescribeCompactionServiceSwitchResponseBodyData {
	return s.Data
}

func (s *DescribeCompactionServiceSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCompactionServiceSwitchResponseBody) SetData(v *DescribeCompactionServiceSwitchResponseBodyData) *DescribeCompactionServiceSwitchResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCompactionServiceSwitchResponseBody) SetRequestId(v string) *DescribeCompactionServiceSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCompactionServiceSwitchResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCompactionServiceSwitchResponseBodyData struct {
	// Indicates whether the remote build feature is enabled.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableCompactionService *bool `json:"EnableCompactionService,omitempty" xml:"EnableCompactionService,omitempty"`
}

func (s DescribeCompactionServiceSwitchResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCompactionServiceSwitchResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCompactionServiceSwitchResponseBodyData) GetEnableCompactionService() *bool {
	return s.EnableCompactionService
}

func (s *DescribeCompactionServiceSwitchResponseBodyData) SetEnableCompactionService(v bool) *DescribeCompactionServiceSwitchResponseBodyData {
	s.EnableCompactionService = &v
	return s
}

func (s *DescribeCompactionServiceSwitchResponseBodyData) Validate() error {
	return dara.Validate(s)
}
