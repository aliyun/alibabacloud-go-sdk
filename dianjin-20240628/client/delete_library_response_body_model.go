// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLibraryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *DeleteLibraryResponseBody
	GetErrCode() *string
	SetMessage(v string) *DeleteLibraryResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteLibraryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteLibraryResponseBody
	GetSuccess() *bool
}

type DeleteLibraryResponseBody struct {
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 30F6AD44-F078-540D-B5A5-1E519C8E9E6D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteLibraryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLibraryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLibraryResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DeleteLibraryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteLibraryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLibraryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteLibraryResponseBody) SetErrCode(v string) *DeleteLibraryResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteLibraryResponseBody) SetMessage(v string) *DeleteLibraryResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteLibraryResponseBody) SetRequestId(v string) *DeleteLibraryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLibraryResponseBody) SetSuccess(v bool) *DeleteLibraryResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteLibraryResponseBody) Validate() error {
	return dara.Validate(s)
}
