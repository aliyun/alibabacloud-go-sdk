// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSlrAndSlsProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v string) *CreateSlrAndSlsProjectRequest
	GetBusinessType() *string
	SetRegion(v string) *CreateSlrAndSlsProjectRequest
	GetRegion() *string
}

type CreateSlrAndSlsProjectRequest struct {
	// The type of the collected logs. Default value: cdn_log_access_l1. Valid values:
	//
	// 	- **cdn_log_access_l1**: access logs of L1 Dynamic Route for CDN (DCDN) points of presence (POPs)
	//
	// 	- **cdn_log_origin**: back-to-origin logs
	//
	// 	- **cdn_log_er**: EdgeRoutine logs
	//
	// example:
	//
	// cdn_log_access_l1
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The region where Log Service resides. Valid values:
	//
	// 	- **cn-hangzhou**
	//
	// 	- **cn-shanghai**
	//
	// 	- **cn-qingdao**
	//
	// 	- **cn-beijing**
	//
	// 	- **cn-zhangjiakou**
	//
	// 	- **cn-shenzhen**
	//
	// 	- **eu-central-1**
	//
	// 	- **us-west-1**
	//
	// 	- **ap-south-1**
	//
	// 	- **ap-southeast-1**
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s CreateSlrAndSlsProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSlrAndSlsProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateSlrAndSlsProjectRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *CreateSlrAndSlsProjectRequest) GetRegion() *string {
	return s.Region
}

func (s *CreateSlrAndSlsProjectRequest) SetBusinessType(v string) *CreateSlrAndSlsProjectRequest {
	s.BusinessType = &v
	return s
}

func (s *CreateSlrAndSlsProjectRequest) SetRegion(v string) *CreateSlrAndSlsProjectRequest {
	s.Region = &v
	return s
}

func (s *CreateSlrAndSlsProjectRequest) Validate() error {
	return dara.Validate(s)
}
