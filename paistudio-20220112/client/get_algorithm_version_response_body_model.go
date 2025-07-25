// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlgorithmVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithmId(v string) *GetAlgorithmVersionResponseBody
	GetAlgorithmId() *string
	SetAlgorithmName(v string) *GetAlgorithmVersionResponseBody
	GetAlgorithmName() *string
	SetAlgorithmProvider(v string) *GetAlgorithmVersionResponseBody
	GetAlgorithmProvider() *string
	SetAlgorithmSpec(v *AlgorithmSpec) *GetAlgorithmVersionResponseBody
	GetAlgorithmSpec() *AlgorithmSpec
	SetAlgorithmVersion(v string) *GetAlgorithmVersionResponseBody
	GetAlgorithmVersion() *string
	SetGmtCreateTime(v string) *GetAlgorithmVersionResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *GetAlgorithmVersionResponseBody
	GetGmtModifiedTime() *string
	SetTenantId(v string) *GetAlgorithmVersionResponseBody
	GetTenantId() *string
	SetUserId(v string) *GetAlgorithmVersionResponseBody
	GetUserId() *string
}

type GetAlgorithmVersionResponseBody struct {
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
	AlgorithmProvider *string        `json:"AlgorithmProvider,omitempty" xml:"AlgorithmProvider,omitempty"`
	AlgorithmSpec     *AlgorithmSpec `json:"AlgorithmSpec,omitempty" xml:"AlgorithmSpec,omitempty"`
	// example:
	//
	// v0.0.1
	AlgorithmVersion *string `json:"AlgorithmVersion,omitempty" xml:"AlgorithmVersion,omitempty"`
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
	// 123456789
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// 123456789
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetAlgorithmVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAlgorithmVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlgorithmVersionResponseBody) GetAlgorithmId() *string {
	return s.AlgorithmId
}

func (s *GetAlgorithmVersionResponseBody) GetAlgorithmName() *string {
	return s.AlgorithmName
}

func (s *GetAlgorithmVersionResponseBody) GetAlgorithmProvider() *string {
	return s.AlgorithmProvider
}

func (s *GetAlgorithmVersionResponseBody) GetAlgorithmSpec() *AlgorithmSpec {
	return s.AlgorithmSpec
}

func (s *GetAlgorithmVersionResponseBody) GetAlgorithmVersion() *string {
	return s.AlgorithmVersion
}

func (s *GetAlgorithmVersionResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetAlgorithmVersionResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetAlgorithmVersionResponseBody) GetTenantId() *string {
	return s.TenantId
}

func (s *GetAlgorithmVersionResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *GetAlgorithmVersionResponseBody) SetAlgorithmId(v string) *GetAlgorithmVersionResponseBody {
	s.AlgorithmId = &v
	return s
}

func (s *GetAlgorithmVersionResponseBody) SetAlgorithmName(v string) *GetAlgorithmVersionResponseBody {
	s.AlgorithmName = &v
	return s
}

func (s *GetAlgorithmVersionResponseBody) SetAlgorithmProvider(v string) *GetAlgorithmVersionResponseBody {
	s.AlgorithmProvider = &v
	return s
}

func (s *GetAlgorithmVersionResponseBody) SetAlgorithmSpec(v *AlgorithmSpec) *GetAlgorithmVersionResponseBody {
	s.AlgorithmSpec = v
	return s
}

func (s *GetAlgorithmVersionResponseBody) SetAlgorithmVersion(v string) *GetAlgorithmVersionResponseBody {
	s.AlgorithmVersion = &v
	return s
}

func (s *GetAlgorithmVersionResponseBody) SetGmtCreateTime(v string) *GetAlgorithmVersionResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetAlgorithmVersionResponseBody) SetGmtModifiedTime(v string) *GetAlgorithmVersionResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetAlgorithmVersionResponseBody) SetTenantId(v string) *GetAlgorithmVersionResponseBody {
	s.TenantId = &v
	return s
}

func (s *GetAlgorithmVersionResponseBody) SetUserId(v string) *GetAlgorithmVersionResponseBody {
	s.UserId = &v
	return s
}

func (s *GetAlgorithmVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
