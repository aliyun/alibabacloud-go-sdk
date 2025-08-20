// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSerDeInfoModel interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *SerDeInfoModel
	GetName() *string
	SetParameters(v map[string]*string) *SerDeInfoModel
	GetParameters() map[string]*string
	SetSerDeId(v int64) *SerDeInfoModel
	GetSerDeId() *int64
	SetSerializationLib(v string) *SerDeInfoModel
	GetSerializationLib() *string
}

type SerDeInfoModel struct {
	Name             *string            `json:"Name,omitempty" xml:"Name,omitempty"`
	Parameters       map[string]*string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	SerDeId          *int64             `json:"SerDeId,omitempty" xml:"SerDeId,omitempty"`
	SerializationLib *string            `json:"SerializationLib,omitempty" xml:"SerializationLib,omitempty"`
}

func (s SerDeInfoModel) String() string {
	return dara.Prettify(s)
}

func (s SerDeInfoModel) GoString() string {
	return s.String()
}

func (s *SerDeInfoModel) GetName() *string {
	return s.Name
}

func (s *SerDeInfoModel) GetParameters() map[string]*string {
	return s.Parameters
}

func (s *SerDeInfoModel) GetSerDeId() *int64 {
	return s.SerDeId
}

func (s *SerDeInfoModel) GetSerializationLib() *string {
	return s.SerializationLib
}

func (s *SerDeInfoModel) SetName(v string) *SerDeInfoModel {
	s.Name = &v
	return s
}

func (s *SerDeInfoModel) SetParameters(v map[string]*string) *SerDeInfoModel {
	s.Parameters = v
	return s
}

func (s *SerDeInfoModel) SetSerDeId(v int64) *SerDeInfoModel {
	s.SerDeId = &v
	return s
}

func (s *SerDeInfoModel) SetSerializationLib(v string) *SerDeInfoModel {
	s.SerializationLib = &v
	return s
}

func (s *SerDeInfoModel) Validate() error {
	return dara.Validate(s)
}
