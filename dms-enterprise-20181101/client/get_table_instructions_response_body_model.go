// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableInstructionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *TableInstructionsVO) *GetTableInstructionsResponseBody
	GetData() *TableInstructionsVO
	SetErrorCode(v string) *GetTableInstructionsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetTableInstructionsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetTableInstructionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTableInstructionsResponseBody
	GetSuccess() *bool
}

type GetTableInstructionsResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 1001
	Data *TableInstructionsVO `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTableInstructionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTableInstructionsResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableInstructionsResponseBody) GetData() *TableInstructionsVO {
	return s.Data
}

func (s *GetTableInstructionsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetTableInstructionsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetTableInstructionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTableInstructionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTableInstructionsResponseBody) SetData(v *TableInstructionsVO) *GetTableInstructionsResponseBody {
	s.Data = v
	return s
}

func (s *GetTableInstructionsResponseBody) SetErrorCode(v string) *GetTableInstructionsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTableInstructionsResponseBody) SetErrorMessage(v string) *GetTableInstructionsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetTableInstructionsResponseBody) SetRequestId(v string) *GetTableInstructionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableInstructionsResponseBody) SetSuccess(v bool) *GetTableInstructionsResponseBody {
	s.Success = &v
	return s
}

func (s *GetTableInstructionsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
