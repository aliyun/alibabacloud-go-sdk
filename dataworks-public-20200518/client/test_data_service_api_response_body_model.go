// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestDataServiceApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *TestDataServiceApiResponseBodyData) *TestDataServiceApiResponseBody
	GetData() *TestDataServiceApiResponseBodyData
	SetRequestId(v string) *TestDataServiceApiResponseBody
	GetRequestId() *string
}

type TestDataServiceApiResponseBody struct {
	// The information about the test.
	Data *TestDataServiceApiResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// abdsfewe
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TestDataServiceApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TestDataServiceApiResponseBody) GoString() string {
	return s.String()
}

func (s *TestDataServiceApiResponseBody) GetData() *TestDataServiceApiResponseBodyData {
	return s.Data
}

func (s *TestDataServiceApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TestDataServiceApiResponseBody) SetData(v *TestDataServiceApiResponseBodyData) *TestDataServiceApiResponseBody {
	s.Data = v
	return s
}

func (s *TestDataServiceApiResponseBody) SetRequestId(v string) *TestDataServiceApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *TestDataServiceApiResponseBody) Validate() error {
	return dara.Validate(s)
}

type TestDataServiceApiResponseBodyData struct {
	// The ID of the test.
	//
	// example:
	//
	// 232432
	TestId *string `json:"TestId,omitempty" xml:"TestId,omitempty"`
}

func (s TestDataServiceApiResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TestDataServiceApiResponseBodyData) GoString() string {
	return s.String()
}

func (s *TestDataServiceApiResponseBodyData) GetTestId() *string {
	return s.TestId
}

func (s *TestDataServiceApiResponseBodyData) SetTestId(v string) *TestDataServiceApiResponseBodyData {
	s.TestId = &v
	return s
}

func (s *TestDataServiceApiResponseBodyData) Validate() error {
	return dara.Validate(s)
}
