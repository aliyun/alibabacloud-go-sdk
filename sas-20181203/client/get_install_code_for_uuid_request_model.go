// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstallCodeForUuidRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUuid(v string) *GetInstallCodeForUuidRequest
	GetUuid() *string
}

type GetInstallCodeForUuidRequest struct {
	// The UUID of the server whose agent installation code you want to query.
	//
	// > Call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// eae0b46e-2155-422e-9565-ecc52c69****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetInstallCodeForUuidRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstallCodeForUuidRequest) GoString() string {
	return s.String()
}

func (s *GetInstallCodeForUuidRequest) GetUuid() *string {
	return s.Uuid
}

func (s *GetInstallCodeForUuidRequest) SetUuid(v string) *GetInstallCodeForUuidRequest {
	s.Uuid = &v
	return s
}

func (s *GetInstallCodeForUuidRequest) Validate() error {
	return dara.Validate(s)
}
