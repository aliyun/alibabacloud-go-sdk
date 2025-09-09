// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsSqlAuditStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDrdsInstanceId(v string) *DescribeDrdsSqlAuditStatusRequest
	GetDrdsInstanceId() *string
}

type DescribeDrdsSqlAuditStatusRequest struct {
	// The ID of the PolarDB-X 1.0 instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds************
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeDrdsSqlAuditStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsSqlAuditStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsSqlAuditStatusRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeDrdsSqlAuditStatusRequest) SetDrdsInstanceId(v string) *DescribeDrdsSqlAuditStatusRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsSqlAuditStatusRequest) Validate() error {
	return dara.Validate(s)
}
