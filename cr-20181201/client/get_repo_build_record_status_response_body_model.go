// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRepoBuildRecordStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBuildStatus(v string) *GetRepoBuildRecordStatusResponseBody
	GetBuildStatus() *string
	SetCode(v string) *GetRepoBuildRecordStatusResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *GetRepoBuildRecordStatusResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *GetRepoBuildRecordStatusResponseBody
	GetRequestId() *string
}

type GetRepoBuildRecordStatusResponseBody struct {
	// The status of the image building.
	//
	// example:
	//
	// success
	BuildStatus *string `json:"BuildStatus,omitempty" xml:"BuildStatus,omitempty"`
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
	// 79174CBA-8556-443A-8976-22C922D7BE37
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRepoBuildRecordStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRepoBuildRecordStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetRepoBuildRecordStatusResponseBody) GetBuildStatus() *string {
	return s.BuildStatus
}

func (s *GetRepoBuildRecordStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetRepoBuildRecordStatusResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetRepoBuildRecordStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRepoBuildRecordStatusResponseBody) SetBuildStatus(v string) *GetRepoBuildRecordStatusResponseBody {
	s.BuildStatus = &v
	return s
}

func (s *GetRepoBuildRecordStatusResponseBody) SetCode(v string) *GetRepoBuildRecordStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetRepoBuildRecordStatusResponseBody) SetIsSuccess(v bool) *GetRepoBuildRecordStatusResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetRepoBuildRecordStatusResponseBody) SetRequestId(v string) *GetRepoBuildRecordStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRepoBuildRecordStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
