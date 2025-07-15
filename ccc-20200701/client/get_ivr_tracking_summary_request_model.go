// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIvrTrackingSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v string) *GetIvrTrackingSummaryRequest
	GetContactId() *string
	SetInstanceId(v string) *GetIvrTrackingSummaryRequest
	GetInstanceId() *string
}

type GetIvrTrackingSummaryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// job-489361145506897920
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0f7ad007-ab50-4b3d-a87a-56864eb40dab
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetIvrTrackingSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIvrTrackingSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetIvrTrackingSummaryRequest) GetContactId() *string {
	return s.ContactId
}

func (s *GetIvrTrackingSummaryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetIvrTrackingSummaryRequest) SetContactId(v string) *GetIvrTrackingSummaryRequest {
	s.ContactId = &v
	return s
}

func (s *GetIvrTrackingSummaryRequest) SetInstanceId(v string) *GetIvrTrackingSummaryRequest {
	s.InstanceId = &v
	return s
}

func (s *GetIvrTrackingSummaryRequest) Validate() error {
	return dara.Validate(s)
}
