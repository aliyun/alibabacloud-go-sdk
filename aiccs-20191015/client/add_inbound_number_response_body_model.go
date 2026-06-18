// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddInboundNumberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *AddInboundNumberResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *AddInboundNumberResponseBody
	GetCode() *string
	SetData(v []*AddInboundNumberResponseBodyData) *AddInboundNumberResponseBody
	GetData() []*AddInboundNumberResponseBodyData
	SetMessage(v string) *AddInboundNumberResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddInboundNumberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddInboundNumberResponseBody
	GetSuccess() *bool
}

type AddInboundNumberResponseBody struct {
	// Details about why access was denied.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data []*AddInboundNumberResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The status code description.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D6A51251-F7C4-596A-9F45-3C3219A5450D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded.
	//
	// - **true**: The call succeeded.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddInboundNumberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddInboundNumberResponseBody) GoString() string {
	return s.String()
}

func (s *AddInboundNumberResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *AddInboundNumberResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddInboundNumberResponseBody) GetData() []*AddInboundNumberResponseBodyData {
	return s.Data
}

func (s *AddInboundNumberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddInboundNumberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddInboundNumberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddInboundNumberResponseBody) SetAccessDeniedDetail(v string) *AddInboundNumberResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AddInboundNumberResponseBody) SetCode(v string) *AddInboundNumberResponseBody {
	s.Code = &v
	return s
}

func (s *AddInboundNumberResponseBody) SetData(v []*AddInboundNumberResponseBodyData) *AddInboundNumberResponseBody {
	s.Data = v
	return s
}

func (s *AddInboundNumberResponseBody) SetMessage(v string) *AddInboundNumberResponseBody {
	s.Message = &v
	return s
}

func (s *AddInboundNumberResponseBody) SetRequestId(v string) *AddInboundNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddInboundNumberResponseBody) SetSuccess(v bool) *AddInboundNumberResponseBody {
	s.Success = &v
	return s
}

func (s *AddInboundNumberResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddInboundNumberResponseBodyData struct {
	// The inbound number.
	//
	// example:
	//
	// 234234238**33
	InboundNumber *string `json:"InboundNumber,omitempty" xml:"InboundNumber,omitempty"`
	// The error description.
	//
	// example:
	//
	// 号码不合法
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the number was added successfully. `true` indicates success, and `false` indicates failure.
	//
	// example:
	//
	// false
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s AddInboundNumberResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddInboundNumberResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddInboundNumberResponseBodyData) GetInboundNumber() *string {
	return s.InboundNumber
}

func (s *AddInboundNumberResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *AddInboundNumberResponseBodyData) GetResult() *bool {
	return s.Result
}

func (s *AddInboundNumberResponseBodyData) SetInboundNumber(v string) *AddInboundNumberResponseBodyData {
	s.InboundNumber = &v
	return s
}

func (s *AddInboundNumberResponseBodyData) SetMessage(v string) *AddInboundNumberResponseBodyData {
	s.Message = &v
	return s
}

func (s *AddInboundNumberResponseBodyData) SetResult(v bool) *AddInboundNumberResponseBodyData {
	s.Result = &v
	return s
}

func (s *AddInboundNumberResponseBodyData) Validate() error {
	return dara.Validate(s)
}
