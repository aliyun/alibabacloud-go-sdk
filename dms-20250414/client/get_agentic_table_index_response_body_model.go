// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgenticTableIndexResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*AgenticTableIndex) *GetAgenticTableIndexResponseBody
	GetData() []*AgenticTableIndex
	SetErrorCode(v string) *GetAgenticTableIndexResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetAgenticTableIndexResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetAgenticTableIndexResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAgenticTableIndexResponseBody
	GetSuccess() *bool
}

type GetAgenticTableIndexResponseBody struct {
	// The list of indexes.
	Data []*AgenticTableIndex `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code returned when a request exception occurs.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the call fails.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - **true**: The request is successful.
	//
	// - **false**: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAgenticTableIndexResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgenticTableIndexResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgenticTableIndexResponseBody) GetData() []*AgenticTableIndex {
	return s.Data
}

func (s *GetAgenticTableIndexResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetAgenticTableIndexResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetAgenticTableIndexResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgenticTableIndexResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAgenticTableIndexResponseBody) SetData(v []*AgenticTableIndex) *GetAgenticTableIndexResponseBody {
	s.Data = v
	return s
}

func (s *GetAgenticTableIndexResponseBody) SetErrorCode(v string) *GetAgenticTableIndexResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetAgenticTableIndexResponseBody) SetErrorMessage(v string) *GetAgenticTableIndexResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetAgenticTableIndexResponseBody) SetRequestId(v string) *GetAgenticTableIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgenticTableIndexResponseBody) SetSuccess(v bool) *GetAgenticTableIndexResponseBody {
	s.Success = &v
	return s
}

func (s *GetAgenticTableIndexResponseBody) Validate() error {
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
