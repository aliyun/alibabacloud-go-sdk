// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIServicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ListAIServicesRequest
	GetDBInstanceId() *string
	SetPageNumber(v string) *ListAIServicesRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListAIServicesRequest
	GetPageSize() *string
	SetType(v string) *ListAIServicesRequest
	GetType() *string
}

type ListAIServicesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// drama
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListAIServicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAIServicesRequest) GoString() string {
	return s.String()
}

func (s *ListAIServicesRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ListAIServicesRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListAIServicesRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListAIServicesRequest) GetType() *string {
	return s.Type
}

func (s *ListAIServicesRequest) SetDBInstanceId(v string) *ListAIServicesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ListAIServicesRequest) SetPageNumber(v string) *ListAIServicesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAIServicesRequest) SetPageSize(v string) *ListAIServicesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAIServicesRequest) SetType(v string) *ListAIServicesRequest {
	s.Type = &v
	return s
}

func (s *ListAIServicesRequest) Validate() error {
	return dara.Validate(s)
}
