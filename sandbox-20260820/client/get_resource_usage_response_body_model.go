// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetResourceUsageResponseBody
	GetCode() *string
	SetMessage(v string) *GetResourceUsageResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetResourceUsageResponseBody
	GetRequestId() *string
	SetResourceUsage(v *ResourceUsage) *GetResourceUsageResponseBody
	GetResourceUsage() *ResourceUsage
}

type GetResourceUsageResponseBody struct {
	// The response code. A value of 200 indicates success. Other values indicate specific error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The message that describes the result of the request.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The unique identifier of the request.
	//
	// example:
	//
	// 12345678-1234-5678-9abc-123456789abc
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The usage results for the query interval and each group.
	//
	// example:
	//
	// {"startTime":"2026-11-05T00:00:00Z","endTime":"2026-11-05T02:00:00Z","granularity":"Hour","groupBy":["teamID"],"groups":[{"dimensions":{"teamID":"70d1c834-0333-58d8-97ac-5336eb9179eb"},"items":[{"usageName":"cpuActive","unit":"CoreSecond","total":5400,"buckets":[{"startTime":"2026-11-05T00:00:00Z","endTime":"2026-11-05T01:00:00Z","quantity":3600},{"startTime":"2026-11-05T01:00:00Z","endTime":"2026-11-05T02:00:00Z","quantity":1800}]}]}]}
	ResourceUsage *ResourceUsage `json:"resourceUsage,omitempty" xml:"resourceUsage,omitempty"`
}

func (s GetResourceUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceUsageResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceUsageResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetResourceUsageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetResourceUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceUsageResponseBody) GetResourceUsage() *ResourceUsage {
	return s.ResourceUsage
}

func (s *GetResourceUsageResponseBody) SetCode(v string) *GetResourceUsageResponseBody {
	s.Code = &v
	return s
}

func (s *GetResourceUsageResponseBody) SetMessage(v string) *GetResourceUsageResponseBody {
	s.Message = &v
	return s
}

func (s *GetResourceUsageResponseBody) SetRequestId(v string) *GetResourceUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceUsageResponseBody) SetResourceUsage(v *ResourceUsage) *GetResourceUsageResponseBody {
	s.ResourceUsage = v
	return s
}

func (s *GetResourceUsageResponseBody) Validate() error {
	if s.ResourceUsage != nil {
		if err := s.ResourceUsage.Validate(); err != nil {
			return err
		}
	}
	return nil
}
