// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFigureClusterForReq interface {
	dara.Model
	String() string
	GoString() string
	SetCover(v *FigureClusterForReqCover) *FigureClusterForReq
	GetCover() *FigureClusterForReqCover
	SetCustomId(v string) *FigureClusterForReq
	GetCustomId() *string
	SetCustomLabels(v map[string]interface{}) *FigureClusterForReq
	GetCustomLabels() map[string]interface{}
	SetMetaLockVersion(v int64) *FigureClusterForReq
	GetMetaLockVersion() *int64
	SetName(v string) *FigureClusterForReq
	GetName() *string
	SetObjectId(v string) *FigureClusterForReq
	GetObjectId() *string
}

type FigureClusterForReq struct {
	// The cover of the cluster.
	Cover *FigureClusterForReqCover `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	// The custom ID.
	//
	// example:
	//
	// abc
	CustomId *string `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	// A key-value map of custom labels that can be used to search for the figure cluster.
	//
	// example:
	//
	// {"Bucket": "examplebucket"}
	CustomLabels map[string]interface{} `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	// The metadata lock version. This parameter enables optimistic locking to prevent conflicting updates during high concurrency. To ensure data consistency, pass the current MetaLockVersion value when updating a figure cluster. You can retrieve this value from a Get or List operation. The service proceeds with the update only if the provided version matches the system\\"s current version. The initial value is 0 and is incremented by 1 after each successful update.
	//
	// example:
	//
	// 1
	MetaLockVersion *int64 `json:"MetaLockVersion,omitempty" xml:"MetaLockVersion,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// abc
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The object ID of the cluster.
	//
	// example:
	//
	// Cluster-99b1c333-86dc-45da-8c6****
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
}

func (s FigureClusterForReq) String() string {
	return dara.Prettify(s)
}

func (s FigureClusterForReq) GoString() string {
	return s.String()
}

func (s *FigureClusterForReq) GetCover() *FigureClusterForReqCover {
	return s.Cover
}

func (s *FigureClusterForReq) GetCustomId() *string {
	return s.CustomId
}

func (s *FigureClusterForReq) GetCustomLabels() map[string]interface{} {
	return s.CustomLabels
}

func (s *FigureClusterForReq) GetMetaLockVersion() *int64 {
	return s.MetaLockVersion
}

func (s *FigureClusterForReq) GetName() *string {
	return s.Name
}

func (s *FigureClusterForReq) GetObjectId() *string {
	return s.ObjectId
}

func (s *FigureClusterForReq) SetCover(v *FigureClusterForReqCover) *FigureClusterForReq {
	s.Cover = v
	return s
}

func (s *FigureClusterForReq) SetCustomId(v string) *FigureClusterForReq {
	s.CustomId = &v
	return s
}

func (s *FigureClusterForReq) SetCustomLabels(v map[string]interface{}) *FigureClusterForReq {
	s.CustomLabels = v
	return s
}

func (s *FigureClusterForReq) SetMetaLockVersion(v int64) *FigureClusterForReq {
	s.MetaLockVersion = &v
	return s
}

func (s *FigureClusterForReq) SetName(v string) *FigureClusterForReq {
	s.Name = &v
	return s
}

func (s *FigureClusterForReq) SetObjectId(v string) *FigureClusterForReq {
	s.ObjectId = &v
	return s
}

func (s *FigureClusterForReq) Validate() error {
	if s.Cover != nil {
		if err := s.Cover.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FigureClusterForReqCover struct {
	// The figure list.
	Figures []*FigureClusterForReqCoverFigures `json:"Figures,omitempty" xml:"Figures,omitempty" type:"Repeated"`
}

func (s FigureClusterForReqCover) String() string {
	return dara.Prettify(s)
}

func (s FigureClusterForReqCover) GoString() string {
	return s.String()
}

func (s *FigureClusterForReqCover) GetFigures() []*FigureClusterForReqCoverFigures {
	return s.Figures
}

func (s *FigureClusterForReqCover) SetFigures(v []*FigureClusterForReqCoverFigures) *FigureClusterForReqCover {
	s.Figures = v
	return s
}

func (s *FigureClusterForReqCover) Validate() error {
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

type FigureClusterForReqCoverFigures struct {
	// The figure ID.
	//
	// example:
	//
	// 2cb3c51e-b406-4b0c-af1b-897d88e1****
	FigureId *string `json:"FigureId,omitempty" xml:"FigureId,omitempty"`
}

func (s FigureClusterForReqCoverFigures) String() string {
	return dara.Prettify(s)
}

func (s FigureClusterForReqCoverFigures) GoString() string {
	return s.String()
}

func (s *FigureClusterForReqCoverFigures) GetFigureId() *string {
	return s.FigureId
}

func (s *FigureClusterForReqCoverFigures) SetFigureId(v string) *FigureClusterForReqCoverFigures {
	s.FigureId = &v
	return s
}

func (s *FigureClusterForReqCoverFigures) Validate() error {
	return dara.Validate(s)
}
