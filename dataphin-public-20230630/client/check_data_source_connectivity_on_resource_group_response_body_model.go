// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDataSourceConnectivityOnResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CheckDataSourceConnectivityOnResourceGroupResponseBody
	GetCode() *string
	SetData(v string) *CheckDataSourceConnectivityOnResourceGroupResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CheckDataSourceConnectivityOnResourceGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CheckDataSourceConnectivityOnResourceGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *CheckDataSourceConnectivityOnResourceGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CheckDataSourceConnectivityOnResourceGroupResponseBody
	GetSuccess() *bool
}

type CheckDataSourceConnectivityOnResourceGroupResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 129837xxxx
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckDataSourceConnectivityOnResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckDataSourceConnectivityOnResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDataSourceConnectivityOnResourceGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *CheckDataSourceConnectivityOnResourceGroupResponseBody) GetData() *string {
	return s.Data
}

func (s *CheckDataSourceConnectivityOnResourceGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CheckDataSourceConnectivityOnResourceGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CheckDataSourceConnectivityOnResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckDataSourceConnectivityOnResourceGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CheckDataSourceConnectivityOnResourceGroupResponseBody) SetCode(v string) *CheckDataSourceConnectivityOnResourceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CheckDataSourceConnectivityOnResourceGroupResponseBody) SetData(v string) *CheckDataSourceConnectivityOnResourceGroupResponseBody {
	s.Data = &v
	return s
}

func (s *CheckDataSourceConnectivityOnResourceGroupResponseBody) SetHttpStatusCode(v int32) *CheckDataSourceConnectivityOnResourceGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CheckDataSourceConnectivityOnResourceGroupResponseBody) SetMessage(v string) *CheckDataSourceConnectivityOnResourceGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CheckDataSourceConnectivityOnResourceGroupResponseBody) SetRequestId(v string) *CheckDataSourceConnectivityOnResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckDataSourceConnectivityOnResourceGroupResponseBody) SetSuccess(v bool) *CheckDataSourceConnectivityOnResourceGroupResponseBody {
	s.Success = &v
	return s
}

func (s *CheckDataSourceConnectivityOnResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
