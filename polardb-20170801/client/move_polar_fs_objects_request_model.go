// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMovePolarFsObjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetObjectsToMove(v []*MovePolarFsObjectsRequestObjectsToMove) *MovePolarFsObjectsRequest
	GetObjectsToMove() []*MovePolarFsObjectsRequestObjectsToMove
	SetPolarFsInstanceId(v string) *MovePolarFsObjectsRequest
	GetPolarFsInstanceId() *string
}

type MovePolarFsObjectsRequest struct {
	ObjectsToMove []*MovePolarFsObjectsRequestObjectsToMove `json:"ObjectsToMove,omitempty" xml:"ObjectsToMove,omitempty" type:"Repeated"`
	// example:
	//
	// pfs-test*****
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
}

func (s MovePolarFsObjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s MovePolarFsObjectsRequest) GoString() string {
	return s.String()
}

func (s *MovePolarFsObjectsRequest) GetObjectsToMove() []*MovePolarFsObjectsRequestObjectsToMove {
	return s.ObjectsToMove
}

func (s *MovePolarFsObjectsRequest) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *MovePolarFsObjectsRequest) SetObjectsToMove(v []*MovePolarFsObjectsRequestObjectsToMove) *MovePolarFsObjectsRequest {
	s.ObjectsToMove = v
	return s
}

func (s *MovePolarFsObjectsRequest) SetPolarFsInstanceId(v string) *MovePolarFsObjectsRequest {
	s.PolarFsInstanceId = &v
	return s
}

func (s *MovePolarFsObjectsRequest) Validate() error {
	if s.ObjectsToMove != nil {
		for _, item := range s.ObjectsToMove {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MovePolarFsObjectsRequestObjectsToMove struct {
	// example:
	//
	// /test1
	DestinationPath *string `json:"DestinationPath,omitempty" xml:"DestinationPath,omitempty"`
	// example:
	//
	// /test
	SourcePath *string `json:"SourcePath,omitempty" xml:"SourcePath,omitempty"`
}

func (s MovePolarFsObjectsRequestObjectsToMove) String() string {
	return dara.Prettify(s)
}

func (s MovePolarFsObjectsRequestObjectsToMove) GoString() string {
	return s.String()
}

func (s *MovePolarFsObjectsRequestObjectsToMove) GetDestinationPath() *string {
	return s.DestinationPath
}

func (s *MovePolarFsObjectsRequestObjectsToMove) GetSourcePath() *string {
	return s.SourcePath
}

func (s *MovePolarFsObjectsRequestObjectsToMove) SetDestinationPath(v string) *MovePolarFsObjectsRequestObjectsToMove {
	s.DestinationPath = &v
	return s
}

func (s *MovePolarFsObjectsRequestObjectsToMove) SetSourcePath(v string) *MovePolarFsObjectsRequestObjectsToMove {
	s.SourcePath = &v
	return s
}

func (s *MovePolarFsObjectsRequestObjectsToMove) Validate() error {
	return dara.Validate(s)
}
