// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExportMigrationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *CreateExportMigrationResponseBody
	GetData() *int64
	SetRequestId(v string) *CreateExportMigrationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateExportMigrationResponseBody
	GetSuccess() *bool
}

type CreateExportMigrationResponseBody struct {
	// The export task ID.
	//
	// example:
	//
	// 1234
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID. You can locate logs and troubleshoot issues based on the ID.
	//
	// example:
	//
	// 34267E2E-0335-1A60-A1F0-ADA530890CBA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateExportMigrationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateExportMigrationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExportMigrationResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateExportMigrationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateExportMigrationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateExportMigrationResponseBody) SetData(v int64) *CreateExportMigrationResponseBody {
	s.Data = &v
	return s
}

func (s *CreateExportMigrationResponseBody) SetRequestId(v string) *CreateExportMigrationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateExportMigrationResponseBody) SetSuccess(v bool) *CreateExportMigrationResponseBody {
	s.Success = &v
	return s
}

func (s *CreateExportMigrationResponseBody) Validate() error {
	return dara.Validate(s)
}
