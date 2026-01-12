// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListWorkflowResponseBodyData) *ListWorkflowResponseBody
	GetData() *ListWorkflowResponseBodyData
	SetHttpCode(v int64) *ListWorkflowResponseBody
	GetHttpCode() *int64
	SetMessage(v string) *ListWorkflowResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListWorkflowResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListWorkflowResponseBody
	GetSuccess() *bool
}

type ListWorkflowResponseBody struct {
	Data      *ListWorkflowResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpCode  *int64                        `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListWorkflowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkflowResponseBody) GetData() *ListWorkflowResponseBodyData {
	return s.Data
}

func (s *ListWorkflowResponseBody) GetHttpCode() *int64 {
	return s.HttpCode
}

func (s *ListWorkflowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListWorkflowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkflowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListWorkflowResponseBody) SetData(v *ListWorkflowResponseBodyData) *ListWorkflowResponseBody {
	s.Data = v
	return s
}

func (s *ListWorkflowResponseBody) SetHttpCode(v int64) *ListWorkflowResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListWorkflowResponseBody) SetMessage(v string) *ListWorkflowResponseBody {
	s.Message = &v
	return s
}

func (s *ListWorkflowResponseBody) SetRequestId(v string) *ListWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkflowResponseBody) SetSuccess(v bool) *ListWorkflowResponseBody {
	s.Success = &v
	return s
}

func (s *ListWorkflowResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListWorkflowResponseBodyData struct {
	HumanPose        []*ListWorkflowResponseBodyDataHumanPose        `json:"HumanPose,omitempty" xml:"HumanPose,omitempty" type:"Repeated"`
	Mannequins       []*ListWorkflowResponseBodyDataMannequins       `json:"Mannequins,omitempty" xml:"Mannequins,omitempty" type:"Repeated"`
	Object           []*ListWorkflowResponseBodyDataObject           `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
	ObjectGeneration []*ListWorkflowResponseBodyDataObjectGeneration `json:"ObjectGeneration,omitempty" xml:"ObjectGeneration,omitempty" type:"Repeated"`
	Scene            []*ListWorkflowResponseBodyDataScene            `json:"Scene,omitempty" xml:"Scene,omitempty" type:"Repeated"`
}

func (s ListWorkflowResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListWorkflowResponseBodyData) GetHumanPose() []*ListWorkflowResponseBodyDataHumanPose {
	return s.HumanPose
}

func (s *ListWorkflowResponseBodyData) GetMannequins() []*ListWorkflowResponseBodyDataMannequins {
	return s.Mannequins
}

func (s *ListWorkflowResponseBodyData) GetObject() []*ListWorkflowResponseBodyDataObject {
	return s.Object
}

func (s *ListWorkflowResponseBodyData) GetObjectGeneration() []*ListWorkflowResponseBodyDataObjectGeneration {
	return s.ObjectGeneration
}

func (s *ListWorkflowResponseBodyData) GetScene() []*ListWorkflowResponseBodyDataScene {
	return s.Scene
}

func (s *ListWorkflowResponseBodyData) SetHumanPose(v []*ListWorkflowResponseBodyDataHumanPose) *ListWorkflowResponseBodyData {
	s.HumanPose = v
	return s
}

func (s *ListWorkflowResponseBodyData) SetMannequins(v []*ListWorkflowResponseBodyDataMannequins) *ListWorkflowResponseBodyData {
	s.Mannequins = v
	return s
}

func (s *ListWorkflowResponseBodyData) SetObject(v []*ListWorkflowResponseBodyDataObject) *ListWorkflowResponseBodyData {
	s.Object = v
	return s
}

func (s *ListWorkflowResponseBodyData) SetObjectGeneration(v []*ListWorkflowResponseBodyDataObjectGeneration) *ListWorkflowResponseBodyData {
	s.ObjectGeneration = v
	return s
}

func (s *ListWorkflowResponseBodyData) SetScene(v []*ListWorkflowResponseBodyDataScene) *ListWorkflowResponseBodyData {
	s.Scene = v
	return s
}

func (s *ListWorkflowResponseBodyData) Validate() error {
	if s.HumanPose != nil {
		for _, item := range s.HumanPose {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Mannequins != nil {
		for _, item := range s.Mannequins {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Object != nil {
		for _, item := range s.Object {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ObjectGeneration != nil {
		for _, item := range s.ObjectGeneration {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Scene != nil {
		for _, item := range s.Scene {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWorkflowResponseBodyDataHumanPose struct {
	BizUsage   *string `json:"BizUsage,omitempty" xml:"BizUsage,omitempty"`
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	Status     *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListWorkflowResponseBodyDataHumanPose) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowResponseBodyDataHumanPose) GoString() string {
	return s.String()
}

func (s *ListWorkflowResponseBodyDataHumanPose) GetBizUsage() *string {
	return s.BizUsage
}

func (s *ListWorkflowResponseBodyDataHumanPose) GetId() *int64 {
	return s.Id
}

func (s *ListWorkflowResponseBodyDataHumanPose) GetName() *string {
	return s.Name
}

func (s *ListWorkflowResponseBodyDataHumanPose) GetObjectType() *string {
	return s.ObjectType
}

func (s *ListWorkflowResponseBodyDataHumanPose) GetStatus() *int64 {
	return s.Status
}

func (s *ListWorkflowResponseBodyDataHumanPose) SetBizUsage(v string) *ListWorkflowResponseBodyDataHumanPose {
	s.BizUsage = &v
	return s
}

func (s *ListWorkflowResponseBodyDataHumanPose) SetId(v int64) *ListWorkflowResponseBodyDataHumanPose {
	s.Id = &v
	return s
}

func (s *ListWorkflowResponseBodyDataHumanPose) SetName(v string) *ListWorkflowResponseBodyDataHumanPose {
	s.Name = &v
	return s
}

func (s *ListWorkflowResponseBodyDataHumanPose) SetObjectType(v string) *ListWorkflowResponseBodyDataHumanPose {
	s.ObjectType = &v
	return s
}

func (s *ListWorkflowResponseBodyDataHumanPose) SetStatus(v int64) *ListWorkflowResponseBodyDataHumanPose {
	s.Status = &v
	return s
}

func (s *ListWorkflowResponseBodyDataHumanPose) Validate() error {
	return dara.Validate(s)
}

type ListWorkflowResponseBodyDataMannequins struct {
	BizUsage   *string `json:"BizUsage,omitempty" xml:"BizUsage,omitempty"`
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	Status     *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListWorkflowResponseBodyDataMannequins) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowResponseBodyDataMannequins) GoString() string {
	return s.String()
}

func (s *ListWorkflowResponseBodyDataMannequins) GetBizUsage() *string {
	return s.BizUsage
}

func (s *ListWorkflowResponseBodyDataMannequins) GetId() *int64 {
	return s.Id
}

func (s *ListWorkflowResponseBodyDataMannequins) GetName() *string {
	return s.Name
}

func (s *ListWorkflowResponseBodyDataMannequins) GetObjectType() *string {
	return s.ObjectType
}

func (s *ListWorkflowResponseBodyDataMannequins) GetStatus() *int64 {
	return s.Status
}

func (s *ListWorkflowResponseBodyDataMannequins) SetBizUsage(v string) *ListWorkflowResponseBodyDataMannequins {
	s.BizUsage = &v
	return s
}

func (s *ListWorkflowResponseBodyDataMannequins) SetId(v int64) *ListWorkflowResponseBodyDataMannequins {
	s.Id = &v
	return s
}

func (s *ListWorkflowResponseBodyDataMannequins) SetName(v string) *ListWorkflowResponseBodyDataMannequins {
	s.Name = &v
	return s
}

func (s *ListWorkflowResponseBodyDataMannequins) SetObjectType(v string) *ListWorkflowResponseBodyDataMannequins {
	s.ObjectType = &v
	return s
}

func (s *ListWorkflowResponseBodyDataMannequins) SetStatus(v int64) *ListWorkflowResponseBodyDataMannequins {
	s.Status = &v
	return s
}

func (s *ListWorkflowResponseBodyDataMannequins) Validate() error {
	return dara.Validate(s)
}

type ListWorkflowResponseBodyDataObject struct {
	BizUsage   *string `json:"BizUsage,omitempty" xml:"BizUsage,omitempty"`
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	Status     *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListWorkflowResponseBodyDataObject) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowResponseBodyDataObject) GoString() string {
	return s.String()
}

func (s *ListWorkflowResponseBodyDataObject) GetBizUsage() *string {
	return s.BizUsage
}

func (s *ListWorkflowResponseBodyDataObject) GetId() *int64 {
	return s.Id
}

func (s *ListWorkflowResponseBodyDataObject) GetName() *string {
	return s.Name
}

func (s *ListWorkflowResponseBodyDataObject) GetObjectType() *string {
	return s.ObjectType
}

func (s *ListWorkflowResponseBodyDataObject) GetStatus() *int64 {
	return s.Status
}

func (s *ListWorkflowResponseBodyDataObject) SetBizUsage(v string) *ListWorkflowResponseBodyDataObject {
	s.BizUsage = &v
	return s
}

func (s *ListWorkflowResponseBodyDataObject) SetId(v int64) *ListWorkflowResponseBodyDataObject {
	s.Id = &v
	return s
}

func (s *ListWorkflowResponseBodyDataObject) SetName(v string) *ListWorkflowResponseBodyDataObject {
	s.Name = &v
	return s
}

func (s *ListWorkflowResponseBodyDataObject) SetObjectType(v string) *ListWorkflowResponseBodyDataObject {
	s.ObjectType = &v
	return s
}

func (s *ListWorkflowResponseBodyDataObject) SetStatus(v int64) *ListWorkflowResponseBodyDataObject {
	s.Status = &v
	return s
}

func (s *ListWorkflowResponseBodyDataObject) Validate() error {
	return dara.Validate(s)
}

type ListWorkflowResponseBodyDataObjectGeneration struct {
	BizUsage   *string `json:"BizUsage,omitempty" xml:"BizUsage,omitempty"`
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	Status     *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListWorkflowResponseBodyDataObjectGeneration) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowResponseBodyDataObjectGeneration) GoString() string {
	return s.String()
}

func (s *ListWorkflowResponseBodyDataObjectGeneration) GetBizUsage() *string {
	return s.BizUsage
}

func (s *ListWorkflowResponseBodyDataObjectGeneration) GetId() *int64 {
	return s.Id
}

func (s *ListWorkflowResponseBodyDataObjectGeneration) GetName() *string {
	return s.Name
}

func (s *ListWorkflowResponseBodyDataObjectGeneration) GetObjectType() *string {
	return s.ObjectType
}

func (s *ListWorkflowResponseBodyDataObjectGeneration) GetStatus() *int64 {
	return s.Status
}

func (s *ListWorkflowResponseBodyDataObjectGeneration) SetBizUsage(v string) *ListWorkflowResponseBodyDataObjectGeneration {
	s.BizUsage = &v
	return s
}

func (s *ListWorkflowResponseBodyDataObjectGeneration) SetId(v int64) *ListWorkflowResponseBodyDataObjectGeneration {
	s.Id = &v
	return s
}

func (s *ListWorkflowResponseBodyDataObjectGeneration) SetName(v string) *ListWorkflowResponseBodyDataObjectGeneration {
	s.Name = &v
	return s
}

func (s *ListWorkflowResponseBodyDataObjectGeneration) SetObjectType(v string) *ListWorkflowResponseBodyDataObjectGeneration {
	s.ObjectType = &v
	return s
}

func (s *ListWorkflowResponseBodyDataObjectGeneration) SetStatus(v int64) *ListWorkflowResponseBodyDataObjectGeneration {
	s.Status = &v
	return s
}

func (s *ListWorkflowResponseBodyDataObjectGeneration) Validate() error {
	return dara.Validate(s)
}

type ListWorkflowResponseBodyDataScene struct {
	BizUsage   *string `json:"BizUsage,omitempty" xml:"BizUsage,omitempty"`
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	Status     *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListWorkflowResponseBodyDataScene) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowResponseBodyDataScene) GoString() string {
	return s.String()
}

func (s *ListWorkflowResponseBodyDataScene) GetBizUsage() *string {
	return s.BizUsage
}

func (s *ListWorkflowResponseBodyDataScene) GetId() *int64 {
	return s.Id
}

func (s *ListWorkflowResponseBodyDataScene) GetName() *string {
	return s.Name
}

func (s *ListWorkflowResponseBodyDataScene) GetObjectType() *string {
	return s.ObjectType
}

func (s *ListWorkflowResponseBodyDataScene) GetStatus() *int64 {
	return s.Status
}

func (s *ListWorkflowResponseBodyDataScene) SetBizUsage(v string) *ListWorkflowResponseBodyDataScene {
	s.BizUsage = &v
	return s
}

func (s *ListWorkflowResponseBodyDataScene) SetId(v int64) *ListWorkflowResponseBodyDataScene {
	s.Id = &v
	return s
}

func (s *ListWorkflowResponseBodyDataScene) SetName(v string) *ListWorkflowResponseBodyDataScene {
	s.Name = &v
	return s
}

func (s *ListWorkflowResponseBodyDataScene) SetObjectType(v string) *ListWorkflowResponseBodyDataScene {
	s.ObjectType = &v
	return s
}

func (s *ListWorkflowResponseBodyDataScene) SetStatus(v int64) *ListWorkflowResponseBodyDataScene {
	s.Status = &v
	return s
}

func (s *ListWorkflowResponseBodyDataScene) Validate() error {
	return dara.Validate(s)
}
