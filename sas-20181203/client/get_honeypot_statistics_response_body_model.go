// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHoneypotStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetHoneypotStatisticsResponseBody
	GetCode() *string
	SetData(v *GetHoneypotStatisticsResponseBodyData) *GetHoneypotStatisticsResponseBody
	GetData() *GetHoneypotStatisticsResponseBodyData
	SetHttpStatusCode(v int32) *GetHoneypotStatisticsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetHoneypotStatisticsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHoneypotStatisticsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetHoneypotStatisticsResponseBody
	GetSuccess() *bool
}

type GetHoneypotStatisticsResponseBody struct {
	// The status code that is returned. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The honeypot usage statistics.
	Data *GetHoneypotStatisticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EB7181CB-32F3-5189-A935-4E24DD1A****
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

func (s GetHoneypotStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHoneypotStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetHoneypotStatisticsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetHoneypotStatisticsResponseBody) GetData() *GetHoneypotStatisticsResponseBodyData {
	return s.Data
}

func (s *GetHoneypotStatisticsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetHoneypotStatisticsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHoneypotStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHoneypotStatisticsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetHoneypotStatisticsResponseBody) SetCode(v string) *GetHoneypotStatisticsResponseBody {
	s.Code = &v
	return s
}

func (s *GetHoneypotStatisticsResponseBody) SetData(v *GetHoneypotStatisticsResponseBodyData) *GetHoneypotStatisticsResponseBody {
	s.Data = v
	return s
}

func (s *GetHoneypotStatisticsResponseBody) SetHttpStatusCode(v int32) *GetHoneypotStatisticsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetHoneypotStatisticsResponseBody) SetMessage(v string) *GetHoneypotStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *GetHoneypotStatisticsResponseBody) SetRequestId(v string) *GetHoneypotStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHoneypotStatisticsResponseBody) SetSuccess(v bool) *GetHoneypotStatisticsResponseBody {
	s.Success = &v
	return s
}

func (s *GetHoneypotStatisticsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHoneypotStatisticsResponseBodyData struct {
	// The total number of honeypots.
	//
	// example:
	//
	// 20
	TotalHoneypotCount *int32 `json:"TotalHoneypotCount,omitempty" xml:"TotalHoneypotCount,omitempty"`
	// The health status of the management node. Valid values:
	//
	// 	- 1: normal
	//
	// 	- 2: abnormal
	//
	// example:
	//
	// 1
	TotalNodeStatus *int32 `json:"TotalNodeStatus,omitempty" xml:"TotalNodeStatus,omitempty"`
	// The total number of authorized probes.
	//
	// example:
	//
	// 40
	TotalProbeCount *int32 `json:"TotalProbeCount,omitempty" xml:"TotalProbeCount,omitempty"`
	// The number of deployed honeypots.
	//
	// example:
	//
	// 7
	UsedHoneypotCount *int32 `json:"UsedHoneypotCount,omitempty" xml:"UsedHoneypotCount,omitempty"`
	// The number of deployed host probes.
	//
	// example:
	//
	// 9
	UsedHostProbeCount *int32 `json:"UsedHostProbeCount,omitempty" xml:"UsedHostProbeCount,omitempty"`
	// The number of deployed probes.
	//
	// example:
	//
	// 15
	UsedProbeCount *int32 `json:"UsedProbeCount,omitempty" xml:"UsedProbeCount,omitempty"`
	// The number of deployed VPC probes.
	//
	// example:
	//
	// 6
	UsedVpcProbeCount *int32 `json:"UsedVpcProbeCount,omitempty" xml:"UsedVpcProbeCount,omitempty"`
}

func (s GetHoneypotStatisticsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetHoneypotStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHoneypotStatisticsResponseBodyData) GetTotalHoneypotCount() *int32 {
	return s.TotalHoneypotCount
}

func (s *GetHoneypotStatisticsResponseBodyData) GetTotalNodeStatus() *int32 {
	return s.TotalNodeStatus
}

func (s *GetHoneypotStatisticsResponseBodyData) GetTotalProbeCount() *int32 {
	return s.TotalProbeCount
}

func (s *GetHoneypotStatisticsResponseBodyData) GetUsedHoneypotCount() *int32 {
	return s.UsedHoneypotCount
}

func (s *GetHoneypotStatisticsResponseBodyData) GetUsedHostProbeCount() *int32 {
	return s.UsedHostProbeCount
}

func (s *GetHoneypotStatisticsResponseBodyData) GetUsedProbeCount() *int32 {
	return s.UsedProbeCount
}

func (s *GetHoneypotStatisticsResponseBodyData) GetUsedVpcProbeCount() *int32 {
	return s.UsedVpcProbeCount
}

func (s *GetHoneypotStatisticsResponseBodyData) SetTotalHoneypotCount(v int32) *GetHoneypotStatisticsResponseBodyData {
	s.TotalHoneypotCount = &v
	return s
}

func (s *GetHoneypotStatisticsResponseBodyData) SetTotalNodeStatus(v int32) *GetHoneypotStatisticsResponseBodyData {
	s.TotalNodeStatus = &v
	return s
}

func (s *GetHoneypotStatisticsResponseBodyData) SetTotalProbeCount(v int32) *GetHoneypotStatisticsResponseBodyData {
	s.TotalProbeCount = &v
	return s
}

func (s *GetHoneypotStatisticsResponseBodyData) SetUsedHoneypotCount(v int32) *GetHoneypotStatisticsResponseBodyData {
	s.UsedHoneypotCount = &v
	return s
}

func (s *GetHoneypotStatisticsResponseBodyData) SetUsedHostProbeCount(v int32) *GetHoneypotStatisticsResponseBodyData {
	s.UsedHostProbeCount = &v
	return s
}

func (s *GetHoneypotStatisticsResponseBodyData) SetUsedProbeCount(v int32) *GetHoneypotStatisticsResponseBodyData {
	s.UsedProbeCount = &v
	return s
}

func (s *GetHoneypotStatisticsResponseBodyData) SetUsedVpcProbeCount(v int32) *GetHoneypotStatisticsResponseBodyData {
	s.UsedVpcProbeCount = &v
	return s
}

func (s *GetHoneypotStatisticsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
