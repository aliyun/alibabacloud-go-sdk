// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlgorithmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithmDescription(v string) *GetAlgorithmResponseBody
	GetAlgorithmDescription() *string
	SetAlgorithmId(v string) *GetAlgorithmResponseBody
	GetAlgorithmId() *string
	SetAlgorithmName(v string) *GetAlgorithmResponseBody
	GetAlgorithmName() *string
	SetAlgorithmProvider(v string) *GetAlgorithmResponseBody
	GetAlgorithmProvider() *string
	SetDisplayName(v string) *GetAlgorithmResponseBody
	GetDisplayName() *string
	SetGmtCreateTime(v string) *GetAlgorithmResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *GetAlgorithmResponseBody
	GetGmtModifiedTime() *string
	SetRequestId(v string) *GetAlgorithmResponseBody
	GetRequestId() *string
	SetTenantId(v string) *GetAlgorithmResponseBody
	GetTenantId() *string
	SetUserId(v string) *GetAlgorithmResponseBody
	GetUserId() *string
	SetWorkspaceId(v string) *GetAlgorithmResponseBody
	GetWorkspaceId() *string
}

type GetAlgorithmResponseBody struct {
	AlgorithmDescription *string `json:"AlgorithmDescription,omitempty" xml:"AlgorithmDescription,omitempty"`
	// example:
	//
	// algo-xsldfvu1334
	AlgorithmId *string `json:"AlgorithmId,omitempty" xml:"AlgorithmId,omitempty"`
	// example:
	//
	// llm_training
	AlgorithmName *string `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	// example:
	//
	// pai
	AlgorithmProvider *string `json:"AlgorithmProvider,omitempty" xml:"AlgorithmProvider,omitempty"`
	// example:
	//
	// llm_training
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 2024-07-10T11:49:47Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2024-07-10T11:49:47Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 123456789
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// 123456789
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 12345
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetAlgorithmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAlgorithmResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlgorithmResponseBody) GetAlgorithmDescription() *string {
	return s.AlgorithmDescription
}

func (s *GetAlgorithmResponseBody) GetAlgorithmId() *string {
	return s.AlgorithmId
}

func (s *GetAlgorithmResponseBody) GetAlgorithmName() *string {
	return s.AlgorithmName
}

func (s *GetAlgorithmResponseBody) GetAlgorithmProvider() *string {
	return s.AlgorithmProvider
}

func (s *GetAlgorithmResponseBody) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetAlgorithmResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetAlgorithmResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetAlgorithmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAlgorithmResponseBody) GetTenantId() *string {
	return s.TenantId
}

func (s *GetAlgorithmResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *GetAlgorithmResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetAlgorithmResponseBody) SetAlgorithmDescription(v string) *GetAlgorithmResponseBody {
	s.AlgorithmDescription = &v
	return s
}

func (s *GetAlgorithmResponseBody) SetAlgorithmId(v string) *GetAlgorithmResponseBody {
	s.AlgorithmId = &v
	return s
}

func (s *GetAlgorithmResponseBody) SetAlgorithmName(v string) *GetAlgorithmResponseBody {
	s.AlgorithmName = &v
	return s
}

func (s *GetAlgorithmResponseBody) SetAlgorithmProvider(v string) *GetAlgorithmResponseBody {
	s.AlgorithmProvider = &v
	return s
}

func (s *GetAlgorithmResponseBody) SetDisplayName(v string) *GetAlgorithmResponseBody {
	s.DisplayName = &v
	return s
}

func (s *GetAlgorithmResponseBody) SetGmtCreateTime(v string) *GetAlgorithmResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetAlgorithmResponseBody) SetGmtModifiedTime(v string) *GetAlgorithmResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetAlgorithmResponseBody) SetRequestId(v string) *GetAlgorithmResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAlgorithmResponseBody) SetTenantId(v string) *GetAlgorithmResponseBody {
	s.TenantId = &v
	return s
}

func (s *GetAlgorithmResponseBody) SetUserId(v string) *GetAlgorithmResponseBody {
	s.UserId = &v
	return s
}

func (s *GetAlgorithmResponseBody) SetWorkspaceId(v string) *GetAlgorithmResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *GetAlgorithmResponseBody) Validate() error {
	return dara.Validate(s)
}
