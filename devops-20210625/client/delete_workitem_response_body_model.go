// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkitemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteFlag(v bool) *DeleteWorkitemResponseBody
	GetDeleteFlag() *bool
	SetErrorCode(v string) *DeleteWorkitemResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *DeleteWorkitemResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *DeleteWorkitemResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteWorkitemResponseBody
	GetSuccess() *bool
}

type DeleteWorkitemResponseBody struct {
	// example:
	//
	// true/false
	DeleteFlag *bool `json:"deleteFlag,omitempty" xml:"deleteFlag,omitempty"`
	// example:
	//
	// InvalidTagGroup.IdNotFound
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
	// true/false
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteWorkitemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkitemResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkitemResponseBody) GetDeleteFlag() *bool {
	return s.DeleteFlag
}

func (s *DeleteWorkitemResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteWorkitemResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *DeleteWorkitemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWorkitemResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteWorkitemResponseBody) SetDeleteFlag(v bool) *DeleteWorkitemResponseBody {
	s.DeleteFlag = &v
	return s
}

func (s *DeleteWorkitemResponseBody) SetErrorCode(v string) *DeleteWorkitemResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteWorkitemResponseBody) SetErrorMsg(v string) *DeleteWorkitemResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DeleteWorkitemResponseBody) SetRequestId(v string) *DeleteWorkitemResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWorkitemResponseBody) SetSuccess(v bool) *DeleteWorkitemResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteWorkitemResponseBody) Validate() error {
	return dara.Validate(s)
}
