// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSqlAuditInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeSqlAuditInfoResponseBodyData) *DescribeSqlAuditInfoResponseBody
	GetData() *DescribeSqlAuditInfoResponseBodyData
	SetRequestId(v string) *DescribeSqlAuditInfoResponseBody
	GetRequestId() *string
}

type DescribeSqlAuditInfoResponseBody struct {
	Data *DescribeSqlAuditInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// DC3ABA3E-0F8A-4596-9104-F5155C******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSqlAuditInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlAuditInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSqlAuditInfoResponseBody) GetData() *DescribeSqlAuditInfoResponseBodyData {
	return s.Data
}

func (s *DescribeSqlAuditInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSqlAuditInfoResponseBody) SetData(v *DescribeSqlAuditInfoResponseBodyData) *DescribeSqlAuditInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSqlAuditInfoResponseBody) SetRequestId(v string) *DescribeSqlAuditInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSqlAuditInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSqlAuditInfoResponseBodyData struct {
	// example:
	//
	// true
	IsEnabled *bool `json:"IsEnabled,omitempty" xml:"IsEnabled,omitempty"`
	// example:
	//
	// polardbx-sqlaudit-log
	SLSLogStore *string `json:"SLSLogStore,omitempty" xml:"SLSLogStore,omitempty"`
	// example:
	//
	// polardbx-sqlaudit-cn-hangzhou-123456789
	SLSProject *string `json:"SLSProject,omitempty" xml:"SLSProject,omitempty"`
}

func (s DescribeSqlAuditInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlAuditInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSqlAuditInfoResponseBodyData) GetIsEnabled() *bool {
	return s.IsEnabled
}

func (s *DescribeSqlAuditInfoResponseBodyData) GetSLSLogStore() *string {
	return s.SLSLogStore
}

func (s *DescribeSqlAuditInfoResponseBodyData) GetSLSProject() *string {
	return s.SLSProject
}

func (s *DescribeSqlAuditInfoResponseBodyData) SetIsEnabled(v bool) *DescribeSqlAuditInfoResponseBodyData {
	s.IsEnabled = &v
	return s
}

func (s *DescribeSqlAuditInfoResponseBodyData) SetSLSLogStore(v string) *DescribeSqlAuditInfoResponseBodyData {
	s.SLSLogStore = &v
	return s
}

func (s *DescribeSqlAuditInfoResponseBodyData) SetSLSProject(v string) *DescribeSqlAuditInfoResponseBodyData {
	s.SLSProject = &v
	return s
}

func (s *DescribeSqlAuditInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
