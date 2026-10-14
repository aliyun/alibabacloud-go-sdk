// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgenticTableColumnResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*AgenticTableColumn) *GetAgenticTableColumnResponseBody
	GetData() []*AgenticTableColumn
	SetErrorCode(v string) *GetAgenticTableColumnResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetAgenticTableColumnResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetAgenticTableColumnResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAgenticTableColumnResponseBody
	GetSuccess() *bool
}

type GetAgenticTableColumnResponseBody struct {
	// The list of fields.
	Data []*AgenticTableColumn `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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

func (s GetAgenticTableColumnResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgenticTableColumnResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgenticTableColumnResponseBody) GetData() []*AgenticTableColumn {
	return s.Data
}

func (s *GetAgenticTableColumnResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetAgenticTableColumnResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetAgenticTableColumnResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgenticTableColumnResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAgenticTableColumnResponseBody) SetData(v []*AgenticTableColumn) *GetAgenticTableColumnResponseBody {
	s.Data = v
	return s
}

func (s *GetAgenticTableColumnResponseBody) SetErrorCode(v string) *GetAgenticTableColumnResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetAgenticTableColumnResponseBody) SetErrorMessage(v string) *GetAgenticTableColumnResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetAgenticTableColumnResponseBody) SetRequestId(v string) *GetAgenticTableColumnResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgenticTableColumnResponseBody) SetSuccess(v bool) *GetAgenticTableColumnResponseBody {
	s.Success = &v
	return s
}

func (s *GetAgenticTableColumnResponseBody) Validate() error {
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
