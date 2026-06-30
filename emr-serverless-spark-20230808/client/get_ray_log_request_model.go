// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRayLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucketName(v string) *GetRayLogRequest
	GetBucketName() *string
	SetPath(v string) *GetRayLogRequest
	GetPath() *string
}

type GetRayLogRequest struct {
	// The bucket name.
	//
	// example:
	//
	// mybucket
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// The log file path.
	//
	// example:
	//
	// /w-xxxxxxx/ray/logs/rj-xxxxxxxxxx_default/xxxx/rj-xxxx_driver.log
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
}

func (s GetRayLogRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRayLogRequest) GoString() string {
	return s.String()
}

func (s *GetRayLogRequest) GetBucketName() *string {
	return s.BucketName
}

func (s *GetRayLogRequest) GetPath() *string {
	return s.Path
}

func (s *GetRayLogRequest) SetBucketName(v string) *GetRayLogRequest {
	s.BucketName = &v
	return s
}

func (s *GetRayLogRequest) SetPath(v string) *GetRayLogRequest {
	s.Path = &v
	return s
}

func (s *GetRayLogRequest) Validate() error {
	return dara.Validate(s)
}
