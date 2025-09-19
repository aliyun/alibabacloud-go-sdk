// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOpaStrategyNewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateOpaStrategyNewResponseBody
	GetCode() *string
	SetData(v []*string) *CreateOpaStrategyNewResponseBody
	GetData() []*string
	SetHttpStatusCode(v int32) *CreateOpaStrategyNewResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateOpaStrategyNewResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateOpaStrategyNewResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateOpaStrategyNewResponseBody
	GetSuccess() *bool
}

type CreateOpaStrategyNewResponseBody struct {
	// The status code returned. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The IDs of the clusters that failed to be saved.
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CD380235-A0B8-540D-A0D5-D62884469E3C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
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

func (s CreateOpaStrategyNewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOpaStrategyNewResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOpaStrategyNewResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateOpaStrategyNewResponseBody) GetData() []*string {
	return s.Data
}

func (s *CreateOpaStrategyNewResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateOpaStrategyNewResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateOpaStrategyNewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOpaStrategyNewResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateOpaStrategyNewResponseBody) SetCode(v string) *CreateOpaStrategyNewResponseBody {
	s.Code = &v
	return s
}

func (s *CreateOpaStrategyNewResponseBody) SetData(v []*string) *CreateOpaStrategyNewResponseBody {
	s.Data = v
	return s
}

func (s *CreateOpaStrategyNewResponseBody) SetHttpStatusCode(v int32) *CreateOpaStrategyNewResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateOpaStrategyNewResponseBody) SetMessage(v string) *CreateOpaStrategyNewResponseBody {
	s.Message = &v
	return s
}

func (s *CreateOpaStrategyNewResponseBody) SetRequestId(v string) *CreateOpaStrategyNewResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOpaStrategyNewResponseBody) SetSuccess(v bool) *CreateOpaStrategyNewResponseBody {
	s.Success = &v
	return s
}

func (s *CreateOpaStrategyNewResponseBody) Validate() error {
	return dara.Validate(s)
}
