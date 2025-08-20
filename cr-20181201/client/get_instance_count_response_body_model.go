// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetInstanceCountResponseBody
	GetCode() *string
	SetCount(v int32) *GetInstanceCountResponseBody
	GetCount() *int32
	SetIsSuccess(v bool) *GetInstanceCountResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *GetInstanceCountResponseBody
	GetRequestId() *string
}

type GetInstanceCountResponseBody struct {
	// Return value
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Number of instances
	//
	// example:
	//
	// 5
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// Indicates whether the API call was successful. Values:
	//
	// - `true`: The API call was successful.
	//
	// - `false`: The API call failed.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// Request ID
	//
	// example:
	//
	// BC648259-91A7-4502-BED3-EDF64361FA83
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInstanceCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceCountResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceCountResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetInstanceCountResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *GetInstanceCountResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetInstanceCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceCountResponseBody) SetCode(v string) *GetInstanceCountResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceCountResponseBody) SetCount(v int32) *GetInstanceCountResponseBody {
	s.Count = &v
	return s
}

func (s *GetInstanceCountResponseBody) SetIsSuccess(v bool) *GetInstanceCountResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetInstanceCountResponseBody) SetRequestId(v string) *GetInstanceCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceCountResponseBody) Validate() error {
	return dara.Validate(s)
}
