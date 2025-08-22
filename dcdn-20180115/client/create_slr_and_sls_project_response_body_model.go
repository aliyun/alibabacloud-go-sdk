// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSlrAndSlsProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSlrAndSlsProjectResponseBody
	GetRequestId() *string
	SetSlsInfo(v *CreateSlrAndSlsProjectResponseBodySlsInfo) *CreateSlrAndSlsProjectResponseBody
	GetSlsInfo() *CreateSlrAndSlsProjectResponseBodySlsInfo
}

type CreateSlrAndSlsProjectResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// fe33a379-5053-4f22-a73c-826e2b44355d
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about Log Service.
	SlsInfo *CreateSlrAndSlsProjectResponseBodySlsInfo `json:"SlsInfo,omitempty" xml:"SlsInfo,omitempty" type:"Struct"`
}

func (s CreateSlrAndSlsProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSlrAndSlsProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSlrAndSlsProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSlrAndSlsProjectResponseBody) GetSlsInfo() *CreateSlrAndSlsProjectResponseBodySlsInfo {
	return s.SlsInfo
}

func (s *CreateSlrAndSlsProjectResponseBody) SetRequestId(v string) *CreateSlrAndSlsProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSlrAndSlsProjectResponseBody) SetSlsInfo(v *CreateSlrAndSlsProjectResponseBodySlsInfo) *CreateSlrAndSlsProjectResponseBody {
	s.SlsInfo = v
	return s
}

func (s *CreateSlrAndSlsProjectResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateSlrAndSlsProjectResponseBodySlsInfo struct {
	// The endpoint of Log Service.
	//
	// example:
	//
	// cn-shanghai.log.*.com
	EndPoint *string `json:"EndPoint,omitempty" xml:"EndPoint,omitempty"`
	// The Logstore of Log Service.
	//
	// example:
	//
	// dcdn-edge-trlog
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	// The project of Log Service.
	//
	// example:
	//
	// dcdn-edge-rtlog-cn-cfc7****
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The region where Log Service resides.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s CreateSlrAndSlsProjectResponseBodySlsInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateSlrAndSlsProjectResponseBodySlsInfo) GoString() string {
	return s.String()
}

func (s *CreateSlrAndSlsProjectResponseBodySlsInfo) GetEndPoint() *string {
	return s.EndPoint
}

func (s *CreateSlrAndSlsProjectResponseBodySlsInfo) GetLogStore() *string {
	return s.LogStore
}

func (s *CreateSlrAndSlsProjectResponseBodySlsInfo) GetProject() *string {
	return s.Project
}

func (s *CreateSlrAndSlsProjectResponseBodySlsInfo) GetRegion() *string {
	return s.Region
}

func (s *CreateSlrAndSlsProjectResponseBodySlsInfo) SetEndPoint(v string) *CreateSlrAndSlsProjectResponseBodySlsInfo {
	s.EndPoint = &v
	return s
}

func (s *CreateSlrAndSlsProjectResponseBodySlsInfo) SetLogStore(v string) *CreateSlrAndSlsProjectResponseBodySlsInfo {
	s.LogStore = &v
	return s
}

func (s *CreateSlrAndSlsProjectResponseBodySlsInfo) SetProject(v string) *CreateSlrAndSlsProjectResponseBodySlsInfo {
	s.Project = &v
	return s
}

func (s *CreateSlrAndSlsProjectResponseBodySlsInfo) SetRegion(v string) *CreateSlrAndSlsProjectResponseBodySlsInfo {
	s.Region = &v
	return s
}

func (s *CreateSlrAndSlsProjectResponseBodySlsInfo) Validate() error {
	return dara.Validate(s)
}
