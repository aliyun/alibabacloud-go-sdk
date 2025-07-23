// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLicenseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetLicenseRequest
	GetId() *int64
	SetInstanceId(v string) *GetLicenseRequest
	GetInstanceId() *string
}

type GetLicenseRequest struct {
	// ID
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// c31238fcb74e482588a72de90cd7dba3
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetLicenseRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLicenseRequest) GoString() string {
	return s.String()
}

func (s *GetLicenseRequest) GetId() *int64 {
	return s.Id
}

func (s *GetLicenseRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetLicenseRequest) SetId(v int64) *GetLicenseRequest {
	s.Id = &v
	return s
}

func (s *GetLicenseRequest) SetInstanceId(v string) *GetLicenseRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLicenseRequest) Validate() error {
	return dara.Validate(s)
}
