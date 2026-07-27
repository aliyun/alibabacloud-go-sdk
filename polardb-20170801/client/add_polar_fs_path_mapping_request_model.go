// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPolarFsPathMappingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomBucketPathList(v []*AddPolarFsPathMappingRequestCustomBucketPathList) *AddPolarFsPathMappingRequest
	GetCustomBucketPathList() []*AddPolarFsPathMappingRequestCustomBucketPathList
	SetDBClusterId(v string) *AddPolarFsPathMappingRequest
	GetDBClusterId() *string
	SetPolarFsInstanceId(v string) *AddPolarFsPathMappingRequest
	GetPolarFsInstanceId() *string
}

type AddPolarFsPathMappingRequest struct {
	// The bucket and corresponding path information.
	CustomBucketPathList []*AddPolarFsPathMappingRequestCustomBucketPathList `json:"CustomBucketPathList,omitempty" xml:"CustomBucketPathList,omitempty" type:"Repeated"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The PolarFS instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pfs-2ze0i74ka607*****
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
}

func (s AddPolarFsPathMappingRequest) String() string {
	return dara.Prettify(s)
}

func (s AddPolarFsPathMappingRequest) GoString() string {
	return s.String()
}

func (s *AddPolarFsPathMappingRequest) GetCustomBucketPathList() []*AddPolarFsPathMappingRequestCustomBucketPathList {
	return s.CustomBucketPathList
}

func (s *AddPolarFsPathMappingRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *AddPolarFsPathMappingRequest) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *AddPolarFsPathMappingRequest) SetCustomBucketPathList(v []*AddPolarFsPathMappingRequestCustomBucketPathList) *AddPolarFsPathMappingRequest {
	s.CustomBucketPathList = v
	return s
}

func (s *AddPolarFsPathMappingRequest) SetDBClusterId(v string) *AddPolarFsPathMappingRequest {
	s.DBClusterId = &v
	return s
}

func (s *AddPolarFsPathMappingRequest) SetPolarFsInstanceId(v string) *AddPolarFsPathMappingRequest {
	s.PolarFsInstanceId = &v
	return s
}

func (s *AddPolarFsPathMappingRequest) Validate() error {
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

type AddPolarFsPathMappingRequestCustomBucketPathList struct {
	// The bucket name.
	//
	// example:
	//
	// Bucket1
	Bucket                *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	BucketAccessKeyId     *string `json:"BucketAccessKeyId,omitempty" xml:"BucketAccessKeyId,omitempty"`
	BucketAccessKeySecret *string `json:"BucketAccessKeySecret,omitempty" xml:"BucketAccessKeySecret,omitempty"`
	// The custom storage path.
	//
	// example:
	//
	// /data1
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s AddPolarFsPathMappingRequestCustomBucketPathList) String() string {
	return dara.Prettify(s)
}

func (s AddPolarFsPathMappingRequestCustomBucketPathList) GoString() string {
	return s.String()
}

func (s *AddPolarFsPathMappingRequestCustomBucketPathList) GetBucket() *string {
	return s.Bucket
}

func (s *AddPolarFsPathMappingRequestCustomBucketPathList) GetBucketAccessKeyId() *string {
	return s.BucketAccessKeyId
}

func (s *AddPolarFsPathMappingRequestCustomBucketPathList) GetBucketAccessKeySecret() *string {
	return s.BucketAccessKeySecret
}

func (s *AddPolarFsPathMappingRequestCustomBucketPathList) GetPath() *string {
	return s.Path
}

func (s *AddPolarFsPathMappingRequestCustomBucketPathList) SetBucket(v string) *AddPolarFsPathMappingRequestCustomBucketPathList {
	s.Bucket = &v
	return s
}

func (s *AddPolarFsPathMappingRequestCustomBucketPathList) SetBucketAccessKeyId(v string) *AddPolarFsPathMappingRequestCustomBucketPathList {
	s.BucketAccessKeyId = &v
	return s
}

func (s *AddPolarFsPathMappingRequestCustomBucketPathList) SetBucketAccessKeySecret(v string) *AddPolarFsPathMappingRequestCustomBucketPathList {
	s.BucketAccessKeySecret = &v
	return s
}

func (s *AddPolarFsPathMappingRequestCustomBucketPathList) SetPath(v string) *AddPolarFsPathMappingRequestCustomBucketPathList {
	s.Path = &v
	return s
}

func (s *AddPolarFsPathMappingRequestCustomBucketPathList) Validate() error {
	return dara.Validate(s)
}
