// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLogSyncToSLSResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetLogSyncToSLSResponseBodyData) *GetLogSyncToSLSResponseBody
	GetData() *GetLogSyncToSLSResponseBodyData
	SetRequestId(v string) *GetLogSyncToSLSResponseBody
	GetRequestId() *string
}

type GetLogSyncToSLSResponseBody struct {
	// The returned data.
	Data *GetLogSyncToSLSResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLogSyncToSLSResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLogSyncToSLSResponseBody) GoString() string {
	return s.String()
}

func (s *GetLogSyncToSLSResponseBody) GetData() *GetLogSyncToSLSResponseBodyData {
	return s.Data
}

func (s *GetLogSyncToSLSResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLogSyncToSLSResponseBody) SetData(v *GetLogSyncToSLSResponseBodyData) *GetLogSyncToSLSResponseBody {
	s.Data = v
	return s
}

func (s *GetLogSyncToSLSResponseBody) SetRequestId(v string) *GetLogSyncToSLSResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLogSyncToSLSResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLogSyncToSLSResponseBodyData struct {
	// The log synchronization status. Valid values:
	//
	// - on: Synchronization is enabled.
	//
	// - off: Synchronization is disabled.
	//
	// example:
	//
	// off
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The Simple Log Service Logstore.
	//
	// example:
	//
	// adbmysql-audit-log
	TargetLogStore *string `json:"TargetLogStore,omitempty" xml:"TargetLogStore,omitempty"`
	// The Simple Log Service project.
	//
	// example:
	//
	// log-service-****-cn-shenzhen
	TargetProject *string `json:"TargetProject,omitempty" xml:"TargetProject,omitempty"`
}

func (s GetLogSyncToSLSResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetLogSyncToSLSResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetLogSyncToSLSResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetLogSyncToSLSResponseBodyData) GetTargetLogStore() *string {
	return s.TargetLogStore
}

func (s *GetLogSyncToSLSResponseBodyData) GetTargetProject() *string {
	return s.TargetProject
}

func (s *GetLogSyncToSLSResponseBodyData) SetStatus(v string) *GetLogSyncToSLSResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetLogSyncToSLSResponseBodyData) SetTargetLogStore(v string) *GetLogSyncToSLSResponseBodyData {
	s.TargetLogStore = &v
	return s
}

func (s *GetLogSyncToSLSResponseBodyData) SetTargetProject(v string) *GetLogSyncToSLSResponseBodyData {
	s.TargetProject = &v
	return s
}

func (s *GetLogSyncToSLSResponseBodyData) Validate() error {
	return dara.Validate(s)
}
