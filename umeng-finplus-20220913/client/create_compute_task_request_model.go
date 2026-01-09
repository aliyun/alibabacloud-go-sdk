// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComputeTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v int64) *CreateComputeTaskRequest
	GetAppId() *int64
	SetDataSetIds(v string) *CreateComputeTaskRequest
	GetDataSetIds() *string
	SetMorseInfoList(v []*CreateComputeTaskRequestMorseInfoList) *CreateComputeTaskRequest
	GetMorseInfoList() []*CreateComputeTaskRequestMorseInfoList
	SetName(v string) *CreateComputeTaskRequest
	GetName() *string
	SetRemarks(v string) *CreateComputeTaskRequest
	GetRemarks() *string
	SetType(v string) *CreateComputeTaskRequest
	GetType() *string
}

type CreateComputeTaskRequest struct {
	AppId         *int64                                   `json:"appId,omitempty" xml:"appId,omitempty"`
	DataSetIds    *string                                  `json:"dataSetIds,omitempty" xml:"dataSetIds,omitempty"`
	MorseInfoList []*CreateComputeTaskRequestMorseInfoList `json:"morseInfoList,omitempty" xml:"morseInfoList,omitempty" type:"Repeated"`
	Name          *string                                  `json:"name,omitempty" xml:"name,omitempty"`
	Remarks       *string                                  `json:"remarks,omitempty" xml:"remarks,omitempty"`
	Type          *string                                  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateComputeTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateComputeTaskRequest) GetAppId() *int64 {
	return s.AppId
}

func (s *CreateComputeTaskRequest) GetDataSetIds() *string {
	return s.DataSetIds
}

func (s *CreateComputeTaskRequest) GetMorseInfoList() []*CreateComputeTaskRequestMorseInfoList {
	return s.MorseInfoList
}

func (s *CreateComputeTaskRequest) GetName() *string {
	return s.Name
}

func (s *CreateComputeTaskRequest) GetRemarks() *string {
	return s.Remarks
}

func (s *CreateComputeTaskRequest) GetType() *string {
	return s.Type
}

func (s *CreateComputeTaskRequest) SetAppId(v int64) *CreateComputeTaskRequest {
	s.AppId = &v
	return s
}

func (s *CreateComputeTaskRequest) SetDataSetIds(v string) *CreateComputeTaskRequest {
	s.DataSetIds = &v
	return s
}

func (s *CreateComputeTaskRequest) SetMorseInfoList(v []*CreateComputeTaskRequestMorseInfoList) *CreateComputeTaskRequest {
	s.MorseInfoList = v
	return s
}

func (s *CreateComputeTaskRequest) SetName(v string) *CreateComputeTaskRequest {
	s.Name = &v
	return s
}

func (s *CreateComputeTaskRequest) SetRemarks(v string) *CreateComputeTaskRequest {
	s.Remarks = &v
	return s
}

func (s *CreateComputeTaskRequest) SetType(v string) *CreateComputeTaskRequest {
	s.Type = &v
	return s
}

func (s *CreateComputeTaskRequest) Validate() error {
	if s.MorseInfoList != nil {
		for _, item := range s.MorseInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateComputeTaskRequestMorseInfoList struct {
	CuId      *string `json:"cuId,omitempty" xml:"cuId,omitempty"`
	CuVersion *string `json:"cuVersion,omitempty" xml:"cuVersion,omitempty"`
}

func (s CreateComputeTaskRequestMorseInfoList) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeTaskRequestMorseInfoList) GoString() string {
	return s.String()
}

func (s *CreateComputeTaskRequestMorseInfoList) GetCuId() *string {
	return s.CuId
}

func (s *CreateComputeTaskRequestMorseInfoList) GetCuVersion() *string {
	return s.CuVersion
}

func (s *CreateComputeTaskRequestMorseInfoList) SetCuId(v string) *CreateComputeTaskRequestMorseInfoList {
	s.CuId = &v
	return s
}

func (s *CreateComputeTaskRequestMorseInfoList) SetCuVersion(v string) *CreateComputeTaskRequestMorseInfoList {
	s.CuVersion = &v
	return s
}

func (s *CreateComputeTaskRequestMorseInfoList) Validate() error {
	return dara.Validate(s)
}
