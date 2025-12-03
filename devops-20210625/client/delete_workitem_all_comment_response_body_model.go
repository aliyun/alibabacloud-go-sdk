// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkitemAllCommentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteFlag(v bool) *DeleteWorkitemAllCommentResponseBody
	GetDeleteFlag() *bool
	SetErrorCode(v string) *DeleteWorkitemAllCommentResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *DeleteWorkitemAllCommentResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *DeleteWorkitemAllCommentResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DeleteWorkitemAllCommentResponseBody
	GetSuccess() *string
}

type DeleteWorkitemAllCommentResponseBody struct {
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
	// A7586FEB-E48D-5579-983F-74981FBFF627
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteWorkitemAllCommentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkitemAllCommentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkitemAllCommentResponseBody) GetDeleteFlag() *bool {
	return s.DeleteFlag
}

func (s *DeleteWorkitemAllCommentResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteWorkitemAllCommentResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *DeleteWorkitemAllCommentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWorkitemAllCommentResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DeleteWorkitemAllCommentResponseBody) SetDeleteFlag(v bool) *DeleteWorkitemAllCommentResponseBody {
	s.DeleteFlag = &v
	return s
}

func (s *DeleteWorkitemAllCommentResponseBody) SetErrorCode(v string) *DeleteWorkitemAllCommentResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteWorkitemAllCommentResponseBody) SetErrorMsg(v string) *DeleteWorkitemAllCommentResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DeleteWorkitemAllCommentResponseBody) SetRequestId(v string) *DeleteWorkitemAllCommentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWorkitemAllCommentResponseBody) SetSuccess(v string) *DeleteWorkitemAllCommentResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteWorkitemAllCommentResponseBody) Validate() error {
	return dara.Validate(s)
}
