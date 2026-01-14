// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectApplicationMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DetectApplicationMonitorRequest
	GetClientToken() *string
	SetRegionId(v string) *DetectApplicationMonitorRequest
	GetRegionId() *string
	SetTaskId(v string) *DetectApplicationMonitorRequest
	GetTaskId() *string
}

type DetectApplicationMonitorRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the region where the Global Accelerator (GA) instance is deployed. Set the value to **cn-hangzhou**.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the origin probing task for which you want to enable the diagnostic feature.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2c8dbdf9-a3ab-46a1-85a4-f094965e****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DetectApplicationMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectApplicationMonitorRequest) GoString() string {
	return s.String()
}

func (s *DetectApplicationMonitorRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DetectApplicationMonitorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetectApplicationMonitorRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DetectApplicationMonitorRequest) SetClientToken(v string) *DetectApplicationMonitorRequest {
	s.ClientToken = &v
	return s
}

func (s *DetectApplicationMonitorRequest) SetRegionId(v string) *DetectApplicationMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *DetectApplicationMonitorRequest) SetTaskId(v string) *DetectApplicationMonitorRequest {
	s.TaskId = &v
	return s
}

func (s *DetectApplicationMonitorRequest) Validate() error {
	return dara.Validate(s)
}
