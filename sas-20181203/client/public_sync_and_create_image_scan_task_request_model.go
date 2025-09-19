// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublicSyncAndCreateImageScanTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImages(v string) *PublicSyncAndCreateImageScanTaskRequest
	GetImages() *string
	SetSourceIp(v string) *PublicSyncAndCreateImageScanTaskRequest
	GetSourceIp() *string
}

type PublicSyncAndCreateImageScanTaskRequest struct {
	// The information about the images. The value of this parameter is in the JSON format and contains the following fields:
	//
	// 	- **RegistryType**: the type of the image repository. Valid values:
	//
	//     	- acr
	//
	//     	- harbor
	//
	//     	- quay
	//
	// 	- **RepoId**: the ID of the image repository.
	//
	// 	- **InstanceId**: the ID of the Container Registry instance to which the image repository belongs.
	//
	// 	- **RepoNamespace**: the namespace to which the image repository belongs.
	//
	// 	- **RegionId**: the region ID of the image.
	//
	// 	- **RepoName**: the name of the image repository.
	//
	// 	- **Digest**: the digest of the image.
	//
	// 	- **Tag**: the tag that is added to the image.
	//
	// 	- **CreateTime**: the timestamp when the image was created. Unit: milliseconds.
	//
	// 	- **UpdateTime**: the timestamp when the image was updated. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"registryType":"acr","repoId":"crr-red9kgqqi0fh75d5","instanceId":"cri-1x2qtjd75ofqvfa8","repoNamespace":"1125","regionId":"cn-hangzhou","repoName":"script1","digest":"b1064a5007786b6f82a50772858e8156bc1e5fbddb7dd20af6922761ad197fe4","tag":"v0113-ver2","createTime":1641267666666,"updateTime":1641267909090}]
	Images *string `json:"Images,omitempty" xml:"Images,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 1.2.X.X
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s PublicSyncAndCreateImageScanTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s PublicSyncAndCreateImageScanTaskRequest) GoString() string {
	return s.String()
}

func (s *PublicSyncAndCreateImageScanTaskRequest) GetImages() *string {
	return s.Images
}

func (s *PublicSyncAndCreateImageScanTaskRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *PublicSyncAndCreateImageScanTaskRequest) SetImages(v string) *PublicSyncAndCreateImageScanTaskRequest {
	s.Images = &v
	return s
}

func (s *PublicSyncAndCreateImageScanTaskRequest) SetSourceIp(v string) *PublicSyncAndCreateImageScanTaskRequest {
	s.SourceIp = &v
	return s
}

func (s *PublicSyncAndCreateImageScanTaskRequest) Validate() error {
	return dara.Validate(s)
}
