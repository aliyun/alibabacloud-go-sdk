// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSmoothExpandPreCheckTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *SubmitSmoothExpandPreCheckTaskRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *SubmitSmoothExpandPreCheckTaskRequest
	GetDrdsInstanceId() *string
}

type SubmitSmoothExpandPreCheckTaskRequest struct {
	// The name of the PolarDB-X 1.0 database.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds823s4esd
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s SubmitSmoothExpandPreCheckTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmoothExpandPreCheckTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitSmoothExpandPreCheckTaskRequest) GetDbName() *string {
	return s.DbName
}

func (s *SubmitSmoothExpandPreCheckTaskRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *SubmitSmoothExpandPreCheckTaskRequest) SetDbName(v string) *SubmitSmoothExpandPreCheckTaskRequest {
	s.DbName = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckTaskRequest) SetDrdsInstanceId(v string) *SubmitSmoothExpandPreCheckTaskRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckTaskRequest) Validate() error {
	return dara.Validate(s)
}
