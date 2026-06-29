// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubtaskItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetSubtaskItemResponseBody
	GetCode() *int32
	SetDetails(v string) *GetSubtaskItemResponseBody
	GetDetails() *string
	SetErrorCode(v string) *GetSubtaskItemResponseBody
	GetErrorCode() *string
	SetItem(v *SubtaskItemDetail) *GetSubtaskItemResponseBody
	GetItem() *SubtaskItemDetail
	SetMessage(v string) *GetSubtaskItemResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSubtaskItemResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSubtaskItemResponseBody
	GetSuccess() *bool
}

type GetSubtaskItemResponseBody struct {
	// Return code. The default value is 0, indicating normal execution.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Details.
	//
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// Error code.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Data item.
	Item *SubtaskItemDetail `json:"Item,omitempty" xml:"Item,omitempty"`
	// Return message of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded. Valid values:
	//
	// - true: The request succeeded.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSubtaskItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSubtaskItemResponseBody) GoString() string {
	return s.String()
}

func (s *GetSubtaskItemResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetSubtaskItemResponseBody) GetDetails() *string {
	return s.Details
}

func (s *GetSubtaskItemResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetSubtaskItemResponseBody) GetItem() *SubtaskItemDetail {
	return s.Item
}

func (s *GetSubtaskItemResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSubtaskItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSubtaskItemResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSubtaskItemResponseBody) SetCode(v int32) *GetSubtaskItemResponseBody {
	s.Code = &v
	return s
}

func (s *GetSubtaskItemResponseBody) SetDetails(v string) *GetSubtaskItemResponseBody {
	s.Details = &v
	return s
}

func (s *GetSubtaskItemResponseBody) SetErrorCode(v string) *GetSubtaskItemResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetSubtaskItemResponseBody) SetItem(v *SubtaskItemDetail) *GetSubtaskItemResponseBody {
	s.Item = v
	return s
}

func (s *GetSubtaskItemResponseBody) SetMessage(v string) *GetSubtaskItemResponseBody {
	s.Message = &v
	return s
}

func (s *GetSubtaskItemResponseBody) SetRequestId(v string) *GetSubtaskItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSubtaskItemResponseBody) SetSuccess(v bool) *GetSubtaskItemResponseBody {
	s.Success = &v
	return s
}

func (s *GetSubtaskItemResponseBody) Validate() error {
	if s.Item != nil {
		if err := s.Item.Validate(); err != nil {
			return err
		}
	}
	return nil
}
