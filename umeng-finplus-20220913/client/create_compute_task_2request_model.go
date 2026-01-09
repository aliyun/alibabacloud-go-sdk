// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComputeTask2Request interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v int64) *CreateComputeTask2Request
	GetAppId() *int64
	SetClientId(v int64) *CreateComputeTask2Request
	GetClientId() *int64
	SetDataSetIds(v string) *CreateComputeTask2Request
	GetDataSetIds() *string
	SetMorseInfoList(v []*CreateComputeTask2RequestMorseInfoList) *CreateComputeTask2Request
	GetMorseInfoList() []*CreateComputeTask2RequestMorseInfoList
	SetName(v string) *CreateComputeTask2Request
	GetName() *string
	SetRemarks(v string) *CreateComputeTask2Request
	GetRemarks() *string
	SetTaskSource(v string) *CreateComputeTask2Request
	GetTaskSource() *string
	SetType(v string) *CreateComputeTask2Request
	GetType() *string
}

type CreateComputeTask2Request struct {
	AppId         *int64                                    `json:"appId,omitempty" xml:"appId,omitempty"`
	ClientId      *int64                                    `json:"clientId,omitempty" xml:"clientId,omitempty"`
	DataSetIds    *string                                   `json:"dataSetIds,omitempty" xml:"dataSetIds,omitempty"`
	MorseInfoList []*CreateComputeTask2RequestMorseInfoList `json:"morseInfoList,omitempty" xml:"morseInfoList,omitempty" type:"Repeated"`
	Name          *string                                   `json:"name,omitempty" xml:"name,omitempty"`
	Remarks       *string                                   `json:"remarks,omitempty" xml:"remarks,omitempty"`
	TaskSource    *string                                   `json:"taskSource,omitempty" xml:"taskSource,omitempty"`
	Type          *string                                   `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateComputeTask2Request) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeTask2Request) GoString() string {
	return s.String()
}

func (s *CreateComputeTask2Request) GetAppId() *int64 {
	return s.AppId
}

func (s *CreateComputeTask2Request) GetClientId() *int64 {
	return s.ClientId
}

func (s *CreateComputeTask2Request) GetDataSetIds() *string {
	return s.DataSetIds
}

func (s *CreateComputeTask2Request) GetMorseInfoList() []*CreateComputeTask2RequestMorseInfoList {
	return s.MorseInfoList
}

func (s *CreateComputeTask2Request) GetName() *string {
	return s.Name
}

func (s *CreateComputeTask2Request) GetRemarks() *string {
	return s.Remarks
}

func (s *CreateComputeTask2Request) GetTaskSource() *string {
	return s.TaskSource
}

func (s *CreateComputeTask2Request) GetType() *string {
	return s.Type
}

func (s *CreateComputeTask2Request) SetAppId(v int64) *CreateComputeTask2Request {
	s.AppId = &v
	return s
}

func (s *CreateComputeTask2Request) SetClientId(v int64) *CreateComputeTask2Request {
	s.ClientId = &v
	return s
}

func (s *CreateComputeTask2Request) SetDataSetIds(v string) *CreateComputeTask2Request {
	s.DataSetIds = &v
	return s
}

func (s *CreateComputeTask2Request) SetMorseInfoList(v []*CreateComputeTask2RequestMorseInfoList) *CreateComputeTask2Request {
	s.MorseInfoList = v
	return s
}

func (s *CreateComputeTask2Request) SetName(v string) *CreateComputeTask2Request {
	s.Name = &v
	return s
}

func (s *CreateComputeTask2Request) SetRemarks(v string) *CreateComputeTask2Request {
	s.Remarks = &v
	return s
}

func (s *CreateComputeTask2Request) SetTaskSource(v string) *CreateComputeTask2Request {
	s.TaskSource = &v
	return s
}

func (s *CreateComputeTask2Request) SetType(v string) *CreateComputeTask2Request {
	s.Type = &v
	return s
}

func (s *CreateComputeTask2Request) Validate() error {
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

type CreateComputeTask2RequestMorseInfoList struct {
	CuId      *string `json:"cuId,omitempty" xml:"cuId,omitempty"`
	CuVersion *string `json:"cuVersion,omitempty" xml:"cuVersion,omitempty"`
}

func (s CreateComputeTask2RequestMorseInfoList) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeTask2RequestMorseInfoList) GoString() string {
	return s.String()
}

func (s *CreateComputeTask2RequestMorseInfoList) GetCuId() *string {
	return s.CuId
}

func (s *CreateComputeTask2RequestMorseInfoList) GetCuVersion() *string {
	return s.CuVersion
}

func (s *CreateComputeTask2RequestMorseInfoList) SetCuId(v string) *CreateComputeTask2RequestMorseInfoList {
	s.CuId = &v
	return s
}

func (s *CreateComputeTask2RequestMorseInfoList) SetCuVersion(v string) *CreateComputeTask2RequestMorseInfoList {
	s.CuVersion = &v
	return s
}

func (s *CreateComputeTask2RequestMorseInfoList) Validate() error {
	return dara.Validate(s)
}
