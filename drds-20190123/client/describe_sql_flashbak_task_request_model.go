// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSqlFlashbakTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDrdsInstanceId(v string) *DescribeSqlFlashbakTaskRequest
	GetDrdsInstanceId() *string
}

type DescribeSqlFlashbakTaskRequest struct {
	// The ID of the PolarDB-X 1.0 instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds****c6vxxyzd
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeSqlFlashbakTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlFlashbakTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeSqlFlashbakTaskRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeSqlFlashbakTaskRequest) SetDrdsInstanceId(v string) *DescribeSqlFlashbakTaskRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeSqlFlashbakTaskRequest) Validate() error {
	return dara.Validate(s)
}
