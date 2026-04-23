// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckInventoryJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *KnowledgeJobInfoVO) *CheckInventoryJobResponseBody
	GetData() *KnowledgeJobInfoVO
	SetErrorCode(v string) *CheckInventoryJobResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CheckInventoryJobResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CheckInventoryJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CheckInventoryJobResponseBody
	GetSuccess() *bool
}

type CheckInventoryJobResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	Data *KnowledgeJobInfoVO `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s CheckInventoryJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckInventoryJobResponseBody) GoString() string {
	return s.String()
}

func (s *CheckInventoryJobResponseBody) GetData() *KnowledgeJobInfoVO {
	return s.Data
}

func (s *CheckInventoryJobResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CheckInventoryJobResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CheckInventoryJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckInventoryJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CheckInventoryJobResponseBody) SetData(v *KnowledgeJobInfoVO) *CheckInventoryJobResponseBody {
	s.Data = v
	return s
}

func (s *CheckInventoryJobResponseBody) SetErrorCode(v string) *CheckInventoryJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CheckInventoryJobResponseBody) SetErrorMessage(v string) *CheckInventoryJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CheckInventoryJobResponseBody) SetRequestId(v string) *CheckInventoryJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckInventoryJobResponseBody) SetSuccess(v bool) *CheckInventoryJobResponseBody {
	s.Success = &v
	return s
}

func (s *CheckInventoryJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
