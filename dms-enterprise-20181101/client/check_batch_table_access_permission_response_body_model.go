// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckBatchTableAccessPermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*CheckBatchTableAccessPermissionResponseBodyData) *CheckBatchTableAccessPermissionResponseBody
	GetData() []*CheckBatchTableAccessPermissionResponseBodyData
	SetErrorCode(v string) *CheckBatchTableAccessPermissionResponseBody
	GetErrorCode() *string
	SetRequestId(v string) *CheckBatchTableAccessPermissionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CheckBatchTableAccessPermissionResponseBody
	GetSuccess() *bool
}

type CheckBatchTableAccessPermissionResponseBody struct {
	Data []*CheckBatchTableAccessPermissionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// B43AD641-49C2-5299-9E06-1B37EC1B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckBatchTableAccessPermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckBatchTableAccessPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *CheckBatchTableAccessPermissionResponseBody) GetData() []*CheckBatchTableAccessPermissionResponseBodyData {
	return s.Data
}

func (s *CheckBatchTableAccessPermissionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CheckBatchTableAccessPermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckBatchTableAccessPermissionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CheckBatchTableAccessPermissionResponseBody) SetData(v []*CheckBatchTableAccessPermissionResponseBodyData) *CheckBatchTableAccessPermissionResponseBody {
	s.Data = v
	return s
}

func (s *CheckBatchTableAccessPermissionResponseBody) SetErrorCode(v string) *CheckBatchTableAccessPermissionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CheckBatchTableAccessPermissionResponseBody) SetRequestId(v string) *CheckBatchTableAccessPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckBatchTableAccessPermissionResponseBody) SetSuccess(v bool) *CheckBatchTableAccessPermissionResponseBody {
	s.Success = &v
	return s
}

func (s *CheckBatchTableAccessPermissionResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CheckBatchTableAccessPermissionResponseBodyData struct {
	// example:
	//
	// The productKey is empty.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// True
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// tab_add_teacher_record
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s CheckBatchTableAccessPermissionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CheckBatchTableAccessPermissionResponseBodyData) GoString() string {
	return s.String()
}

func (s *CheckBatchTableAccessPermissionResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CheckBatchTableAccessPermissionResponseBodyData) GetSuccess() *string {
	return s.Success
}

func (s *CheckBatchTableAccessPermissionResponseBodyData) GetTableName() *string {
	return s.TableName
}

func (s *CheckBatchTableAccessPermissionResponseBodyData) SetErrorMessage(v string) *CheckBatchTableAccessPermissionResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *CheckBatchTableAccessPermissionResponseBodyData) SetSuccess(v string) *CheckBatchTableAccessPermissionResponseBodyData {
	s.Success = &v
	return s
}

func (s *CheckBatchTableAccessPermissionResponseBodyData) SetTableName(v string) *CheckBatchTableAccessPermissionResponseBodyData {
	s.TableName = &v
	return s
}

func (s *CheckBatchTableAccessPermissionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
