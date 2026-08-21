// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceHealthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *ListInstanceHealthRequest
	GetXDebugId() *string
	SetCluster(v string) *ListInstanceHealthRequest
	GetCluster() *string
	SetCurrent(v int32) *ListInstanceHealthRequest
	GetCurrent() *int32
	SetEnd(v float32) *ListInstanceHealthRequest
	GetEnd() *float32
	SetInstance(v string) *ListInstanceHealthRequest
	GetInstance() *string
	SetPageSize(v int32) *ListInstanceHealthRequest
	GetPageSize() *int32
	SetStart(v float32) *ListInstanceHealthRequest
	GetStart() *float32
	SetXSysomInvokeSource(v string) *ListInstanceHealthRequest
	GetXSysomInvokeSource() *string
}

type ListInstanceHealthRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// 2dc6f3be-bf24-4070-8d3a-c03c26850532
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	Current *int32 `json:"current,omitempty" xml:"current,omitempty"`
	// The end time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1725801327754
	End *float32 `json:"end,omitempty" xml:"end,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-wz9d00ut2ska3mlyhn6j
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// The number of entries per page. Default value: 5. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1725797727754
	Start              *float32 `json:"start,omitempty" xml:"start,omitempty"`
	XSysomInvokeSource *string  `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s ListInstanceHealthRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceHealthRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceHealthRequest) GetXDebugId() *string {
	return s.XDebugId
}

func (s *ListInstanceHealthRequest) GetCluster() *string {
	return s.Cluster
}

func (s *ListInstanceHealthRequest) GetCurrent() *int32 {
	return s.Current
}

func (s *ListInstanceHealthRequest) GetEnd() *float32 {
	return s.End
}

func (s *ListInstanceHealthRequest) GetInstance() *string {
	return s.Instance
}

func (s *ListInstanceHealthRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstanceHealthRequest) GetStart() *float32 {
	return s.Start
}

func (s *ListInstanceHealthRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *ListInstanceHealthRequest) SetXDebugId(v string) *ListInstanceHealthRequest {
	s.XDebugId = &v
	return s
}

func (s *ListInstanceHealthRequest) SetCluster(v string) *ListInstanceHealthRequest {
	s.Cluster = &v
	return s
}

func (s *ListInstanceHealthRequest) SetCurrent(v int32) *ListInstanceHealthRequest {
	s.Current = &v
	return s
}

func (s *ListInstanceHealthRequest) SetEnd(v float32) *ListInstanceHealthRequest {
	s.End = &v
	return s
}

func (s *ListInstanceHealthRequest) SetInstance(v string) *ListInstanceHealthRequest {
	s.Instance = &v
	return s
}

func (s *ListInstanceHealthRequest) SetPageSize(v int32) *ListInstanceHealthRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstanceHealthRequest) SetStart(v float32) *ListInstanceHealthRequest {
	s.Start = &v
	return s
}

func (s *ListInstanceHealthRequest) SetXSysomInvokeSource(v string) *ListInstanceHealthRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *ListInstanceHealthRequest) Validate() error {
	return dara.Validate(s)
}
