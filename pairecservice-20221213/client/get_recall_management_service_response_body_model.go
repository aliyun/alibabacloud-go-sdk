// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecallManagementServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentRecallManagementServiceVersionId(v string) *GetRecallManagementServiceResponseBody
	GetCurrentRecallManagementServiceVersionId() *string
	SetCurrentRecallManagementServiceVersionName(v string) *GetRecallManagementServiceResponseBody
	GetCurrentRecallManagementServiceVersionName() *string
	SetDescription(v string) *GetRecallManagementServiceResponseBody
	GetDescription() *string
	SetGmtCreateTime(v string) *GetRecallManagementServiceResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *GetRecallManagementServiceResponseBody
	GetGmtModifiedTime() *string
	SetName(v string) *GetRecallManagementServiceResponseBody
	GetName() *string
	SetRecallManagementServiceId(v string) *GetRecallManagementServiceResponseBody
	GetRecallManagementServiceId() *string
	SetRequestId(v string) *GetRecallManagementServiceResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetRecallManagementServiceResponseBody
	GetStatus() *string
}

type GetRecallManagementServiceResponseBody struct {
	// example:
	//
	// 1
	CurrentRecallManagementServiceVersionId *string `json:"CurrentRecallManagementServiceVersionId,omitempty" xml:"CurrentRecallManagementServiceVersionId,omitempty"`
	// example:
	//
	// version-1
	CurrentRecallManagementServiceVersionName *string `json:"CurrentRecallManagementServiceVersionName,omitempty" xml:"CurrentRecallManagementServiceVersionName,omitempty"`
	// example:
	//
	// this is a test recall
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// hot_group_recall
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 3
	RecallManagementServiceId *string `json:"RecallManagementServiceId,omitempty" xml:"RecallManagementServiceId,omitempty"`
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetRecallManagementServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRecallManagementServiceResponseBody) GoString() string {
	return s.String()
}

func (s *GetRecallManagementServiceResponseBody) GetCurrentRecallManagementServiceVersionId() *string {
	return s.CurrentRecallManagementServiceVersionId
}

func (s *GetRecallManagementServiceResponseBody) GetCurrentRecallManagementServiceVersionName() *string {
	return s.CurrentRecallManagementServiceVersionName
}

func (s *GetRecallManagementServiceResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetRecallManagementServiceResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetRecallManagementServiceResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetRecallManagementServiceResponseBody) GetName() *string {
	return s.Name
}

func (s *GetRecallManagementServiceResponseBody) GetRecallManagementServiceId() *string {
	return s.RecallManagementServiceId
}

func (s *GetRecallManagementServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRecallManagementServiceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetRecallManagementServiceResponseBody) SetCurrentRecallManagementServiceVersionId(v string) *GetRecallManagementServiceResponseBody {
	s.CurrentRecallManagementServiceVersionId = &v
	return s
}

func (s *GetRecallManagementServiceResponseBody) SetCurrentRecallManagementServiceVersionName(v string) *GetRecallManagementServiceResponseBody {
	s.CurrentRecallManagementServiceVersionName = &v
	return s
}

func (s *GetRecallManagementServiceResponseBody) SetDescription(v string) *GetRecallManagementServiceResponseBody {
	s.Description = &v
	return s
}

func (s *GetRecallManagementServiceResponseBody) SetGmtCreateTime(v string) *GetRecallManagementServiceResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetRecallManagementServiceResponseBody) SetGmtModifiedTime(v string) *GetRecallManagementServiceResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetRecallManagementServiceResponseBody) SetName(v string) *GetRecallManagementServiceResponseBody {
	s.Name = &v
	return s
}

func (s *GetRecallManagementServiceResponseBody) SetRecallManagementServiceId(v string) *GetRecallManagementServiceResponseBody {
	s.RecallManagementServiceId = &v
	return s
}

func (s *GetRecallManagementServiceResponseBody) SetRequestId(v string) *GetRecallManagementServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRecallManagementServiceResponseBody) SetStatus(v string) *GetRecallManagementServiceResponseBody {
	s.Status = &v
	return s
}

func (s *GetRecallManagementServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
