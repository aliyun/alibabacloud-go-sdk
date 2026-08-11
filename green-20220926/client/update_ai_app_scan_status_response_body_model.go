// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAiAppScanStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedAppIds(v []*string) *UpdateAiAppScanStatusResponseBody
	GetFailedAppIds() []*string
	SetRequestId(v string) *UpdateAiAppScanStatusResponseBody
	GetRequestId() *string
	SetStatus(v string) *UpdateAiAppScanStatusResponseBody
	GetStatus() *string
	SetSuccessAppIds(v []*string) *UpdateAiAppScanStatusResponseBody
	GetSuccessAppIds() []*string
}

type UpdateAiAppScanStatusResponseBody struct {
	// The list of application IDs that failed.
	FailedAppIds []*string `json:"FailedAppIds,omitempty" xml:"FailedAppIds,omitempty" type:"Repeated"`
	// The ID assigned by the backend to uniquely identify a request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status. Valid values:
	//
	// - SUCCESS: Succeeded.
	//
	// - PARTIAL_SUCCESS: Partially succeeded.
	//
	// - FAILED: Failed.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of application IDs that succeeded.
	SuccessAppIds []*string `json:"SuccessAppIds,omitempty" xml:"SuccessAppIds,omitempty" type:"Repeated"`
}

func (s UpdateAiAppScanStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAiAppScanStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAiAppScanStatusResponseBody) GetFailedAppIds() []*string {
	return s.FailedAppIds
}

func (s *UpdateAiAppScanStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAiAppScanStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UpdateAiAppScanStatusResponseBody) GetSuccessAppIds() []*string {
	return s.SuccessAppIds
}

func (s *UpdateAiAppScanStatusResponseBody) SetFailedAppIds(v []*string) *UpdateAiAppScanStatusResponseBody {
	s.FailedAppIds = v
	return s
}

func (s *UpdateAiAppScanStatusResponseBody) SetRequestId(v string) *UpdateAiAppScanStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAiAppScanStatusResponseBody) SetStatus(v string) *UpdateAiAppScanStatusResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateAiAppScanStatusResponseBody) SetSuccessAppIds(v []*string) *UpdateAiAppScanStatusResponseBody {
	s.SuccessAppIds = v
	return s
}

func (s *UpdateAiAppScanStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
