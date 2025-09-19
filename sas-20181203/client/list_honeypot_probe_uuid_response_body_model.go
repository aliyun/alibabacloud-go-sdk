// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHoneypotProbeUuidResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListHoneypotProbeUuidResponseBody
	GetCode() *string
	SetCount(v int32) *ListHoneypotProbeUuidResponseBody
	GetCount() *int32
	SetData(v []*string) *ListHoneypotProbeUuidResponseBody
	GetData() []*string
	SetHttpStatusCode(v int32) *ListHoneypotProbeUuidResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListHoneypotProbeUuidResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListHoneypotProbeUuidResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListHoneypotProbeUuidResponseBody
	GetSuccess() *bool
}

type ListHoneypotProbeUuidResponseBody struct {
	// The response code. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The resources of the probe.
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 32C9C2A6-B837-538E-921B-90746CB*****
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

func (s ListHoneypotProbeUuidResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotProbeUuidResponseBody) GoString() string {
	return s.String()
}

func (s *ListHoneypotProbeUuidResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListHoneypotProbeUuidResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ListHoneypotProbeUuidResponseBody) GetData() []*string {
	return s.Data
}

func (s *ListHoneypotProbeUuidResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListHoneypotProbeUuidResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListHoneypotProbeUuidResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHoneypotProbeUuidResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListHoneypotProbeUuidResponseBody) SetCode(v string) *ListHoneypotProbeUuidResponseBody {
	s.Code = &v
	return s
}

func (s *ListHoneypotProbeUuidResponseBody) SetCount(v int32) *ListHoneypotProbeUuidResponseBody {
	s.Count = &v
	return s
}

func (s *ListHoneypotProbeUuidResponseBody) SetData(v []*string) *ListHoneypotProbeUuidResponseBody {
	s.Data = v
	return s
}

func (s *ListHoneypotProbeUuidResponseBody) SetHttpStatusCode(v int32) *ListHoneypotProbeUuidResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListHoneypotProbeUuidResponseBody) SetMessage(v string) *ListHoneypotProbeUuidResponseBody {
	s.Message = &v
	return s
}

func (s *ListHoneypotProbeUuidResponseBody) SetRequestId(v string) *ListHoneypotProbeUuidResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHoneypotProbeUuidResponseBody) SetSuccess(v bool) *ListHoneypotProbeUuidResponseBody {
	s.Success = &v
	return s
}

func (s *ListHoneypotProbeUuidResponseBody) Validate() error {
	return dara.Validate(s)
}
