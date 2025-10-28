// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *StopApplicationRequest
	GetAppId() *string
	SetEccInfo(v string) *StopApplicationRequest
	GetEccInfo() *string
}

type StopApplicationRequest struct {
	// The ID of the application. You can call the ListApplication operation to query the application ID. For more information, see [ListApplication](https://help.aliyun.com/document_detail/149390.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// c627c157-560d*******
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the elastic compute container (ECC) that corresponds to the Elastic Compute Service (ECS) instance on which you want to stop the application. You can call the QueryApplicationStatus operation to query the ECC ID. For more information, see [QueryApplicationStatus](https://help.aliyun.com/document_detail/149394.html).
	//
	// 	- If you want to stop the application on multiple ECS instances, separate the ECC IDs with commas (,).
	//
	// 	- If you leave this parameter empty, the application will be stopped on all ECS instances.
	//
	// example:
	//
	// 74ee9166-****1f6-bcb60e5b****
	EccInfo *string `json:"EccInfo,omitempty" xml:"EccInfo,omitempty"`
}

func (s StopApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s StopApplicationRequest) GoString() string {
	return s.String()
}

func (s *StopApplicationRequest) GetAppId() *string {
	return s.AppId
}

func (s *StopApplicationRequest) GetEccInfo() *string {
	return s.EccInfo
}

func (s *StopApplicationRequest) SetAppId(v string) *StopApplicationRequest {
	s.AppId = &v
	return s
}

func (s *StopApplicationRequest) SetEccInfo(v string) *StopApplicationRequest {
	s.EccInfo = &v
	return s
}

func (s *StopApplicationRequest) Validate() error {
	return dara.Validate(s)
}
