// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkitemCommentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteFlag(v bool) *DeleteWorkitemCommentResponseBody
	GetDeleteFlag() *bool
	SetErrorCode(v string) *DeleteWorkitemCommentResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *DeleteWorkitemCommentResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *DeleteWorkitemCommentResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DeleteWorkitemCommentResponseBody
	GetSuccess() *string
}

type DeleteWorkitemCommentResponseBody struct {
	// example:
	//
	// true/false
	DeleteFlag *bool `json:"deleteFlag,omitempty" xml:"deleteFlag,omitempty"`
	// example:
	//
	// success
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// error
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteWorkitemCommentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkitemCommentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkitemCommentResponseBody) GetDeleteFlag() *bool {
	return s.DeleteFlag
}

func (s *DeleteWorkitemCommentResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteWorkitemCommentResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *DeleteWorkitemCommentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWorkitemCommentResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DeleteWorkitemCommentResponseBody) SetDeleteFlag(v bool) *DeleteWorkitemCommentResponseBody {
	s.DeleteFlag = &v
	return s
}

func (s *DeleteWorkitemCommentResponseBody) SetErrorCode(v string) *DeleteWorkitemCommentResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteWorkitemCommentResponseBody) SetErrorMsg(v string) *DeleteWorkitemCommentResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DeleteWorkitemCommentResponseBody) SetRequestId(v string) *DeleteWorkitemCommentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWorkitemCommentResponseBody) SetSuccess(v string) *DeleteWorkitemCommentResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteWorkitemCommentResponseBody) Validate() error {
	return dara.Validate(s)
}
