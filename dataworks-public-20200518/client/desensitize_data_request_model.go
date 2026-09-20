// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDesensitizeDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *DesensitizeDataRequest
	GetData() *string
	SetSceneCode(v string) *DesensitizeDataRequest
	GetSceneCode() *string
}

type DesensitizeDataRequest struct {
	// The input data to be masked.
	//
	// This parameter is required.
	//
	// example:
	//
	// 15365291784
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The masking scene code. You can view this on the Data Masking Management page of DataWorks Data Protection Umbrella in the DataWorks console.
	//
	// You can obtain this value from Data[].SceneCode in the response of DsgSceneQuerySceneListByName.
	//
	// This parameter is required.
	//
	// example:
	//
	// _default_scene_code
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
}

func (s DesensitizeDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DesensitizeDataRequest) GoString() string {
	return s.String()
}

func (s *DesensitizeDataRequest) GetData() *string {
	return s.Data
}

func (s *DesensitizeDataRequest) GetSceneCode() *string {
	return s.SceneCode
}

func (s *DesensitizeDataRequest) SetData(v string) *DesensitizeDataRequest {
	s.Data = &v
	return s
}

func (s *DesensitizeDataRequest) SetSceneCode(v string) *DesensitizeDataRequest {
	s.SceneCode = &v
	return s
}

func (s *DesensitizeDataRequest) Validate() error {
	return dara.Validate(s)
}
