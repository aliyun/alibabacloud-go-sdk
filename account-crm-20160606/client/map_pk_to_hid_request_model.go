// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMapPkToHidRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *MapPkToHidRequest
	GetAppName() *string
	SetMappingScenes(v string) *MapPkToHidRequest
	GetMappingScenes() *string
	SetPk(v string) *MapPkToHidRequest
	GetPk() *string
}

type MapPkToHidRequest struct {
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	MappingScenes *string `json:"MappingScenes,omitempty" xml:"MappingScenes,omitempty"`
	// This parameter is required.
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
}

func (s MapPkToHidRequest) String() string {
	return dara.Prettify(s)
}

func (s MapPkToHidRequest) GoString() string {
	return s.String()
}

func (s *MapPkToHidRequest) GetAppName() *string {
	return s.AppName
}

func (s *MapPkToHidRequest) GetMappingScenes() *string {
	return s.MappingScenes
}

func (s *MapPkToHidRequest) GetPk() *string {
	return s.Pk
}

func (s *MapPkToHidRequest) SetAppName(v string) *MapPkToHidRequest {
	s.AppName = &v
	return s
}

func (s *MapPkToHidRequest) SetMappingScenes(v string) *MapPkToHidRequest {
	s.MappingScenes = &v
	return s
}

func (s *MapPkToHidRequest) SetPk(v string) *MapPkToHidRequest {
	s.Pk = &v
	return s
}

func (s *MapPkToHidRequest) Validate() error {
	return dara.Validate(s)
}
