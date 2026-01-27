// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImportTaskValidationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDetail(v string) *DescribeImportTaskValidationResponseBody
	GetDetail() *string
	SetRequestId(v string) *DescribeImportTaskValidationResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeImportTaskValidationResponseBody
	GetStatus() *string
	SetSuccess(v bool) *DescribeImportTaskValidationResponseBody
	GetSuccess() *bool
}

type DescribeImportTaskValidationResponseBody struct {
	// example:
	//
	// {"ValidateAction": "Detail"}
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 3E36DB6E-AE3B-53B6-A703-85F883FD1B2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// COMPLETED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeImportTaskValidationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImportTaskValidationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImportTaskValidationResponseBody) GetDetail() *string {
	return s.Detail
}

func (s *DescribeImportTaskValidationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImportTaskValidationResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeImportTaskValidationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeImportTaskValidationResponseBody) SetDetail(v string) *DescribeImportTaskValidationResponseBody {
	s.Detail = &v
	return s
}

func (s *DescribeImportTaskValidationResponseBody) SetRequestId(v string) *DescribeImportTaskValidationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImportTaskValidationResponseBody) SetStatus(v string) *DescribeImportTaskValidationResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeImportTaskValidationResponseBody) SetSuccess(v bool) *DescribeImportTaskValidationResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeImportTaskValidationResponseBody) Validate() error {
	return dara.Validate(s)
}
