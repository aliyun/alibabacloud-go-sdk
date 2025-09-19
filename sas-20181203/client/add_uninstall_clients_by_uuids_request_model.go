// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUninstallClientsByUuidsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallMethod(v string) *AddUninstallClientsByUuidsRequest
	GetCallMethod() *string
	SetFeedback(v string) *AddUninstallClientsByUuidsRequest
	GetFeedback() *string
	SetRegion(v string) *AddUninstallClientsByUuidsRequest
	GetRegion() *string
	SetSourceIp(v string) *AddUninstallClientsByUuidsRequest
	GetSourceIp() *string
	SetUuids(v string) *AddUninstallClientsByUuidsRequest
	GetUuids() *string
}

type AddUninstallClientsByUuidsRequest struct {
	// The method name. Default value: init.
	//
	// example:
	//
	// init
	CallMethod *string `json:"CallMethod,omitempty" xml:"CallMethod,omitempty"`
	// The feedback.
	//
	// example:
	//
	// reinstall
	Feedback *string `json:"Feedback,omitempty" xml:"Feedback,omitempty"`
	// The region in which the server resides.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The source IP address of the request. You do not need to specify this parameter. It is automatically obtained by the system.
	//
	// example:
	//
	// 1.2.XXX.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The UUID of the server that you want to unbind. Separate multiple UUIDs with commas (,).
	//
	// >  You can call the [DescribeCloudCenterInstances](https://help.aliyun.com/document_detail/141932.html) operation to query the UUIDs of servers.
	//
	// This parameter is required.
	//
	// example:
	//
	// inet-183707ae-3bdf-4db0-b771-3e9962bf****,inet-49dceccc-4f01-469b-8411-2416ea12****
	Uuids *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s AddUninstallClientsByUuidsRequest) String() string {
	return dara.Prettify(s)
}

func (s AddUninstallClientsByUuidsRequest) GoString() string {
	return s.String()
}

func (s *AddUninstallClientsByUuidsRequest) GetCallMethod() *string {
	return s.CallMethod
}

func (s *AddUninstallClientsByUuidsRequest) GetFeedback() *string {
	return s.Feedback
}

func (s *AddUninstallClientsByUuidsRequest) GetRegion() *string {
	return s.Region
}

func (s *AddUninstallClientsByUuidsRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *AddUninstallClientsByUuidsRequest) GetUuids() *string {
	return s.Uuids
}

func (s *AddUninstallClientsByUuidsRequest) SetCallMethod(v string) *AddUninstallClientsByUuidsRequest {
	s.CallMethod = &v
	return s
}

func (s *AddUninstallClientsByUuidsRequest) SetFeedback(v string) *AddUninstallClientsByUuidsRequest {
	s.Feedback = &v
	return s
}

func (s *AddUninstallClientsByUuidsRequest) SetRegion(v string) *AddUninstallClientsByUuidsRequest {
	s.Region = &v
	return s
}

func (s *AddUninstallClientsByUuidsRequest) SetSourceIp(v string) *AddUninstallClientsByUuidsRequest {
	s.SourceIp = &v
	return s
}

func (s *AddUninstallClientsByUuidsRequest) SetUuids(v string) *AddUninstallClientsByUuidsRequest {
	s.Uuids = &v
	return s
}

func (s *AddUninstallClientsByUuidsRequest) Validate() error {
	return dara.Validate(s)
}
