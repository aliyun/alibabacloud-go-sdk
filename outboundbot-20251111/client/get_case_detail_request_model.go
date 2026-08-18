// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCaseDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCaseId(v string) *GetCaseDetailRequest
	GetCaseId() *string
	SetInstanceId(v string) *GetCaseDetailRequest
	GetInstanceId() *string
	SetProductCode(v string) *GetCaseDetailRequest
	GetProductCode() *string
}

type GetCaseDetailRequest struct {
	// The case ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 35fdb65e-9c20-42af-9d89-a24068547cb6
	CaseId *string `json:"CaseId,omitempty" xml:"CaseId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4eee9bf8-1319-468f-ac82-83c50ae389f8
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The product code.
	//
	// example:
	//
	// OUTBOUND_BOT
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s GetCaseDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCaseDetailRequest) GoString() string {
	return s.String()
}

func (s *GetCaseDetailRequest) GetCaseId() *string {
	return s.CaseId
}

func (s *GetCaseDetailRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetCaseDetailRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *GetCaseDetailRequest) SetCaseId(v string) *GetCaseDetailRequest {
	s.CaseId = &v
	return s
}

func (s *GetCaseDetailRequest) SetInstanceId(v string) *GetCaseDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *GetCaseDetailRequest) SetProductCode(v string) *GetCaseDetailRequest {
	s.ProductCode = &v
	return s
}

func (s *GetCaseDetailRequest) Validate() error {
	return dara.Validate(s)
}
