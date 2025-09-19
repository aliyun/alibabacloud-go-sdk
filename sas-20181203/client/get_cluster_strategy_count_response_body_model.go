// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterStrategyCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetClusterStrategyCountResponseBody
	GetCode() *string
	SetCount(v int32) *GetClusterStrategyCountResponseBody
	GetCount() *int32
	SetData(v []*GetClusterStrategyCountResponseBodyData) *GetClusterStrategyCountResponseBody
	GetData() []*GetClusterStrategyCountResponseBodyData
	SetHttpStatusCode(v int32) *GetClusterStrategyCountResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetClusterStrategyCountResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetClusterStrategyCountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetClusterStrategyCountResponseBody
	GetSuccess() *bool
}

type GetClusterStrategyCountResponseBody struct {
	// The status code returned. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 3
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The returned data.
	Data []*GetClusterStrategyCountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code that is returned.
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
	// 796348B5-115C-5BFB-83EA-B5C3C12F822F
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

func (s GetClusterStrategyCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetClusterStrategyCountResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterStrategyCountResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetClusterStrategyCountResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *GetClusterStrategyCountResponseBody) GetData() []*GetClusterStrategyCountResponseBodyData {
	return s.Data
}

func (s *GetClusterStrategyCountResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetClusterStrategyCountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetClusterStrategyCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetClusterStrategyCountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetClusterStrategyCountResponseBody) SetCode(v string) *GetClusterStrategyCountResponseBody {
	s.Code = &v
	return s
}

func (s *GetClusterStrategyCountResponseBody) SetCount(v int32) *GetClusterStrategyCountResponseBody {
	s.Count = &v
	return s
}

func (s *GetClusterStrategyCountResponseBody) SetData(v []*GetClusterStrategyCountResponseBodyData) *GetClusterStrategyCountResponseBody {
	s.Data = v
	return s
}

func (s *GetClusterStrategyCountResponseBody) SetHttpStatusCode(v int32) *GetClusterStrategyCountResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetClusterStrategyCountResponseBody) SetMessage(v string) *GetClusterStrategyCountResponseBody {
	s.Message = &v
	return s
}

func (s *GetClusterStrategyCountResponseBody) SetRequestId(v string) *GetClusterStrategyCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClusterStrategyCountResponseBody) SetSuccess(v bool) *GetClusterStrategyCountResponseBody {
	s.Success = &v
	return s
}

func (s *GetClusterStrategyCountResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetClusterStrategyCountResponseBodyData struct {
	// The cluster ID.
	//
	// example:
	//
	// c8ca91e0907d94efaba7fb0827eb9****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The number of policies.
	//
	// example:
	//
	// 10
	StrategyCount *int32 `json:"StrategyCount,omitempty" xml:"StrategyCount,omitempty"`
}

func (s GetClusterStrategyCountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetClusterStrategyCountResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetClusterStrategyCountResponseBodyData) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetClusterStrategyCountResponseBodyData) GetStrategyCount() *int32 {
	return s.StrategyCount
}

func (s *GetClusterStrategyCountResponseBodyData) SetClusterId(v string) *GetClusterStrategyCountResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *GetClusterStrategyCountResponseBodyData) SetStrategyCount(v int32) *GetClusterStrategyCountResponseBodyData {
	s.StrategyCount = &v
	return s
}

func (s *GetClusterStrategyCountResponseBodyData) Validate() error {
	return dara.Validate(s)
}
