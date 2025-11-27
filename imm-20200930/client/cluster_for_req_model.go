// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClusterForReq interface {
	dara.Model
	String() string
	GoString() string
	SetCover(v *ClusterForReqCover) *ClusterForReq
	GetCover() *ClusterForReqCover
	SetCustomId(v string) *ClusterForReq
	GetCustomId() *string
	SetCustomLabels(v map[string]interface{}) *ClusterForReq
	GetCustomLabels() map[string]interface{}
	SetName(v string) *ClusterForReq
	GetName() *string
	SetObjectId(v string) *ClusterForReq
	GetObjectId() *string
}

type ClusterForReq struct {
	Cover        *ClusterForReqCover    `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	CustomId     *string                `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	CustomLabels map[string]interface{} `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	Name         *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	ObjectId     *string                `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
}

func (s ClusterForReq) String() string {
	return dara.Prettify(s)
}

func (s ClusterForReq) GoString() string {
	return s.String()
}

func (s *ClusterForReq) GetCover() *ClusterForReqCover {
	return s.Cover
}

func (s *ClusterForReq) GetCustomId() *string {
	return s.CustomId
}

func (s *ClusterForReq) GetCustomLabels() map[string]interface{} {
	return s.CustomLabels
}

func (s *ClusterForReq) GetName() *string {
	return s.Name
}

func (s *ClusterForReq) GetObjectId() *string {
	return s.ObjectId
}

func (s *ClusterForReq) SetCover(v *ClusterForReqCover) *ClusterForReq {
	s.Cover = v
	return s
}

func (s *ClusterForReq) SetCustomId(v string) *ClusterForReq {
	s.CustomId = &v
	return s
}

func (s *ClusterForReq) SetCustomLabels(v map[string]interface{}) *ClusterForReq {
	s.CustomLabels = v
	return s
}

func (s *ClusterForReq) SetName(v string) *ClusterForReq {
	s.Name = &v
	return s
}

func (s *ClusterForReq) SetObjectId(v string) *ClusterForReq {
	s.ObjectId = &v
	return s
}

func (s *ClusterForReq) Validate() error {
	if s.Cover != nil {
		if err := s.Cover.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClusterForReqCover struct {
	Figures []*ClusterForReqCoverFigures `json:"Figures,omitempty" xml:"Figures,omitempty" type:"Repeated"`
}

func (s ClusterForReqCover) String() string {
	return dara.Prettify(s)
}

func (s ClusterForReqCover) GoString() string {
	return s.String()
}

func (s *ClusterForReqCover) GetFigures() []*ClusterForReqCoverFigures {
	return s.Figures
}

func (s *ClusterForReqCover) SetFigures(v []*ClusterForReqCoverFigures) *ClusterForReqCover {
	s.Figures = v
	return s
}

func (s *ClusterForReqCover) Validate() error {
	if s.Figures != nil {
		for _, item := range s.Figures {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ClusterForReqCoverFigures struct {
	FigureId *string `json:"FigureId,omitempty" xml:"FigureId,omitempty"`
}

func (s ClusterForReqCoverFigures) String() string {
	return dara.Prettify(s)
}

func (s ClusterForReqCoverFigures) GoString() string {
	return s.String()
}

func (s *ClusterForReqCoverFigures) GetFigureId() *string {
	return s.FigureId
}

func (s *ClusterForReqCoverFigures) SetFigureId(v string) *ClusterForReqCoverFigures {
	s.FigureId = &v
	return s
}

func (s *ClusterForReqCoverFigures) Validate() error {
	return dara.Validate(s)
}
