// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRepoBuildRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBuildRecordId(v string) *GetRepoBuildRecordRequest
	GetBuildRecordId() *string
	SetInstanceId(v string) *GetRepoBuildRecordRequest
	GetInstanceId() *string
}

type GetRepoBuildRecordRequest struct {
	// The ID of the image building record.
	//
	// This parameter is required.
	//
	// example:
	//
	// a78ec6fb-16ea-4649-93b7-f52afba7d****
	BuildRecordId *string `json:"BuildRecordId,omitempty" xml:"BuildRecordId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetRepoBuildRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRepoBuildRecordRequest) GoString() string {
	return s.String()
}

func (s *GetRepoBuildRecordRequest) GetBuildRecordId() *string {
	return s.BuildRecordId
}

func (s *GetRepoBuildRecordRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRepoBuildRecordRequest) SetBuildRecordId(v string) *GetRepoBuildRecordRequest {
	s.BuildRecordId = &v
	return s
}

func (s *GetRepoBuildRecordRequest) SetInstanceId(v string) *GetRepoBuildRecordRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRepoBuildRecordRequest) Validate() error {
	return dara.Validate(s)
}
