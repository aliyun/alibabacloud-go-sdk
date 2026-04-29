// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisablePolarClawCronJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DisablePolarClawCronJobResponseBody
	GetApplicationId() *string
	SetCode(v int32) *DisablePolarClawCronJobResponseBody
	GetCode() *int32
	SetJobId(v string) *DisablePolarClawCronJobResponseBody
	GetJobId() *string
	SetMessage(v string) *DisablePolarClawCronJobResponseBody
	GetMessage() *string
	SetOk(v bool) *DisablePolarClawCronJobResponseBody
	GetOk() *bool
	SetRequestId(v string) *DisablePolarClawCronJobResponseBody
	GetRequestId() *string
}

type DisablePolarClawCronJobResponseBody struct {
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 0ee00f56-f467-4d41-858c-ca4ede2c770e
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// true
	Ok *bool `json:"Ok,omitempty" xml:"Ok,omitempty"`
	// example:
	//
	// 2281C6C9-CBAB-1AFD-8400-670750CF6025_2212
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisablePolarClawCronJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisablePolarClawCronJobResponseBody) GoString() string {
	return s.String()
}

func (s *DisablePolarClawCronJobResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DisablePolarClawCronJobResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DisablePolarClawCronJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *DisablePolarClawCronJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DisablePolarClawCronJobResponseBody) GetOk() *bool {
	return s.Ok
}

func (s *DisablePolarClawCronJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisablePolarClawCronJobResponseBody) SetApplicationId(v string) *DisablePolarClawCronJobResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *DisablePolarClawCronJobResponseBody) SetCode(v int32) *DisablePolarClawCronJobResponseBody {
	s.Code = &v
	return s
}

func (s *DisablePolarClawCronJobResponseBody) SetJobId(v string) *DisablePolarClawCronJobResponseBody {
	s.JobId = &v
	return s
}

func (s *DisablePolarClawCronJobResponseBody) SetMessage(v string) *DisablePolarClawCronJobResponseBody {
	s.Message = &v
	return s
}

func (s *DisablePolarClawCronJobResponseBody) SetOk(v bool) *DisablePolarClawCronJobResponseBody {
	s.Ok = &v
	return s
}

func (s *DisablePolarClawCronJobResponseBody) SetRequestId(v string) *DisablePolarClawCronJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisablePolarClawCronJobResponseBody) Validate() error {
	return dara.Validate(s)
}
