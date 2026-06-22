// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTenantCheckAvailableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetTenantCheckAvailableResponseBodyData) *GetTenantCheckAvailableResponseBody
	GetData() *GetTenantCheckAvailableResponseBodyData
	SetRequestId(v string) *GetTenantCheckAvailableResponseBody
	GetRequestId() *string
}

type GetTenantCheckAvailableResponseBody struct {
	// The returned data.
	Data *GetTenantCheckAvailableResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request. Alibaba Cloud generates a unique identifier for each request. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// 69BFFCDE-37D6-5A49-A8BC-BB03AC83****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTenantCheckAvailableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTenantCheckAvailableResponseBody) GoString() string {
	return s.String()
}

func (s *GetTenantCheckAvailableResponseBody) GetData() *GetTenantCheckAvailableResponseBodyData {
	return s.Data
}

func (s *GetTenantCheckAvailableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTenantCheckAvailableResponseBody) SetData(v *GetTenantCheckAvailableResponseBodyData) *GetTenantCheckAvailableResponseBody {
	s.Data = v
	return s
}

func (s *GetTenantCheckAvailableResponseBody) SetRequestId(v string) *GetTenantCheckAvailableResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTenantCheckAvailableResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTenantCheckAvailableResponseBodyData struct {
	// The timestamp of the next time when a one-click scan can be submitted.
	//
	// example:
	//
	// 1725530005357
	NextScanTime *int64 `json:"NextScanTime,omitempty" xml:"NextScanTime,omitempty"`
	// The current status of the one-click scan. Valid values:
	//
	// - 0: The one-click scan can be submitted.
	//
	// - 1: The current task is not complete. The scan cannot be submitted.
	//
	// - 2: The free scan quota for this week has been used. Wait until the next free scan time.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetTenantCheckAvailableResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTenantCheckAvailableResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTenantCheckAvailableResponseBodyData) GetNextScanTime() *int64 {
	return s.NextScanTime
}

func (s *GetTenantCheckAvailableResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *GetTenantCheckAvailableResponseBodyData) SetNextScanTime(v int64) *GetTenantCheckAvailableResponseBodyData {
	s.NextScanTime = &v
	return s
}

func (s *GetTenantCheckAvailableResponseBodyData) SetStatus(v int32) *GetTenantCheckAvailableResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetTenantCheckAvailableResponseBodyData) Validate() error {
	return dara.Validate(s)
}
