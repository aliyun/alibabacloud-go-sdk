// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPatchPolarClawConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *PatchPolarClawConfigResponseBody
	GetApplicationId() *string
	SetCode(v int32) *PatchPolarClawConfigResponseBody
	GetCode() *int32
	SetMessage(v string) *PatchPolarClawConfigResponseBody
	GetMessage() *string
	SetOk(v bool) *PatchPolarClawConfigResponseBody
	GetOk() *bool
	SetPatchedKeys(v []*string) *PatchPolarClawConfigResponseBody
	GetPatchedKeys() []*string
	SetRequestId(v string) *PatchPolarClawConfigResponseBody
	GetRequestId() *string
	SetRestarted(v bool) *PatchPolarClawConfigResponseBody
	GetRestarted() *bool
}

type PatchPolarClawConfigResponseBody struct {
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
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// true
	Ok          *bool     `json:"Ok,omitempty" xml:"Ok,omitempty"`
	PatchedKeys []*string `json:"PatchedKeys,omitempty" xml:"PatchedKeys,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 3E5CD764-FCCA-5C9C-838E-20E0DE84B2AF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Restarted *bool `json:"Restarted,omitempty" xml:"Restarted,omitempty"`
}

func (s PatchPolarClawConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PatchPolarClawConfigResponseBody) GoString() string {
	return s.String()
}

func (s *PatchPolarClawConfigResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *PatchPolarClawConfigResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *PatchPolarClawConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PatchPolarClawConfigResponseBody) GetOk() *bool {
	return s.Ok
}

func (s *PatchPolarClawConfigResponseBody) GetPatchedKeys() []*string {
	return s.PatchedKeys
}

func (s *PatchPolarClawConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PatchPolarClawConfigResponseBody) GetRestarted() *bool {
	return s.Restarted
}

func (s *PatchPolarClawConfigResponseBody) SetApplicationId(v string) *PatchPolarClawConfigResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *PatchPolarClawConfigResponseBody) SetCode(v int32) *PatchPolarClawConfigResponseBody {
	s.Code = &v
	return s
}

func (s *PatchPolarClawConfigResponseBody) SetMessage(v string) *PatchPolarClawConfigResponseBody {
	s.Message = &v
	return s
}

func (s *PatchPolarClawConfigResponseBody) SetOk(v bool) *PatchPolarClawConfigResponseBody {
	s.Ok = &v
	return s
}

func (s *PatchPolarClawConfigResponseBody) SetPatchedKeys(v []*string) *PatchPolarClawConfigResponseBody {
	s.PatchedKeys = v
	return s
}

func (s *PatchPolarClawConfigResponseBody) SetRequestId(v string) *PatchPolarClawConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *PatchPolarClawConfigResponseBody) SetRestarted(v bool) *PatchPolarClawConfigResponseBody {
	s.Restarted = &v
	return s
}

func (s *PatchPolarClawConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
