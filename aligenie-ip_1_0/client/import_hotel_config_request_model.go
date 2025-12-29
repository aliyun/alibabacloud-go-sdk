// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportHotelConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *ImportHotelConfigRequest
	GetHotelId() *string
	SetImportHotelConfig(v *ImportHotelConfigRequestImportHotelConfig) *ImportHotelConfigRequest
	GetImportHotelConfig() *ImportHotelConfigRequestImportHotelConfig
}

type ImportHotelConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a7***83
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	ImportHotelConfig *ImportHotelConfigRequestImportHotelConfig `json:"ImportHotelConfig,omitempty" xml:"ImportHotelConfig,omitempty" type:"Struct"`
}

func (s ImportHotelConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportHotelConfigRequest) GoString() string {
	return s.String()
}

func (s *ImportHotelConfigRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *ImportHotelConfigRequest) GetImportHotelConfig() *ImportHotelConfigRequestImportHotelConfig {
	return s.ImportHotelConfig
}

func (s *ImportHotelConfigRequest) SetHotelId(v string) *ImportHotelConfigRequest {
	s.HotelId = &v
	return s
}

func (s *ImportHotelConfigRequest) SetImportHotelConfig(v *ImportHotelConfigRequestImportHotelConfig) *ImportHotelConfigRequest {
	s.ImportHotelConfig = v
	return s
}

func (s *ImportHotelConfigRequest) Validate() error {
	if s.ImportHotelConfig != nil {
		if err := s.ImportHotelConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImportHotelConfigRequestImportHotelConfig struct {
	RcuCustomScenes []*ImportHotelConfigRequestImportHotelConfigRcuCustomScenes `json:"RcuCustomScenes,omitempty" xml:"RcuCustomScenes,omitempty" type:"Repeated"`
}

func (s ImportHotelConfigRequestImportHotelConfig) String() string {
	return dara.Prettify(s)
}

func (s ImportHotelConfigRequestImportHotelConfig) GoString() string {
	return s.String()
}

func (s *ImportHotelConfigRequestImportHotelConfig) GetRcuCustomScenes() []*ImportHotelConfigRequestImportHotelConfigRcuCustomScenes {
	return s.RcuCustomScenes
}

func (s *ImportHotelConfigRequestImportHotelConfig) SetRcuCustomScenes(v []*ImportHotelConfigRequestImportHotelConfigRcuCustomScenes) *ImportHotelConfigRequestImportHotelConfig {
	s.RcuCustomScenes = v
	return s
}

func (s *ImportHotelConfigRequestImportHotelConfig) Validate() error {
	if s.RcuCustomScenes != nil {
		for _, item := range s.RcuCustomScenes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ImportHotelConfigRequestImportHotelConfigRcuCustomScenes struct {
	// This parameter is required.
	CorpusList  []*string `json:"CorpusList,omitempty" xml:"CorpusList,omitempty" type:"Repeated"`
	Description *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	Icon        *string   `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s ImportHotelConfigRequestImportHotelConfigRcuCustomScenes) String() string {
	return dara.Prettify(s)
}

func (s ImportHotelConfigRequestImportHotelConfigRcuCustomScenes) GoString() string {
	return s.String()
}

func (s *ImportHotelConfigRequestImportHotelConfigRcuCustomScenes) GetCorpusList() []*string {
	return s.CorpusList
}

func (s *ImportHotelConfigRequestImportHotelConfigRcuCustomScenes) GetDescription() *string {
	return s.Description
}

func (s *ImportHotelConfigRequestImportHotelConfigRcuCustomScenes) GetIcon() *string {
	return s.Icon
}

func (s *ImportHotelConfigRequestImportHotelConfigRcuCustomScenes) GetName() *string {
	return s.Name
}

func (s *ImportHotelConfigRequestImportHotelConfigRcuCustomScenes) GetSceneId() *string {
	return s.SceneId
}

func (s *ImportHotelConfigRequestImportHotelConfigRcuCustomScenes) SetCorpusList(v []*string) *ImportHotelConfigRequestImportHotelConfigRcuCustomScenes {
	s.CorpusList = v
	return s
}

func (s *ImportHotelConfigRequestImportHotelConfigRcuCustomScenes) SetDescription(v string) *ImportHotelConfigRequestImportHotelConfigRcuCustomScenes {
	s.Description = &v
	return s
}

func (s *ImportHotelConfigRequestImportHotelConfigRcuCustomScenes) SetIcon(v string) *ImportHotelConfigRequestImportHotelConfigRcuCustomScenes {
	s.Icon = &v
	return s
}

func (s *ImportHotelConfigRequestImportHotelConfigRcuCustomScenes) SetName(v string) *ImportHotelConfigRequestImportHotelConfigRcuCustomScenes {
	s.Name = &v
	return s
}

func (s *ImportHotelConfigRequestImportHotelConfigRcuCustomScenes) SetSceneId(v string) *ImportHotelConfigRequestImportHotelConfigRcuCustomScenes {
	s.SceneId = &v
	return s
}

func (s *ImportHotelConfigRequestImportHotelConfigRcuCustomScenes) Validate() error {
	return dara.Validate(s)
}
