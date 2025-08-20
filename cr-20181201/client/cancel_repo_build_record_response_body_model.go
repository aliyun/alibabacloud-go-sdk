// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelRepoBuildRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CancelRepoBuildRecordResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *CancelRepoBuildRecordResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *CancelRepoBuildRecordResponseBody
	GetRequestId() *string
}

type CancelRepoBuildRecordResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4CE1F661-75DD-4EBD-A4AD-057B26834ABB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelRepoBuildRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelRepoBuildRecordResponseBody) GoString() string {
	return s.String()
}

func (s *CancelRepoBuildRecordResponseBody) GetCode() *string {
	return s.Code
}

func (s *CancelRepoBuildRecordResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *CancelRepoBuildRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelRepoBuildRecordResponseBody) SetCode(v string) *CancelRepoBuildRecordResponseBody {
	s.Code = &v
	return s
}

func (s *CancelRepoBuildRecordResponseBody) SetIsSuccess(v bool) *CancelRepoBuildRecordResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CancelRepoBuildRecordResponseBody) SetRequestId(v string) *CancelRepoBuildRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelRepoBuildRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
