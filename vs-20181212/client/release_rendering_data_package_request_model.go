// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseRenderingDataPackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataPackageId(v string) *ReleaseRenderingDataPackageRequest
	GetDataPackageId() *string
}

type ReleaseRenderingDataPackageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dp-9f8c57355d224ad7beaf95e145f22111
	DataPackageId *string `json:"DataPackageId,omitempty" xml:"DataPackageId,omitempty"`
}

func (s ReleaseRenderingDataPackageRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseRenderingDataPackageRequest) GoString() string {
	return s.String()
}

func (s *ReleaseRenderingDataPackageRequest) GetDataPackageId() *string {
	return s.DataPackageId
}

func (s *ReleaseRenderingDataPackageRequest) SetDataPackageId(v string) *ReleaseRenderingDataPackageRequest {
	s.DataPackageId = &v
	return s
}

func (s *ReleaseRenderingDataPackageRequest) Validate() error {
	return dara.Validate(s)
}
