// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolarFsPathMappingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomBucketPathList(v []*DeletePolarFsPathMappingRequestCustomBucketPathList) *DeletePolarFsPathMappingRequest
	GetCustomBucketPathList() []*DeletePolarFsPathMappingRequestCustomBucketPathList
	SetDBClusterId(v string) *DeletePolarFsPathMappingRequest
	GetDBClusterId() *string
	SetPolarFsInstanceId(v string) *DeletePolarFsPathMappingRequest
	GetPolarFsInstanceId() *string
}

type DeletePolarFsPathMappingRequest struct {
	// The list of bucket-path mappings to delete.
	CustomBucketPathList []*DeletePolarFsPathMappingRequestCustomBucketPathList `json:"CustomBucketPathList,omitempty" xml:"CustomBucketPathList,omitempty" type:"Repeated"`
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the PolarFS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pfs-2ze0i74ka607*****
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
}

func (s DeletePolarFsPathMappingRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePolarFsPathMappingRequest) GoString() string {
	return s.String()
}

func (s *DeletePolarFsPathMappingRequest) GetCustomBucketPathList() []*DeletePolarFsPathMappingRequestCustomBucketPathList {
	return s.CustomBucketPathList
}

func (s *DeletePolarFsPathMappingRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeletePolarFsPathMappingRequest) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *DeletePolarFsPathMappingRequest) SetCustomBucketPathList(v []*DeletePolarFsPathMappingRequestCustomBucketPathList) *DeletePolarFsPathMappingRequest {
	s.CustomBucketPathList = v
	return s
}

func (s *DeletePolarFsPathMappingRequest) SetDBClusterId(v string) *DeletePolarFsPathMappingRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeletePolarFsPathMappingRequest) SetPolarFsInstanceId(v string) *DeletePolarFsPathMappingRequest {
	s.PolarFsInstanceId = &v
	return s
}

func (s *DeletePolarFsPathMappingRequest) Validate() error {
	if s.CustomBucketPathList != nil {
		for _, item := range s.CustomBucketPathList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DeletePolarFsPathMappingRequestCustomBucketPathList struct {
	// The name of the bucket.
	//
	// example:
	//
	// Bucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The path in the bucket.
	//
	// example:
	//
	// /data
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s DeletePolarFsPathMappingRequestCustomBucketPathList) String() string {
	return dara.Prettify(s)
}

func (s DeletePolarFsPathMappingRequestCustomBucketPathList) GoString() string {
	return s.String()
}

func (s *DeletePolarFsPathMappingRequestCustomBucketPathList) GetBucket() *string {
	return s.Bucket
}

func (s *DeletePolarFsPathMappingRequestCustomBucketPathList) GetPath() *string {
	return s.Path
}

func (s *DeletePolarFsPathMappingRequestCustomBucketPathList) SetBucket(v string) *DeletePolarFsPathMappingRequestCustomBucketPathList {
	s.Bucket = &v
	return s
}

func (s *DeletePolarFsPathMappingRequestCustomBucketPathList) SetPath(v string) *DeletePolarFsPathMappingRequestCustomBucketPathList {
	s.Path = &v
	return s
}

func (s *DeletePolarFsPathMappingRequestCustomBucketPathList) Validate() error {
	return dara.Validate(s)
}
