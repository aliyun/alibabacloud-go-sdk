// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileProtectClientRuleDashboardResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetFileProtectClientRuleDashboardResponseBodyData) *GetFileProtectClientRuleDashboardResponseBody
	GetData() *GetFileProtectClientRuleDashboardResponseBodyData
	SetRequestId(v string) *GetFileProtectClientRuleDashboardResponseBody
	GetRequestId() *string
}

type GetFileProtectClientRuleDashboardResponseBody struct {
	Data *GetFileProtectClientRuleDashboardResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 69BFFCDE-37D6-5A49-A8BC-BB03AC83****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFileProtectClientRuleDashboardResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFileProtectClientRuleDashboardResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileProtectClientRuleDashboardResponseBody) GetData() *GetFileProtectClientRuleDashboardResponseBodyData {
	return s.Data
}

func (s *GetFileProtectClientRuleDashboardResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFileProtectClientRuleDashboardResponseBody) SetData(v *GetFileProtectClientRuleDashboardResponseBodyData) *GetFileProtectClientRuleDashboardResponseBody {
	s.Data = v
	return s
}

func (s *GetFileProtectClientRuleDashboardResponseBody) SetRequestId(v string) *GetFileProtectClientRuleDashboardResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFileProtectClientRuleDashboardResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFileProtectClientRuleDashboardResponseBodyData struct {
	// example:
	//
	// 5
	AuthTotal *string `json:"AuthTotal,omitempty" xml:"AuthTotal,omitempty"`
	// example:
	//
	// 4
	BindCount *int32 `json:"BindCount,omitempty" xml:"BindCount,omitempty"`
	// example:
	//
	// 3
	ProtectedDirectoriesCount *int32 `json:"ProtectedDirectoriesCount,omitempty" xml:"ProtectedDirectoriesCount,omitempty"`
	// example:
	//
	// 12
	ProtectedInstancesCount *int32 `json:"ProtectedInstancesCount,omitempty" xml:"ProtectedInstancesCount,omitempty"`
}

func (s GetFileProtectClientRuleDashboardResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetFileProtectClientRuleDashboardResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFileProtectClientRuleDashboardResponseBodyData) GetAuthTotal() *string {
	return s.AuthTotal
}

func (s *GetFileProtectClientRuleDashboardResponseBodyData) GetBindCount() *int32 {
	return s.BindCount
}

func (s *GetFileProtectClientRuleDashboardResponseBodyData) GetProtectedDirectoriesCount() *int32 {
	return s.ProtectedDirectoriesCount
}

func (s *GetFileProtectClientRuleDashboardResponseBodyData) GetProtectedInstancesCount() *int32 {
	return s.ProtectedInstancesCount
}

func (s *GetFileProtectClientRuleDashboardResponseBodyData) SetAuthTotal(v string) *GetFileProtectClientRuleDashboardResponseBodyData {
	s.AuthTotal = &v
	return s
}

func (s *GetFileProtectClientRuleDashboardResponseBodyData) SetBindCount(v int32) *GetFileProtectClientRuleDashboardResponseBodyData {
	s.BindCount = &v
	return s
}

func (s *GetFileProtectClientRuleDashboardResponseBodyData) SetProtectedDirectoriesCount(v int32) *GetFileProtectClientRuleDashboardResponseBodyData {
	s.ProtectedDirectoriesCount = &v
	return s
}

func (s *GetFileProtectClientRuleDashboardResponseBodyData) SetProtectedInstancesCount(v int32) *GetFileProtectClientRuleDashboardResponseBodyData {
	s.ProtectedInstancesCount = &v
	return s
}

func (s *GetFileProtectClientRuleDashboardResponseBodyData) Validate() error {
	return dara.Validate(s)
}
