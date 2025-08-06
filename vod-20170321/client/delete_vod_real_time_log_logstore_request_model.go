// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVodRealTimeLogLogstoreRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLogstore(v string) *DeleteVodRealTimeLogLogstoreRequest
	GetLogstore() *string
	SetOwnerId(v int64) *DeleteVodRealTimeLogLogstoreRequest
	GetOwnerId() *int64
	SetProject(v string) *DeleteVodRealTimeLogLogstoreRequest
	GetProject() *string
	SetRegion(v string) *DeleteVodRealTimeLogLogstoreRequest
	GetRegion() *string
}

type DeleteVodRealTimeLogLogstoreRequest struct {
	// This parameter is required.
	Logstore *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// This parameter is required.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DeleteVodRealTimeLogLogstoreRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVodRealTimeLogLogstoreRequest) GoString() string {
	return s.String()
}

func (s *DeleteVodRealTimeLogLogstoreRequest) GetLogstore() *string {
	return s.Logstore
}

func (s *DeleteVodRealTimeLogLogstoreRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteVodRealTimeLogLogstoreRequest) GetProject() *string {
	return s.Project
}

func (s *DeleteVodRealTimeLogLogstoreRequest) GetRegion() *string {
	return s.Region
}

func (s *DeleteVodRealTimeLogLogstoreRequest) SetLogstore(v string) *DeleteVodRealTimeLogLogstoreRequest {
	s.Logstore = &v
	return s
}

func (s *DeleteVodRealTimeLogLogstoreRequest) SetOwnerId(v int64) *DeleteVodRealTimeLogLogstoreRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteVodRealTimeLogLogstoreRequest) SetProject(v string) *DeleteVodRealTimeLogLogstoreRequest {
	s.Project = &v
	return s
}

func (s *DeleteVodRealTimeLogLogstoreRequest) SetRegion(v string) *DeleteVodRealTimeLogLogstoreRequest {
	s.Region = &v
	return s
}

func (s *DeleteVodRealTimeLogLogstoreRequest) Validate() error {
	return dara.Validate(s)
}
