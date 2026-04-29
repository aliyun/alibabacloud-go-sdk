// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolarClawCronJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DeletePolarClawCronJobResponseBody
	GetApplicationId() *string
	SetCode(v int32) *DeletePolarClawCronJobResponseBody
	GetCode() *int32
	SetJobId(v string) *DeletePolarClawCronJobResponseBody
	GetJobId() *string
	SetMessage(v string) *DeletePolarClawCronJobResponseBody
	GetMessage() *string
	SetOk(v bool) *DeletePolarClawCronJobResponseBody
	GetOk() *bool
	SetRemoved(v bool) *DeletePolarClawCronJobResponseBody
	GetRemoved() *bool
	SetRequestId(v string) *DeletePolarClawCronJobResponseBody
	GetRequestId() *string
}

type DeletePolarClawCronJobResponseBody struct {
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
	// 8006e51c-dab3-4602-bc69-4f728002c6ce
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
	// true
	Removed *bool `json:"Removed,omitempty" xml:"Removed,omitempty"`
	// example:
	//
	// 2281C6C9-CBAB-1AFD-8400-670750CF6025_2212
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePolarClawCronJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePolarClawCronJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePolarClawCronJobResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DeletePolarClawCronJobResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeletePolarClawCronJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *DeletePolarClawCronJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeletePolarClawCronJobResponseBody) GetOk() *bool {
	return s.Ok
}

func (s *DeletePolarClawCronJobResponseBody) GetRemoved() *bool {
	return s.Removed
}

func (s *DeletePolarClawCronJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePolarClawCronJobResponseBody) SetApplicationId(v string) *DeletePolarClawCronJobResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *DeletePolarClawCronJobResponseBody) SetCode(v int32) *DeletePolarClawCronJobResponseBody {
	s.Code = &v
	return s
}

func (s *DeletePolarClawCronJobResponseBody) SetJobId(v string) *DeletePolarClawCronJobResponseBody {
	s.JobId = &v
	return s
}

func (s *DeletePolarClawCronJobResponseBody) SetMessage(v string) *DeletePolarClawCronJobResponseBody {
	s.Message = &v
	return s
}

func (s *DeletePolarClawCronJobResponseBody) SetOk(v bool) *DeletePolarClawCronJobResponseBody {
	s.Ok = &v
	return s
}

func (s *DeletePolarClawCronJobResponseBody) SetRemoved(v bool) *DeletePolarClawCronJobResponseBody {
	s.Removed = &v
	return s
}

func (s *DeletePolarClawCronJobResponseBody) SetRequestId(v string) *DeletePolarClawCronJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePolarClawCronJobResponseBody) Validate() error {
	return dara.Validate(s)
}
