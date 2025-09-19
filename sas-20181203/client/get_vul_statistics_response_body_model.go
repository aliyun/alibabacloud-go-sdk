// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVulStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetVulStatisticsResponseBody
	GetRequestId() *string
	SetVulAsapSum(v int32) *GetVulStatisticsResponseBody
	GetVulAsapSum() *int32
	SetVulLaterSum(v int32) *GetVulStatisticsResponseBody
	GetVulLaterSum() *int32
	SetVulNntfSum(v int32) *GetVulStatisticsResponseBody
	GetVulNntfSum() *int32
}

type GetVulStatisticsResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 3FE272FA-7263-4554-A90F-A7857945A6D5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of vulnerabilities that have the high priority.
	//
	// example:
	//
	// 16
	VulAsapSum *int32 `json:"VulAsapSum,omitempty" xml:"VulAsapSum,omitempty"`
	// The number of vulnerabilities that have the medium priority.
	//
	// example:
	//
	// 0
	VulLaterSum *int32 `json:"VulLaterSum,omitempty" xml:"VulLaterSum,omitempty"`
	// The number of vulnerabilities that have the low priority.
	//
	// example:
	//
	// 0
	VulNntfSum *int32 `json:"VulNntfSum,omitempty" xml:"VulNntfSum,omitempty"`
}

func (s GetVulStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVulStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetVulStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVulStatisticsResponseBody) GetVulAsapSum() *int32 {
	return s.VulAsapSum
}

func (s *GetVulStatisticsResponseBody) GetVulLaterSum() *int32 {
	return s.VulLaterSum
}

func (s *GetVulStatisticsResponseBody) GetVulNntfSum() *int32 {
	return s.VulNntfSum
}

func (s *GetVulStatisticsResponseBody) SetRequestId(v string) *GetVulStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVulStatisticsResponseBody) SetVulAsapSum(v int32) *GetVulStatisticsResponseBody {
	s.VulAsapSum = &v
	return s
}

func (s *GetVulStatisticsResponseBody) SetVulLaterSum(v int32) *GetVulStatisticsResponseBody {
	s.VulLaterSum = &v
	return s
}

func (s *GetVulStatisticsResponseBody) SetVulNntfSum(v int32) *GetVulStatisticsResponseBody {
	s.VulNntfSum = &v
	return s
}

func (s *GetVulStatisticsResponseBody) Validate() error {
	return dara.Validate(s)
}
