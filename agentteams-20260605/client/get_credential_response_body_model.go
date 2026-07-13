// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCredentialResponseBody
	GetCode() *string
	SetData(v *GetCredentialResponseBodyData) *GetCredentialResponseBody
	GetData() *GetCredentialResponseBodyData
	SetHttpStatusCode(v int32) *GetCredentialResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetCredentialResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCredentialResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCredentialResponseBody
	GetSuccess() *bool
}

type GetCredentialResponseBody struct {
	Code           *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetCredentialResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                         `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *GetCredentialResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCredentialResponseBody) GetData() *GetCredentialResponseBodyData {
	return s.Data
}

func (s *GetCredentialResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetCredentialResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCredentialResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCredentialResponseBody) SetCode(v string) *GetCredentialResponseBody {
	s.Code = &v
	return s
}

func (s *GetCredentialResponseBody) SetData(v *GetCredentialResponseBodyData) *GetCredentialResponseBody {
	s.Data = v
	return s
}

func (s *GetCredentialResponseBody) SetHttpStatusCode(v int32) *GetCredentialResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetCredentialResponseBody) SetMessage(v string) *GetCredentialResponseBody {
	s.Message = &v
	return s
}

func (s *GetCredentialResponseBody) SetRequestId(v string) *GetCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCredentialResponseBody) SetSuccess(v bool) *GetCredentialResponseBody {
	s.Success = &v
	return s
}

func (s *GetCredentialResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCredentialResponseBodyData struct {
	BoundWorkers []*GetCredentialResponseBodyDataBoundWorkers `json:"BoundWorkers,omitempty" xml:"BoundWorkers,omitempty" type:"Repeated"`
	CreateTime   *string                                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description  *string                                      `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceId   *string                                      `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name         *string                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId     *string                                      `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status       *string                                      `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime   *string                                      `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetCredentialResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCredentialResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCredentialResponseBodyData) GetBoundWorkers() []*GetCredentialResponseBodyDataBoundWorkers {
	return s.BoundWorkers
}

func (s *GetCredentialResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetCredentialResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetCredentialResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetCredentialResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetCredentialResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetCredentialResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetCredentialResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetCredentialResponseBodyData) SetBoundWorkers(v []*GetCredentialResponseBodyDataBoundWorkers) *GetCredentialResponseBodyData {
	s.BoundWorkers = v
	return s
}

func (s *GetCredentialResponseBodyData) SetCreateTime(v string) *GetCredentialResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetCredentialResponseBodyData) SetDescription(v string) *GetCredentialResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetCredentialResponseBodyData) SetInstanceId(v string) *GetCredentialResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetCredentialResponseBodyData) SetName(v string) *GetCredentialResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetCredentialResponseBodyData) SetRegionId(v string) *GetCredentialResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetCredentialResponseBodyData) SetStatus(v string) *GetCredentialResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetCredentialResponseBodyData) SetUpdateTime(v string) *GetCredentialResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetCredentialResponseBodyData) Validate() error {
	if s.BoundWorkers != nil {
		for _, item := range s.BoundWorkers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCredentialResponseBodyDataBoundWorkers struct {
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetCredentialResponseBodyDataBoundWorkers) String() string {
	return dara.Prettify(s)
}

func (s GetCredentialResponseBodyDataBoundWorkers) GoString() string {
	return s.String()
}

func (s *GetCredentialResponseBodyDataBoundWorkers) GetName() *string {
	return s.Name
}

func (s *GetCredentialResponseBodyDataBoundWorkers) GetStatus() *string {
	return s.Status
}

func (s *GetCredentialResponseBodyDataBoundWorkers) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetCredentialResponseBodyDataBoundWorkers) SetName(v string) *GetCredentialResponseBodyDataBoundWorkers {
	s.Name = &v
	return s
}

func (s *GetCredentialResponseBodyDataBoundWorkers) SetStatus(v string) *GetCredentialResponseBodyDataBoundWorkers {
	s.Status = &v
	return s
}

func (s *GetCredentialResponseBodyDataBoundWorkers) SetUpdateTime(v string) *GetCredentialResponseBodyDataBoundWorkers {
	s.UpdateTime = &v
	return s
}

func (s *GetCredentialResponseBodyDataBoundWorkers) Validate() error {
	return dara.Validate(s)
}
