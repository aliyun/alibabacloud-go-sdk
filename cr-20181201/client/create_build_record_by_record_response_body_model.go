// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBuildRecordByRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBuildRecordId(v string) *CreateBuildRecordByRecordResponseBody
	GetBuildRecordId() *string
	SetCode(v string) *CreateBuildRecordByRecordResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *CreateBuildRecordByRecordResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *CreateBuildRecordByRecordResponseBody
	GetRequestId() *string
}

type CreateBuildRecordByRecordResponseBody struct {
	// The build record ID.
	//
	// example:
	//
	// crbr-ly77w5i3t31f****
	BuildRecordId *string `json:"BuildRecordId,omitempty" xml:"BuildRecordId,omitempty"`
	// The return code of the API operation. A value of 200 indicates success. Other values indicate error codes.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - `true`: The call was successful.
	//
	// - `false`: The call failed.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4CE1F661-75DD-4EBD-A4AD-057B26834ABB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBuildRecordByRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBuildRecordByRecordResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBuildRecordByRecordResponseBody) GetBuildRecordId() *string {
	return s.BuildRecordId
}

func (s *CreateBuildRecordByRecordResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateBuildRecordByRecordResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *CreateBuildRecordByRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBuildRecordByRecordResponseBody) SetBuildRecordId(v string) *CreateBuildRecordByRecordResponseBody {
	s.BuildRecordId = &v
	return s
}

func (s *CreateBuildRecordByRecordResponseBody) SetCode(v string) *CreateBuildRecordByRecordResponseBody {
	s.Code = &v
	return s
}

func (s *CreateBuildRecordByRecordResponseBody) SetIsSuccess(v bool) *CreateBuildRecordByRecordResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateBuildRecordByRecordResponseBody) SetRequestId(v string) *CreateBuildRecordByRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBuildRecordByRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
