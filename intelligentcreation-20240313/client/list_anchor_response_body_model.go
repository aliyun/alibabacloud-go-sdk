// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAnchorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAnchorResponseBody
	GetCode() *string
	SetErrorCode(v string) *ListAnchorResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListAnchorResponseBody
	GetErrorMessage() *string
	SetList(v []*AnchorResponse) *ListAnchorResponseBody
	GetList() []*AnchorResponse
	SetRequestId(v string) *ListAnchorResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAnchorResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *ListAnchorResponseBody
	GetTotal() *int32
}

type ListAnchorResponseBody struct {
	// code
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// Deduct.DeductTaskAlreadySuccess
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// Failed to proxy flink ui request, message: An error occurred: Invalid UUID string: jobsn
	ErrorMessage *string           `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	List         []*AnchorResponse `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// requestId
	//
	// example:
	//
	// 5389BE87-571B-573C-90ED-F07C5E68760B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 10
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListAnchorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAnchorResponseBody) GoString() string {
	return s.String()
}

func (s *ListAnchorResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAnchorResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListAnchorResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListAnchorResponseBody) GetList() []*AnchorResponse {
	return s.List
}

func (s *ListAnchorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAnchorResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAnchorResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListAnchorResponseBody) SetCode(v string) *ListAnchorResponseBody {
	s.Code = &v
	return s
}

func (s *ListAnchorResponseBody) SetErrorCode(v string) *ListAnchorResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListAnchorResponseBody) SetErrorMessage(v string) *ListAnchorResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListAnchorResponseBody) SetList(v []*AnchorResponse) *ListAnchorResponseBody {
	s.List = v
	return s
}

func (s *ListAnchorResponseBody) SetRequestId(v string) *ListAnchorResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAnchorResponseBody) SetSuccess(v bool) *ListAnchorResponseBody {
	s.Success = &v
	return s
}

func (s *ListAnchorResponseBody) SetTotal(v int32) *ListAnchorResponseBody {
	s.Total = &v
	return s
}

func (s *ListAnchorResponseBody) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
