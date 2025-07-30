// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEveryOneSellsFormListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEveryOneSellsFormListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEveryOneSellsFormListResponse
	GetStatusCode() *int32
	SetBody(v []*GetEveryOneSellsFormListResponseBody) *GetEveryOneSellsFormListResponse
	GetBody() []*GetEveryOneSellsFormListResponseBody
}

type GetEveryOneSellsFormListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       []*GetEveryOneSellsFormListResponseBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s GetEveryOneSellsFormListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEveryOneSellsFormListResponse) GoString() string {
	return s.String()
}

func (s *GetEveryOneSellsFormListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEveryOneSellsFormListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEveryOneSellsFormListResponse) GetBody() []*GetEveryOneSellsFormListResponseBody {
	return s.Body
}

func (s *GetEveryOneSellsFormListResponse) SetHeaders(v map[string]*string) *GetEveryOneSellsFormListResponse {
	s.Headers = v
	return s
}

func (s *GetEveryOneSellsFormListResponse) SetStatusCode(v int32) *GetEveryOneSellsFormListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEveryOneSellsFormListResponse) SetBody(v []*GetEveryOneSellsFormListResponseBody) *GetEveryOneSellsFormListResponse {
	s.Body = v
	return s
}

func (s *GetEveryOneSellsFormListResponse) Validate() error {
	return dara.Validate(s)
}

type GetEveryOneSellsFormListResponseBody struct {
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// virtualMainDepartment
	VirtualMainDepartment *string `json:"VirtualMainDepartment,omitempty" xml:"VirtualMainDepartment,omitempty"`
	// example:
	//
	// virtualDepartmentName
	VirtualDepartmentName *string `json:"VirtualDepartmentName,omitempty" xml:"VirtualDepartmentName,omitempty"`
	// example:
	//
	// virtualDepartmentCode
	VirtualDepartmentCode *string `json:"VirtualDepartmentCode,omitempty" xml:"VirtualDepartmentCode,omitempty"`
	// example:
	//
	// virtualParentDepartment
	VirtualParentDepartment *string `json:"VirtualParentDepartment,omitempty" xml:"VirtualParentDepartment,omitempty"`
	// example:
	//
	// virtualDepartmentStatus
	VirtualDepartmentStatus *string `json:"VirtualDepartmentStatus,omitempty" xml:"VirtualDepartmentStatus,omitempty"`
	// example:
	//
	// 1234
	DingId *string `json:"DingId,omitempty" xml:"DingId,omitempty"`
	// example:
	//
	// work
	EmpStatus *string `json:"EmpStatus,omitempty" xml:"EmpStatus,omitempty"`
}

func (s GetEveryOneSellsFormListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEveryOneSellsFormListResponseBody) GoString() string {
	return s.String()
}

func (s *GetEveryOneSellsFormListResponseBody) GetId() *int64 {
	return s.Id
}

func (s *GetEveryOneSellsFormListResponseBody) GetVirtualMainDepartment() *string {
	return s.VirtualMainDepartment
}

func (s *GetEveryOneSellsFormListResponseBody) GetVirtualDepartmentName() *string {
	return s.VirtualDepartmentName
}

func (s *GetEveryOneSellsFormListResponseBody) GetVirtualDepartmentCode() *string {
	return s.VirtualDepartmentCode
}

func (s *GetEveryOneSellsFormListResponseBody) GetVirtualParentDepartment() *string {
	return s.VirtualParentDepartment
}

func (s *GetEveryOneSellsFormListResponseBody) GetVirtualDepartmentStatus() *string {
	return s.VirtualDepartmentStatus
}

func (s *GetEveryOneSellsFormListResponseBody) GetDingId() *string {
	return s.DingId
}

func (s *GetEveryOneSellsFormListResponseBody) GetEmpStatus() *string {
	return s.EmpStatus
}

func (s *GetEveryOneSellsFormListResponseBody) SetId(v int64) *GetEveryOneSellsFormListResponseBody {
	s.Id = &v
	return s
}

func (s *GetEveryOneSellsFormListResponseBody) SetVirtualMainDepartment(v string) *GetEveryOneSellsFormListResponseBody {
	s.VirtualMainDepartment = &v
	return s
}

func (s *GetEveryOneSellsFormListResponseBody) SetVirtualDepartmentName(v string) *GetEveryOneSellsFormListResponseBody {
	s.VirtualDepartmentName = &v
	return s
}

func (s *GetEveryOneSellsFormListResponseBody) SetVirtualDepartmentCode(v string) *GetEveryOneSellsFormListResponseBody {
	s.VirtualDepartmentCode = &v
	return s
}

func (s *GetEveryOneSellsFormListResponseBody) SetVirtualParentDepartment(v string) *GetEveryOneSellsFormListResponseBody {
	s.VirtualParentDepartment = &v
	return s
}

func (s *GetEveryOneSellsFormListResponseBody) SetVirtualDepartmentStatus(v string) *GetEveryOneSellsFormListResponseBody {
	s.VirtualDepartmentStatus = &v
	return s
}

func (s *GetEveryOneSellsFormListResponseBody) SetDingId(v string) *GetEveryOneSellsFormListResponseBody {
	s.DingId = &v
	return s
}

func (s *GetEveryOneSellsFormListResponseBody) SetEmpStatus(v string) *GetEveryOneSellsFormListResponseBody {
	s.EmpStatus = &v
	return s
}

func (s *GetEveryOneSellsFormListResponseBody) Validate() error {
	return dara.Validate(s)
}
