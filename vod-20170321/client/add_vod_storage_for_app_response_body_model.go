// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddVodStorageForAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddVodStorageForAppResponseBody
	GetRequestId() *string
	SetStorageLocation(v string) *AddVodStorageForAppResponseBody
	GetStorageLocation() *string
}

type AddVodStorageForAppResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4A*****F6-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The address of the VOD bucket.
	//
	// example:
	//
	// out-****.oss-cn-shanghai.aliyuncs.com
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
}

func (s AddVodStorageForAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddVodStorageForAppResponseBody) GoString() string {
	return s.String()
}

func (s *AddVodStorageForAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddVodStorageForAppResponseBody) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *AddVodStorageForAppResponseBody) SetRequestId(v string) *AddVodStorageForAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddVodStorageForAppResponseBody) SetStorageLocation(v string) *AddVodStorageForAppResponseBody {
	s.StorageLocation = &v
	return s
}

func (s *AddVodStorageForAppResponseBody) Validate() error {
	return dara.Validate(s)
}
