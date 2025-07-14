// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *SuspendJobRequest
	GetAppId() *string
	SetSuspend(v bool) *SuspendJobRequest
	GetSuspend() *bool
}

type SuspendJobRequest struct {
	// The ID of the job template.
	//
	// This parameter is required.
	//
	// example:
	//
	// ee1a7a07-abcb-4652-a1d3-2d57f415****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Start or suspend a job template.
	//
	// 	- true: Start a job template.
	//
	// 	- false: Suspend a job template.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Suspend *bool `json:"Suspend,omitempty" xml:"Suspend,omitempty"`
}

func (s SuspendJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SuspendJobRequest) GoString() string {
	return s.String()
}

func (s *SuspendJobRequest) GetAppId() *string {
	return s.AppId
}

func (s *SuspendJobRequest) GetSuspend() *bool {
	return s.Suspend
}

func (s *SuspendJobRequest) SetAppId(v string) *SuspendJobRequest {
	s.AppId = &v
	return s
}

func (s *SuspendJobRequest) SetSuspend(v bool) *SuspendJobRequest {
	s.Suspend = &v
	return s
}

func (s *SuspendJobRequest) Validate() error {
	return dara.Validate(s)
}
